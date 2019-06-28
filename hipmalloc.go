package hip

//#include <hip/hip_runtime_api.h>
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/dereklstinson/cutil"
)

type hipmem struct {
	m unsafe.Pointer
}

//DevicePtr is a pointer to device mem
type DevicePtr struct {
	d C.hipDeviceptr_t
}

func (d *DevicePtr) Ptr() unsafe.Pointer {
	return (unsafe.Pointer)(d.d)
}
func (d *DevicePtr) DPtr() *unsafe.Pointer {
	return (*unsafe.Pointer)(&d.d)
}

type HipArrayDescriptor C.HIP_ARRAY_DESCRIPTOR

/*typedef struct HIP_ARRAY_DESCRIPTOR {
    enum hipArray_Format format;
    unsigned int numChannels;
    size_t width;
    size_t height;
    unsigned int flags;
    size_t depth;
}HIP_ARRAY_DESCRIPTOR;
*/
func (h *hipmem) Ptr() unsafe.Pointer {
	return h.m
}
func (h *hipmem) DPtr() *unsafe.Pointer {
	return &h.m
}
func Malloc(mem cutil.Mem, sib uint) error {
	sizet := (C.size_t)(sib)
	err := status(C.hipMalloc(mem.DPtr(), sizet)).error("Malloc")
	if err != nil {
		return err
	}
	runtime.SetFinalizer(mem, hipFree)
	return err
}

func ExtMallocWithFlags(mem cutil.Mem, sib uint, flags MallocFlags) error {
	sizet := (C.size_t)(sib)
	err := status(C.hipExtMallocWithFlags(mem.DPtr(), sizet, flags.c())).error("ExtMallocWithFlags")
	if err != nil {
		return err
	}
	runtime.SetFinalizer(mem, hipFree)
	return err
}

func HostMalloc(mem cutil.Mem, sib uint, flags MallocFlags) error {
	sizet := (C.size_t)(sib)
	err := status(C.hipHostMalloc(mem.DPtr(), sizet, flags.c())).error("HostMalloc")
	if err != nil {
		return err
	}
	runtime.SetFinalizer(mem, hipHostFree)
	return err
}

func MallocManaged(mem cutil.Mem, sib uint, flags MallocFlags) error {
	sizet := (C.size_t)(sib)
	err := status(C.hipMallocManaged(mem.DPtr(), sizet, flags.c())).error("MallocManaged")
	if err != nil {
		return err
	}
	runtime.SetFinalizer(mem, hipFree)
	return err
}

func HostGetDevicePointer(hostmem cutil.Mem, flags MallocFlags) (devicemem cutil.Mem, err error) {
	devicemem = new(hipmem)
	err = status(C.hipHostGetDevicePointer(devicemem.DPtr(), hostmem.Ptr(), flags.c())).error("HostGetDevicePointer")
	return devicemem, err
}

func HostGetFlags(hostmem cutil.Mem) (flags MallocFlags, err error) {
	err = status(C.hipHostGetFlags(flags.cptr(), hostmem.Ptr())).error("HostGetFlags")
	return flags, err
}

func Memcpy(dst, src cutil.Mem, sib uint, kind MemCpyKind) error {
	sizet := (C.size_t)(sib)
	return status(C.hipMemcpy(dst.Ptr(), src.Ptr(), sizet, kind.c())).error("Memcpy")
}
func MemcpyAsync(dst, src cutil.Mem, sib uint, kind MemCpyKind, stream Stream) error {
	sizet := (C.size_t)(sib)
	return status(C.hipMemcpyAsync(dst.Ptr(), src.Ptr(), sizet, kind.c(), stream.s)).error("MemcpyAsync")
}
func MemcpyHtoD(dst *DevicePtr, src cutil.Mem, sib uint) error {
	sizet := (C.size_t)(sib)
	return status(C.hipMemcpyHtoD(dst.d, src.Ptr(), sizet)).error("MemcpyHtoD")
}

func MemcpyDtoH(dst cutil.Mem, src *DevicePtr, sib uint) error {
	sizet := (C.size_t)(sib)
	return status(C.hipMemcpyDtoH(dst.Ptr(), src.d, sizet)).error("MemcpyDtoH")
}

func MemcpyDtoD(dst, src *DevicePtr, sib uint) error {
	sizet := (C.size_t)(sib)
	return status(C.hipMemcpyDtoD(dst.d, src.d, sizet)).error("MemcpyDtoD")
}

