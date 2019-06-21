package hip

//#include <hip/hip_runtime_api.h>
import "C"
import (
	"runtime"
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
func (s *Stream) Query() error {
	return status(C.hipStreamQuery(s.s)).error("(s *Stream)Query()")
}
func (s *Stream) WaitEvent(event *Event, flags int32) error {
	return status(C.hipStreamWaitEvent(s.s, event.e, 0)).error("(s *Stream) WaitEvent()")
}
func (s *Stream) GetFlags() (flag StreamFlag, err error) {
	err = status(C.hipStreamGetFlags(s.s, flag.cptr())).error("(s *Stream) GetFlags()")
	return flag, err
}
func (s *Stream) GetPriority() (priority int32, err error) {
	err = status(C.hipStreamGetPriority(s.s, (*C.int)(&priority))).error(" (s *Stream)GetPriority()")
	return priority, err

}

func CreateStream() (s *Stream, err error) {
	s = new(Stream)
	err = status(C.hipStreamCreate((&s.s))).error("CreateStream")
	runtime.SetFinalizer(s, hipStreamDestroy)
	return s, err
}

func CreateStreamWithFlags(flag StreamFlag) (s *Stream, err error) {
	s = new(Stream)
	err = status(C.hipStreamCreateWithFlags(&s.s, flag.c())).error("CreateStreamWithFlags")
	runtime.SetFinalizer(s, hipStreamDestroy)
	return s, err
}

func CreateStreamWithPriority(flag StreamFlag, priority int32) (s *Stream, err error) {
	s = new(Stream)
	err = status(C.hipStreamCreateWithPriority(&s.s, flag.c(), (C.int)(priority))).error("CreateStreamWithPriority")
	runtime.SetFinalizer(s, hipStreamDestroy)
	return s, err
}
func hipStreamDestroy(s *Stream) error {
	return status(C.hipStreamDestroy(s.s)).error("hipStreamDestroy(hidden)")
}

func DeviceGetStreamPriorityRange() (least, greatest int32, err error) {
	err = status(C.hipDeviceGetStreamPriorityRange((*C.int)(&least), (*C.int)(&greatest))).error("DeviceGetStreamPriorityRange")
	return least, greatest, err
}

type StreamFlag C.uint

func (s StreamFlag) c() C.uint      { return (C.uint)(s) }
func (s *StreamFlag) cptr() *C.uint { return (*C.uint)(s) }

func (s *StreamFlag) Default() StreamFlag     { *s = (StreamFlag)(C.hipStreamDefault); return *s }
func (s *StreamFlag) NonBlocking() StreamFlag { *s = (StreamFlag)(C.hipStreamNonBlocking); return *s }
