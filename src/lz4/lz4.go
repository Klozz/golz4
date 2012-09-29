package lz4

// #cgo CFLAGS: -I/home/albert/projects/ls4-read-only
// #cgo LDFLAGS: -L/home/albert/projects/ls4-read-only -llz4
// #include "lz4.h"
// #include "lz4hc.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func Compress(in, out []byte) int {
	n := C.LZ4_compress((*C.char)(unsafe.Pointer(&in[0])), (*C.char)(unsafe.Pointer(&out[0])), C.int(len(in)))
	return (int)(n)
}

func CompressHC(in, out []byte) int {
	n := C.LZ4_compressHC((*C.char)(unsafe.Pointer(&in[0])), (*C.char)(unsafe.Pointer(&out[0])), C.int(len(in)))
	return (int)(n)
}

func CompressBound(isize int) int {
	return (isize + (isize / 255) + 16)
}

func DecompressUnknownOutputSize(in, out []byte) int {
	n := C.LZ4_uncompress_unknownOutputSize((*C.char)(unsafe.Pointer(&in[0])), (*C.char)(unsafe.Pointer(&out[0])), C.int(len(in)), C.int(len(out)))
	return (int)(n)
}

func Decompress(in, out []byte) error {
	n := C.LZ4_uncompress((*C.char)(unsafe.Pointer(&in[0])), (*C.char)(unsafe.Pointer(&out[0])), C.int(len(out)))
	if n < 0 {
		return fmt.Errorf("Decompressiong error. Got %d", n)
	}
	return nil
}