func MemcpyHtoDAsync(dst *DevicePtr, src cutil.Mem, sib uint, stream Stream) error {
	sizet := (C.size_t)(sib)
	return status(C.hipMemcpyHtoDAsync(dst.d, src.Ptr(), sizet, stream.s)).error("hipMemcpyHtoDAsync")
}
func MemcpyDtoHAsync(dst cutil.Mem, src *DevicePtr, sib uint, stream Stream) error {
	sizet := (C.size_t)(sib)
	return status(C.hipMemcpyDtoHAsync(dst.Ptr(), src.d, sizet, stream.s)).error("hipMemcpyDtoHAsync")
}
func MemcpyDtoDAsync(dst, src *DevicePtr, sib uint, stream Stream) error {
	sizet := (C.size_t)(sib)
	return status(C.hipMemcpyDtoDAsync(dst.d, src.d, sizet, stream.s)).error("hipMemcpyDtoDAsync")
}

func MallocPitch(ptr cutil.Mem, width, height uint) (pitch uint, err error) {
	w := (C.size_t)(width)
	h := (C.size_t)(height)
	var p C.size_t
	err = status(C.hipMallocPitch(ptr.DPtr(), &p, w, h)).error("hipMallocPitch")
	pitch = (uint)(p)
	return pitch, err
}

//func  ModuleGetGlobal(hipDeviceptr_t* dptr, size_t* bytes, hipModule_t hmod, const char* name)error{return status(C.hipModuleGetGlobal(hipDeviceptr_t* dptr, size_t* bytes, hipModule_t hmod, const char* name)).error("hipModuleGetGlobal")}
//func  GetSymbolAddress(void** devPtr, const void* symbolName)error{return status(C.hipGetSymbolAddress(void** devPtr, const void* symbolName)).error("hipGetSymbolAddress")}
//func  GetSymbolSize(size_t* size, const void* symbolName)error{return status(C.hipGetSymbolSize(size_t* size, const void* symbolName)).error("hipGetSymbolSize")}
//func  MemcpyToSymbol(const void* symbolName, const void* src, size_t sizeBytes, size_t offset __dparm(0),hipMemcpyKind kind __dparm(hipMemcpyHostToDevice))error{return status(C.hipMemcpyToSymbol(const void* symbolName, const void* src, size_t sizeBytes, size_t offset __dparm(0),hipMemcpyKind kind __dparm(hipMemcpyHostToDevice))).error("hipMemcpyToSymbol")}
//func  MemcpyToSymbolAsync(const void* symbolName, const void* src,size_t sizeBytes, size_t offset,hipMemcpyKind kind, hipStream_t stream __dparm(0))error{return status(C.hipMemcpyToSymbolAsync(const void* symbolName, const void* src,size_t sizeBytes, size_t offset,hipMemcpyKind kind, hipStream_t stream __dparm(0))).error("hipMemcpyToSymbolAsync")}
//func  MemcpyFromSymbol(void* dst, const void* symbolName,size_t sizeBytes, size_t offset __dparm(0),hipMemcpyKind kind __dparm(hipMemcpyDeviceToHost))error{return status(C.hipMemcpyFromSymbol(void* dst, const void* symbolName,size_t sizeBytes, size_t offset __dparm(0),hipMemcpyKind kind __dparm(hipMemcpyDeviceToHost))).error("hipMemcpyFromSymbol")}
//func  MemcpyFromSymbolAsync(void* dst, const void* symbolName,size_t sizeBytes, size_t offset,hipMemcpyKind kind,hipStream_t stream __dparm(0))error{return status(C.hipMemcpyFromSymbolAsync(void* dst, const void* symbolName,size_t sizeBytes, size_t offset,hipMemcpyKind kind,hipStream_t stream __dparm(0))).error("hipMemcpyFromSymbolAsync")}

