package xtwirp

import (
	"encoding/base64"
	"errors"

	"github.com/twitchtv/twirp"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const DetailsMetaKey = "error_details"

type ErrorMeta interface {
	TwirpErrorMeta() map[string]string
}

// WithDetailsFrom sets gRPC/PSRPC error details from src as Twirp error metadata on dst.
//
// If error details implement ErrorMeta, their custom metadata fields will be included as well.
func WithDetailsFrom(dst twirp.Error, src error) twirp.Error {
	st, ok := status.FromError(src)
	if !ok {
		return dst
	}
	return WithDetailsFromStatus(dst, st)
}

// WithDetailsFromStatus sets gRPC error details from status as Twirp error metadata on dst.
//
// If error details implement ErrorMeta, their custom metadata fields will be included as well.
func WithDetailsFromStatus(dst twirp.Error, st *status.Status) twirp.Error {
	if st == nil {
		return dst
	}
	details := st.Details()
	if len(details) == 0 {
		return dst
	}
	for _, d := range details {
		if m, ok := d.(ErrorMeta); ok {
			for k, v := range m.TwirpErrorMeta() {
				dst = dst.WithMeta(k, v)
			}
		}
	}
	p := st.Proto()
	if len(p.Details) == 0 {
		return dst
	}
	data, err := proto.Marshal(p)
	if err != nil {
		return dst
	}
	val := base64.StdEncoding.EncodeToString(data)
	return dst.WithMeta(DetailsMetaKey, val)
}

// StatusFromError is an analog of gRPCs status.FromError, but it also considers
// error details encoded in Twirp metadata.
func StatusFromError(err error) (*status.Status, bool) {
	if st, ok := status.FromError(err); ok {
		return st, true
	}
	var e twirp.Error
	if !errors.As(err, &e) {
		return nil, false
	}
	val := e.Meta(DetailsMetaKey)
	if val == "" {
		return nil, false
	}
	data, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, false
	}
	var p spb.Status
	if err := proto.Unmarshal(data, &p); err != nil {
		return nil, false
	}
	return status.FromProto(&p), true
}
