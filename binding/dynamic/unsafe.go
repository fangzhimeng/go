package dynamic

import (
	"unsafe"
	"github.com/thrift-iterator/go/spi"
	"github.com/thrift-iterator/go/protocol"
)

type internalDecoder interface {
	decode(ptr unsafe.Pointer, iter spi.Iterator)
}

type valDecoderAdapter struct {
	decoder internalDecoder
}

func (decoder *valDecoderAdapter) Decode(val interface{}, iter spi.Iterator) {
	ptr := (*emptyInterface)(unsafe.Pointer(&val)).word
	decoder.decoder.decode(ptr, iter)
}

type internalEncoder interface {
	encode(ptr unsafe.Pointer, stream spi.Stream)
	thriftType() protocol.TType
}

type valEncoderAdapter struct {
	encoder internalEncoder
}

func (encoder *valEncoderAdapter) Encode(val interface{}, stream spi.Stream) {
	ptr := (*emptyInterface)(unsafe.Pointer(&val)).word
	encoder.encoder.encode(ptr, stream)
}

type ptrEncoderAdapter struct {
	encoder internalEncoder
}

func (encoder *ptrEncoderAdapter) Encode(val interface{}, stream spi.Stream) {
	ptr := (*emptyInterface)(unsafe.Pointer(&val)).word
	encoder.encoder.encode(unsafe.Pointer(&ptr), stream)
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

// sliceHeader is a safe version of SliceHeader used within this package.
type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
