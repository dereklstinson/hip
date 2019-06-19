package hip

//#include <hip/hip_runtime_api.h>
import "C"
import (
	"unsafe"
)

type Stream struct {
	s C.hipStream_t
}

func (s *Stream) Ptr() unsafe.Pointer {
	return (unsafe.Pointer)(s)
}
func (s *Stream) Sync() error {
	return status(C.hipStreamSynchronize((s.s))).error("(s *Stream) Sync()")

}
func CreateStream() (s *Stream, err error) {
	s = new(Stream)
	err = status(C.hipStreamCreate((&s.s))).error("CreateStream()")
	return s, err
}

type StreamFlag C.uint

func (s StreamFlag) c() C.uint      { return (C.uint)(s) }
func (s *StreamFlag) cptr() *C.uint { return (*C.uint)(s) }

func (s *StreamFlag) Default() StreamFlag     { *s = (StreamFlag)(C.hipStreamDefault); return *s }
func (s *StreamFlag) NonBlocking() StreamFlag { *s = (StreamFlag)(C.hipStreamNonBlocking); return *s }
