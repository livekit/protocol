// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
	"text/template"

	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"

	"github.com/puzpuzpuz/xsync/v4"

	"github.com/livekit/protocol/livekit/logger"
	"github.com/livekit/protocol/utils/must"
)

// Proto returns a zapcore.ObjectMarshaler that records each field annotated
// with a (logger.sensitivity) under that sensitivity, and redacts it when the
// sink cannot record that level. SECRET is never recordable.
func Proto(val proto.Message) zapcore.ObjectMarshaler {
	if val == nil {
		return nil
	}
	return protoMarshaller{m: val.ProtoReflect()}
}

// ProtoWithLimit behaves like Proto, but when the message's wire size exceeds
// maxBytes it logs a compact summary instead of the full contents: message
// type, wire size, and top-level scalar fields (tagged or redacted per
// sensitivity, long strings truncated), with repeated/map fields reduced to
// counts and nested messages dropped. Use for messages of unbounded size, such
// as RPC payloads.
func ProtoWithLimit(val proto.Message, maxBytes int) zapcore.ObjectMarshaler {
	if val == nil {
		return nil
	}
	return protoMarshaller{m: val.ProtoReflect(), maxSize: maxBytes}
}

var _ zapcore.ObjectMarshaler = protoMarshaller{}
var _ zapcore.ObjectMarshaler = protoMapMarshaller{}
var _ zapcore.ArrayMarshaler = protoListMarshaller{}

type protoMarshaller struct {
	m protoreflect.Message
	// maxSize is the wire size in bytes above which the message is logged as
	// a summary instead of its full contents. 0 means no limit.
	maxSize int
}

func (p protoMarshaller) MarshalLogObject(e zapcore.ObjectEncoder) error {
	if !p.m.IsValid() {
		return nil
	}

	if p.maxSize > 0 {
		if size := proto.Size(p.m.Interface()); size > p.maxSize {
			return p.marshalSummary(e, size)
		}
	}

	fields := p.m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		k := f.JSONName()
		if proto.HasExtension(f.Options(), logger.E_Name) {
			k = proto.GetExtension(f.Options(), logger.E_Name).(string)
		}
		v := p.m.Get(f)

		if protoFieldIsZero(f, v) {
			continue
		}

		// Never assign to e: sibling fields must not inherit this field's
		// level. enc records this field, and everything nested under it, at
		// the level the field declares.
		enc := e
		if level := fieldSensitivity(f); level > SensitivityNone {
			se := ObjectEncoderFor(e, level)
			if se == nil {
				e.AddString(k, marshalRedacted(f, v))
				continue
			}
			enc = se
		}

		// Sensitivity is not re-checked for map values or list elements: proto
		// cannot annotate them independently of the containing field, so the
		// check above is complete.
		if f.IsMap() {
			if m := v.Map(); m.IsValid() {
				enc.AddObject(k, protoMapMarshaller{f: f, m: m})
			}
		} else if f.IsList() {
			if m := v.List(); m.IsValid() {
				enc.AddArray(k, protoListMarshaller{f: f, m: m})
			}
		} else {
			marshalProtoField(k, f, v, enc)
		}
	}
	return nil
}

// marshalSummary keeps top-level scalar fields, which do not meaningfully
// contribute to message size, and reduces lists, maps, and long strings to
// counts and sizes.
func (p protoMarshaller) marshalSummary(e zapcore.ObjectEncoder, size int) error {
	e.AddString("truncatedProto", string(p.m.Descriptor().FullName()))
	e.AddInt("protoBytes", size)

	fields := p.m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		v := p.m.Get(f)
		if protoFieldIsZero(f, v) {
			continue
		}
		k := f.JSONName()
		if proto.HasExtension(f.Options(), logger.E_Name) {
			k = proto.GetExtension(f.Options(), logger.E_Name).(string)
		}
		// These disclose no field contents, so they precede the sensitivity
		// decision and stay untagged. Moving the decision above them would
		// replace a count with a redaction.
		switch {
		case f.IsMap():
			e.AddInt(k+"Count", v.Map().Len())
			continue
		case f.IsList():
			e.AddInt(k+"Count", v.List().Len())
			continue
		case f.Kind() == protoreflect.MessageKind, f.Kind() == protoreflect.GroupKind:
			continue
		}

		enc := e
		if level := fieldSensitivity(f); level > SensitivityNone {
			se := ObjectEncoderFor(e, level)
			if se == nil {
				e.AddString(k, marshalRedacted(f, v))
				continue
			}
			enc = se
		}

		// Summaries stay bounded even for values the sink may record.
		if f.Kind() == protoreflect.StringKind {
			enc.AddString(k, marshalTruncatedString(v.String()))
		} else {
			marshalProtoField(k, f, v, enc)
		}
	}
	return nil
}

func marshalTruncatedString(s string) string {
	if len(s) <= 256 {
		return s
	}
	return fmt.Sprintf("%s... (%d bytes)", s[:256], len(s))
}

type protoMapMarshaller struct {
	f protoreflect.FieldDescriptor
	m protoreflect.Map
}

func (p protoMapMarshaller) MarshalLogObject(e zapcore.ObjectEncoder) error {
	p.m.Range(func(ki protoreflect.MapKey, vi protoreflect.Value) bool {
		var k string
		switch p.f.MapKey().Kind() {
		case protoreflect.BoolKind:
			k = strconv.FormatBool(ki.Bool())
		case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
			k = strconv.FormatInt(ki.Int(), 10)
		case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
			k = strconv.FormatUint(ki.Uint(), 10)
		case protoreflect.StringKind:
			k = ki.String()
		}
		marshalProtoField(k, p.f.MapValue(), vi, e)
		return true
	})
	return nil
}

