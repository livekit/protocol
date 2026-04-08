package utils

import "unsafe"

func CastStringSlice[OutType, InType ~string](src []InType) []OutType {
	return *(*[]OutType)(unsafe.Pointer(&src))
}
