package genlib_test

import (
	"testing"
	"unsafe"
)

func BenchmarkBranchlessAbs(b *testing.B) {
	var v1, v2, v1a, v2a int64 = 1111, -1111, 0, 0
	var v1neg, v2neg bool = false, false
	for i := 0; i < b.N; i += 1 {
		v1neg = v1 < 0
		v1a = v1 * int64(-*(*uint8)(unsafe.Pointer(&v1neg)))
		v2neg = v2 < 0
		v2a = v2 * int64(-*(*uint8)(unsafe.Pointer(&v2neg)))
		_ = v1a
		_ = v2a
	}
}

func BenchmarkIfElseAbs(b *testing.B) {
	var v1, v2, v1a, v2a int64 = 1111, -1111, 0, 0
	for i := 0; i < b.N; i += 1 {
		if v1 < 0 {
			v1a = -v1
		} else {
			v1a = v1
		}
		if v2 < 0 {
			v2a = -v2
		} else {
			v2a = v2
		}
		_ = v1a
		_ = v2a
	}
}