type protoListMarshaller struct {
	f protoreflect.FieldDescriptor
	m protoreflect.List
}

func (p protoListMarshaller) MarshalLogArray(e zapcore.ArrayEncoder) error {
	for i := 0; i < p.m.Len(); i++ {
		v := p.m.Get(i)
		switch p.f.Kind() {
		case protoreflect.BoolKind:
			e.AppendBool(v.Bool())
		case protoreflect.EnumKind:
			e.AppendString(marshalProtoEnum(p.f, v))
		case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
			e.AppendInt64(v.Int())
		case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
			e.AppendUint64(v.Uint())
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			e.AppendFloat64(v.Float())
		case protoreflect.StringKind:
			e.AppendString(v.String())
		case protoreflect.BytesKind:
			e.AppendString(marshalProtoBytes(v.Bytes()))
		case protoreflect.MessageKind:
			e.AppendObject(protoMarshaller{m: v.Message()})
		}
	}
	return nil
}

func marshalProtoField(k string, f protoreflect.FieldDescriptor, v protoreflect.Value, e zapcore.ObjectEncoder) {
	switch f.Kind() {
	case protoreflect.BoolKind:
		e.AddBool(k, v.Bool())
	case protoreflect.EnumKind:
		e.AddString(k, marshalProtoEnum(f, v))
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		e.AddInt64(k, v.Int())
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		e.AddUint64(k, v.Uint())
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		e.AddFloat64(k, v.Float())
	case protoreflect.StringKind:
		e.AddString(k, v.String())
	case protoreflect.BytesKind:
		e.AddString(k, marshalProtoBytes(v.Bytes()))
	case protoreflect.MessageKind:
		e.AddObject(k, protoMarshaller{m: v.Message()})
	}
}

func marshalProtoEnum(f protoreflect.FieldDescriptor, v protoreflect.Value) string {
	if e := f.Enum().Values().ByNumber(v.Enum()); e != nil {
		return string(e.Name())
	}
	return fmt.Sprintf("<UNDEFINED(%d)>", v.Enum())
}

func marshalProtoBytes(b []byte) string {
	n := len(b)
	if n > 64 {
		b = b[:64]
	}
	s := base64.RawStdEncoding.EncodeToString(b)
	switch {
	case n <= 64:
		return s
	case n < 1<<10:
		return fmt.Sprintf("%s... (%dbytes)", s, n)
	case n < 1<<20:
		return fmt.Sprintf("%s... (%.2fkB)", s, float64(n)/float64(1<<10))
	case n < 1<<30:
		return fmt.Sprintf("%s... (%.2fMB)", s, float64(n)/float64(1<<20))
	default:
		return fmt.Sprintf("%s... (%.2fGB)", s, float64(n)/float64(1<<30))
	}
}

var redactTemplates = xsync.NewMap[string, *template.Template]()

func marshalRedacted(f protoreflect.FieldDescriptor, v protoreflect.Value) string {
	if !proto.HasExtension(f.Options(), logger.E_RedactFormat) {
		return "<redacted>"
	}

	text := proto.GetExtension(f.Options(), logger.E_RedactFormat).(string)
	tpl, _ := redactTemplates.LoadOrCompute(text, func() (*template.Template, bool) {
		return template.Must(template.New("format").Parse(text)), false
	})

	var b bytes.Buffer
	must.Do(tpl.Execute(&b, redactTemplateData{f, v}))
	return b.String()
}

type redactTemplateData struct {
	f protoreflect.FieldDescriptor
	v protoreflect.Value
}

func (d redactTemplateData) TextName() string {
	return d.f.TextName()
}

func (d redactTemplateData) Size() string {
	msg := dynamicpb.NewMessage(d.f.ContainingMessage())

	switch {
	case d.f.IsList():
		dst := msg.Mutable(d.f).List()
		src := d.v.List()
		for i := 0; i < src.Len(); i++ {
			dst.Append(src.Get(i))
		}
	case d.f.IsMap():
		dst := msg.Mutable(d.f).Map()
		src := d.v.Map()
		src.Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
			dst.Set(k, v)
			return true
		})
	default:
		msg.Set(d.f, d.v)
	}

	return strconv.Itoa(proto.Size(msg.Interface()))
}

func fieldSensitivity(f protoreflect.FieldDescriptor) logger.Sensitivity {
	if !proto.HasExtension(f.Options(), logger.E_Sensitivity) {
		return logger.Sensitivity_SENSITIVITY_UNSPECIFIED
	}
	return proto.GetExtension(f.Options(), logger.E_Sensitivity).(logger.Sensitivity)
}

func protoFieldIsZero(f protoreflect.FieldDescriptor, v protoreflect.Value) bool {
	if f.IsList() {
		l := v.List()
		return l == nil || l.Len() == 0
	}
	if f.IsMap() {
		m := v.Map()
		return m == nil || m.Len() == 0
	}
	switch f.Kind() {
	case protoreflect.BoolKind:
		return !v.Bool()
	case protoreflect.EnumKind:
		return false
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		return v.Int() == 0
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		return v.Uint() == 0
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return v.Float() == 0
	case protoreflect.StringKind:
		return v.String() == ""
	case protoreflect.BytesKind:
		return len(v.Bytes()) == 0
	case protoreflect.MessageKind:
		return !v.Message().IsValid()
	}
	return true
}
