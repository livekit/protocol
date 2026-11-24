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

package utils

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/livekit/protocol/livekit/logger"
)

func CloneProto[T proto.Message](m T) T {
	return proto.Clone(m).(T)
}

func CloneProtoSlice[T proto.Message](ms []T) []T {
	cs := make([]T, len(ms))
	for i := range ms {
		cs[i] = CloneProto(ms[i])
	}
	return cs
}

func CloneProtoRedacted[T proto.Message](m T) T {
	clone := proto.Clone(m).(T)

	var redact func(msg proto.Message)
	redact = func(msg proto.Message) {
		if msg == nil {
			return
		}

		reflected := msg.ProtoReflect()
		reflected.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			if proto.HasExtension(fd.Options(), logger.E_Sensitivity) {
				if proto.GetExtension(fd.Options(), logger.E_Sensitivity).(logger.Sensitivity) >= logger.Sensitivity_SENSITIVITY_PII {
					reflected.Clear(fd)
					return true
				}
			}

			switch {
			case fd.IsMap():
				// fd.Kind() is MessageKind for every map (synthetic entry type), so check the value type.
				if fd.MapValue().Message() == nil {
					break
				}
				v.Map().Range(func(_ protoreflect.MapKey, val protoreflect.Value) bool {
					redact(val.Message().Interface())
					return true
				})

			case fd.Message() == nil:

			case fd.IsList():
				src := v.List()
				for i := 0; i < src.Len(); i++ {
					redact(src.Get(i).Message().Interface())
				}

			default:
				redact(v.Message().Interface())
			}

			return true
		})
	}
	redact(clone)
	return clone
}