//func  Memset(void* dst, int value, size_t sizeBytes)error{return status(C.hipMemset(void* dst, int value, size_t sizeBytes)).error("hipMemset")}
//func  MemsetD8(hipDeviceptr_t dest, unsigned char value, size_t sizeBytes)error{return status(C.hipMemsetD8(hipDeviceptr_t dest, unsigned char value, size_t sizeBytes)).error("hipMemsetD8")}
//func  MemsetD32(hipDeviceptr_t dest, int value, size_t count)error{return status(C.hipMemsetD32(hipDeviceptr_t dest, int value, size_t count)).error("hipMemsetD32")}
//func  MemsetAsync(void* dst, int value, size_t sizeBytes, hipStream_t stream __dparm(0))error{return status(C.hipMemsetAsync(void* dst, int value, size_t sizeBytes, hipStream_t stream __dparm(0))).error("hipMemsetAsync")}
//func  MemsetD32Async(hipDeviceptr_t dst, int value, size_t count, hipStream_t stream __dparm(0))error{return status(C.hipMemsetD32Async(hipDeviceptr_t dst, int value, size_t count, hipStream_t stream __dparm(0))).error("hipMemsetD32Async")}
//func  Memset2D(void* dst, size_t pitch, int value, size_t width, size_t height)error{return status(C.hipMemset2D(void* dst, size_t pitch, int value, size_t width, size_t height)).error("hipMemset2D")}
//func  Memset2DAsync(void* dst, size_t pitch, int value, size_t width, size_t height,hipStream_t stream __dparm(0))error{return status(C.hipMemset2DAsync(void* dst, size_t pitch, int value, size_t width, size_t height,hipStream_t stream __dparm(0))).error("hipMemset2DAsync")}
//func  Memset3D(hipPitchedPtr pitchedDevPtr, int  value, hipExtent extent )error{return status(C.hipMemset3D(hipPitchedPtr pitchedDevPtr, int  value, hipExtent extent )).error("hipMemset3D")}
//func  Memset3DAsync(hipPitchedPtr pitchedDevPtr, int  value, hipExtent extent ,hipStream_t stream __dparm(0))error{return status(C.hipMemset3DAsync(hipPitchedPtr pitchedDevPtr, int  value, hipExtent extent ,hipStream_t stream __dparm(0))).error("hipMemset3DAsync")}
//func  MemGetInfo(size_t* free, size_t* total)error{return status(C.hipMemGetInfo(size_t* free, size_t* total)).error("hipMemGetInfo")}
//func  MemPtrGetInfo(void* ptr, size_t* size)error{return status(C.hipMemPtrGetInfo(void* ptr, size_t* size)).error("hipMemPtrGetInfo")}
//func  MallocArray(hipArray** array, const hipChannelFormatDesc* desc, size_t width, size_t height __dparm(0), unsigned int flags __dparm(hipArrayDefault))error{return status(C.hipMallocArray(hipArray** array, const hipChannelFormatDesc* desc, size_t width, size_t height __dparm(0), unsigned int flags __dparm(hipArrayDefault))).error("hipMallocArray")}
//func  ArrayCreate(hipArray** pHandle, const HIP_ARRAY_DESCRIPTOR* pAllocateArray)error{return status(C.hipArrayCreate(hipArray** pHandle, const HIP_ARRAY_DESCRIPTOR* pAllocateArray)).error("hipArrayCreate")}
//func  Array3DCreate(hipArray** array, const HIP_ARRAY_DESCRIPTOR* pAllocateArray)error{return status(C.hipArray3DCreate(hipArray** array, const HIP_ARRAY_DESCRIPTOR* pAllocateArray)).error("hipArray3DCreate")}
//func  Malloc3D(hipPitchedPtr* pitchedDevPtr, hipExtent extent)error{return status(C.hipMalloc3D(hipPitchedPtr* pitchedDevPtr, hipExtent extent)).error("hipMalloc3D")}
//func  FreeArray(hipArray* array)error{return status(C.hipFreeArray(hipArray* array)).error("hipFreeArray")}
//func  Malloc3DArray(hipArray** array, const struct hipChannelFormatDesc* desc, struct hipExtent extent, unsigned int flags)error{return status(C.hipMalloc3DArray(hipArray** array, const struct hipChannelFormatDesc* desc, struct hipExtent extent, unsigned int flags)).error("hipMalloc3DArray")}
//func  Memcpy2D(void* dst, size_t dpitch, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind)error{return status(C.hipMemcpy2D(void* dst, size_t dpitch, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind)).error("hipMemcpy2D")}
//func  MemcpyParam2D(const hip_Memcpy2D* pCopy)error{return status(C.hipMemcpyParam2D(const hip_Memcpy2D* pCopy)).error("hipMemcpyParam2D")}
//func  Memcpy2DAsync(void* dst, size_t dpitch, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind, hipStream_t stream __dparm(0))error{return status(C.hipMemcpy2DAsync(void* dst, size_t dpitch, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind, hipStream_t stream __dparm(0))).error("hipMemcpy2DAsync")}
//func  Memcpy2DToArray(hipArray* dst, size_t wOffset, size_t hOffset, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind)error{return status(C.hipMemcpy2DToArray(hipArray* dst, size_t wOffset, size_t hOffset, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind)).error("hipMemcpy2DToArray")}
//func  MemcpyToArray(hipArray* dst, size_t wOffset, size_t hOffset, const void* src, size_t count, hipMemcpyKind kind)error{return status(C.hipMemcpyToArray(hipArray* dst, size_t wOffset, size_t hOffset, const void* src, size_t count, hipMemcpyKind kind)).error("hipMemcpyToArray")}
//func  MemcpyFromArray(void* dst, hipArray_const_t srcArray, size_t wOffset, size_t hOffset, size_t count, hipMemcpyKind kind)error{return status(C.hipMemcpyFromArray(void* dst, hipArray_const_t srcArray, size_t wOffset, size_t hOffset, size_t count, hipMemcpyKind kind)).error("hipMemcpyFromArray")}
//func  MemcpyAtoH(void* dst, hipArray* srcArray, size_t srcOffset, size_t count)error{return status(C.hipMemcpyAtoH(void* dst, hipArray* srcArray, size_t srcOffset, size_t count)).error("hipMemcpyAtoH")}
//func  MemcpyHtoA(hipArray* dstArray, size_t dstOffset, const void* srcHost, size_t count)error{return status(C.hipMemcpyHtoA(hipArray* dstArray, size_t dstOffset, const void* srcHost, size_t count)).error("hipMemcpyHtoA")}
//func  Memcpy3D(const struct hipMemcpy3DParms* p)error{return status(C.hipMemcpy3D(const struct hipMemcpy3DParms* p)).error("hipMemcpy3D")}
func hipFree(mem cutil.Mem) error {
	return status(C.hipFree(mem.Ptr())).error("hipFree (hidden)")
}
func hipHostFree(mem cutil.Mem) error {
	return status(C.hipHostFree(mem.Ptr())).error("hipHostFree (hidden)")
}

