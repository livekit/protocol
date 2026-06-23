package prototest

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Equals[T proto.Message](t require.TestingT, exp, got T) {
	require.True(t, proto.Equal(exp, got), "\nexp: %v\ngot: %v", exp, got)
}
