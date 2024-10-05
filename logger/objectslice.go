package logger

import "go.uber.org/zap/zapcore"

func ObjectSlice[T zapcore.ObjectMarshaler](s []T) zapcore.ArrayMarshaler {
	return objectSlice[T](s)
}

type objectSlice[T zapcore.ObjectMarshaler] []T

func (a objectSlice[T]) MarshalLogArray(e zapcore.ArrayEncoder) error {
	for _, o := range a {
		e.AppendObject(o)
	}
	return nil
}
