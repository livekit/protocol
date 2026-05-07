package protojson

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type MarshalOptions = protojson.MarshalOptions
type UnmarshalOptions = protojson.UnmarshalOptions

func Format(m proto.Message) string {
	return MarshalOptions{
		Multiline:    true,
		AllowPartial: true,
	}.Format(m)
}

func Marshal(m proto.Message) ([]byte, error) {
	return MarshalOptions{
		AllowPartial: true,
	}.Marshal(m)
}

func Unmarshal(b []byte, m proto.Message) error {
	return UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}.Unmarshal(b, m)
}