type MemCpyKind C.hipMemcpyKind

func (m MemCpyKind) c() C.hipMemcpyKind      { return (C.hipMemcpyKind)(m) }
func (m *MemCpyKind) cptr() *C.hipMemcpyKind { return (*C.hipMemcpyKind)(m) }
func (m *MemCpyKind) HtoH() MemCpyKind {
	*m = (MemCpyKind)(C.hipMemcpyHostToHost)
	return *m
}
func (m *MemCpyKind) HtoD() MemCpyKind {
	*m = (MemCpyKind)(C.hipMemcpyHostToDevice)
	return *m
}
func (m *MemCpyKind) DtoH() MemCpyKind {
	*m = (MemCpyKind)(C.hipMemcpyDeviceToHost)
	return *m
}
func (m *MemCpyKind) DtoD() MemCpyKind {
	*m = (MemCpyKind)(C.hipMemcpyDeviceToDevice)
	return *m
}
func (m *MemCpyKind) Default() MemCpyKind {
	*m = (MemCpyKind)(C.hipMemcpyDefault)
	return *m
}

type MallocFlags C.uint

func (m MallocFlags) c() C.uint      { return (C.uint)(m) }
func (m *MallocFlags) cptr() *C.uint { return (*C.uint)(m) }
func (m *MallocFlags) Default() MallocFlags {
	*m = (C.hipHostMallocDefault)
	return *m
}
func (m *MallocFlags) Portable() MallocFlags {
	*m = (C.hipHostMallocPortable)
	return *m
}

func (m *MallocFlags) Mapped() MallocFlags {
	*m = (C.hipHostMallocMapped)
	return *m
}

func (m *MallocFlags) WriteCombined() MallocFlags {
	*m = (C.hipHostMallocWriteCombined)
	return *m
}

func (m *MallocFlags) Coherent() MallocFlags {
	*m = (C.hipHostMallocCoherent)
	return *m
}

func (m *MallocFlags) NonCoherent() MallocFlags {
	*m = (C.hipHostMallocNonCoherent)
	return *m
}

func (m *MallocFlags) Global() MallocFlags {
	*m = (C.hipMemAttachGlobal)
	return *m
}

func (m *MallocFlags) AttachHost() MallocFlags {
	*m = (C.hipMemAttachHost)
	return *m
}

func (m *MallocFlags) DeviceDefault() MallocFlags {
	*m = (C.hipDeviceMallocDefault)
	return *m
}

func (m *MallocFlags) DeviceFinegrained() MallocFlags {
	*m = (C.hipDeviceMallocFinegrained)
	return *m
}
