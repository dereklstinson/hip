package hip

//#include <hip/hip_runtime_api.h>
import "C"
import "errors"

func init() {
	err := status(C.hipInit(0)).error("go: intit C.hipInit")
	if err != nil {
		panic(err)
	}
}

type IpcMemHandle struct {
	i C.hipIpcMemHandle_t
}
type IpcEventHandle struct {
	i C.hipIpcEventHandle_t
}

type Event struct {
	e C.hipEvent_t
}

type status C.hipError_t

type ComputeMode C.uint

func (c *ComputeMode) Default() ComputeMode {
	*c = (ComputeMode)(C.hipComputeModeDefault)
	return *c
}
func (c *ComputeMode) Exclusive() ComputeMode {
	*c = (ComputeMode)(C.hipComputeModeExclusive)
	return *c
}
func (c *ComputeMode) Prohibitive() ComputeMode {
	*c = (ComputeMode)(C.hipComputeModeProhibited)
	return *c
}
func (c *ComputeMode) EsclusiveProcess() ComputeMode {
	*c = (ComputeMode)(C.hipComputeModeExclusiveProcess)
	return *c
}

func (e status) error(comment string) error {
	switch e {

	case (status)(C.hipSuccess):
		return nil
	case (status)(C.hipErrorOutOfMemory):
		return errors.New(comment + " :hipErrorOutOfMemory")
	case (status)(C.hipErrorNotInitialized):
		return errors.New(comment + " :hipErrorNotInitialized")
	case (status)(C.hipErrorDeinitialized):
		return errors.New(comment + " :hipErrorDeinitialized")
	case (status)(C.hipErrorProfilerDisabled):
		return errors.New(comment + " :hipErrorProfilerDisabled")
	case (status)(C.hipErrorProfilerNotInitialized):
		return errors.New(comment + " :hipErrorProfilerNotInitialized")
	case (status)(C.hipErrorProfilerAlreadyStarted):
		return errors.New(comment + " :hipErrorProfilerAlreadyStarted")
	case (status)(C.hipErrorProfilerAlreadyStopped):
		return errors.New(comment + " :hipErrorProfilerAlreadyStopped")
	case (status)(C.hipErrorInsufficientDriver):
		return errors.New(comment + " :hipErrorInsufficientDriver")
	case (status)(C.hipErrorInvalidImage):
		return errors.New(comment + " :hipErrorInvalidImage")
	case (status)(C.hipErrorInvalidContext):
		return errors.New(comment + " :hipErrorInvalidContext")
	case (status)(C.hipErrorContextAlreadyCurrent):
		return errors.New(comment + " :hipErrorContextAlreadyCurrent")
	case (status)(C.hipErrorMapFailed):
		return errors.New(comment + " :hipErrorMapFailed")
	case (status)(C.hipErrorUnmapFailed):
		return errors.New(comment + " :hipErrorUnmapFailed")
	case (status)(C.hipErrorArrayIsMapped):
		return errors.New(comment + " :hipErrorArrayIsMapped")
	case (status)(C.hipErrorAlreadyMapped):
		return errors.New(comment + " :hipErrorAlreadyMapped")
	case (status)(C.hipErrorNoBinaryForGpu):
		return errors.New(comment + " :hipErrorNoBinaryForGpu")
	case (status)(C.hipErrorAlreadyAcquired):
		return errors.New(comment + " :hipErrorAlreadyAcquired")
	case (status)(C.hipErrorNotMapped):
		return errors.New(comment + " :hipErrorNotMapped")
	case (status)(C.hipErrorNotMappedAsArray):
		return errors.New(comment + " :hipErrorNotMappedAsArray")
	case (status)(C.hipErrorNotMappedAsPointer):
		return errors.New(comment + " :hipErrorNotMappedAsPointer")
	case (status)(C.hipErrorECCNotCorrectable):
		return errors.New(comment + " :hipErrorECCNotCorrectable")
	case (status)(C.hipErrorUnsupportedLimit):
		return errors.New(comment + " :hipErrorUnsupportedLimit")
	case (status)(C.hipErrorContextAlreadyInUse):
		return errors.New(comment + " :hipErrorContextAlreadyInUse")
	case (status)(C.hipErrorPeerAccessUnsupported):
		return errors.New(comment + " :hipErrorPeerAccessUnsupported")
	case (status)(C.hipErrorInvalidKernelFile):
		return errors.New(comment + " :hipErrorInvalidKernelFile")
	case (status)(C.hipErrorInvalidGraphicsContext):
		return errors.New(comment + " :hipErrorInvalidGraphicsContext")
	case (status)(C.hipErrorInvalidSource):
		return errors.New(comment + " :hipErrorInvalidSource")
	case (status)(C.hipErrorFileNotFound):
		return errors.New(comment + " :hipErrorFileNotFound")
	case (status)(C.hipErrorSharedObjectSymbolNotFound):
		return errors.New(comment + " :hipErrorSharedObjectSymbolNotFound")
	case (status)(C.hipErrorSharedObjectInitFailed):
		return errors.New(comment + " :hipErrorSharedObjectInitFailed")
	case (status)(C.hipErrorOperatingSystem):
		return errors.New(comment + " :hipErrorOperatingSystem")
	case (status)(C.hipErrorSetOnActiveProcess):
		return errors.New(comment + " :hipErrorSetOnActiveProcess")
	case (status)(C.hipErrorInvalidHandle):
		return errors.New(comment + " :hipErrorInvalidHandle")
	case (status)(C.hipErrorNotFound):
		return errors.New(comment + " :hipErrorNotFound")
	case (status)(C.hipErrorIllegalAddress):
		return errors.New(comment + " :hipErrorIllegalAddress")
	case (status)(C.hipErrorInvalidSymbol):
		return errors.New(comment + " :hipErrorInvalidSymbol")
	case (status)(C.hipErrorMissingConfiguration):
		return errors.New(comment + " :hipErrorMissingConfiguration")
	case (status)(C.hipErrorMemoryAllocation):
		return errors.New(comment + " :hipErrorMemoryAllocation")
	case (status)(C.hipErrorInitializationError):
		return errors.New(comment + " :hipErrorInitializationError")
	case (status)(C.hipErrorLaunchFailure):
		return errors.New(comment + " :hipErrorLaunchFailure")
	case (status)(C.hipErrorPriorLaunchFailure):
		return errors.New(comment + " :hipErrorPriorLaunchFailure")
	case (status)(C.hipErrorLaunchTimeOut):
		return errors.New(comment + " :hipErrorLaunchTimeOut")
	case (status)(C.hipErrorLaunchOutOfResources):
		return errors.New(comment + " :hipErrorLaunchOutOfResources")
	case (status)(C.hipErrorInvalidDeviceFunction):
		return errors.New(comment + " :hipErrorInvalidDeviceFunction")
	case (status)(C.hipErrorInvalidConfiguration):
		return errors.New(comment + " :hipErrorInvalidConfiguration")
	case (status)(C.hipErrorInvalidDevice):
		return errors.New(comment + " :hipErrorInvalidDevice")
	case (status)(C.hipErrorInvalidValue):
		return errors.New(comment + " :hipErrorInvalidValue")
	case (status)(C.hipErrorInvalidDevicePointer):
		return errors.New(comment + " :hipErrorInvalidDevicePointer")
	case (status)(C.hipErrorInvalidMemcpyDirection):
		return errors.New(comment + " :hipErrorInvalidMemcpyDirection")
	case (status)(C.hipErrorUnknown):
		return errors.New(comment + " :hipErrorUnknown")
	case (status)(C.hipErrorInvalidResourceHandle):
		return errors.New(comment + " :hipErrorInvalidResourceHandle")
	case (status)(C.hipErrorNotReady):
		return errors.New(comment + " :hipErrorNotReady")
	case (status)(C.hipErrorNoDevice):
		return errors.New(comment + " :hipErrorNoDevice")
	case (status)(C.hipErrorPeerAccessAlreadyEnabled):
		return errors.New(comment + " :hipErrorPeerAccessAlreadyEnabled")
	case (status)(C.hipErrorPeerAccessNotEnabled):
		return errors.New(comment + " :hipErrorPeerAccessNotEnabled")
	case (status)(C.hipErrorRuntimeMemory):
		return errors.New(comment + " :hipErrorRuntimeMemory")
	case (status)(C.hipErrorRuntimeOther):
		return errors.New(comment + " :hipErrorRuntimeOther")
	case (status)(C.hipErrorHostMemoryAlreadyRegistered):
		return errors.New(comment + " :hipErrorHostMemoryAlreadyRegistered")
	case (status)(C.hipErrorHostMemoryNotRegistered):
		return errors.New(comment + " :hipErrorHostMemoryNotRegistered")
	case (status)(C.hipErrorMapBufferObjectFailed):
		return errors.New(comment + " :hipErrorMapBufferObjectFailed")
	case (status)(C.hipErrorAssert):
		return errors.New(comment + " :hipErrorAssert")

	}
	return errors.New("hipgo: error unknown")
}
func GetLastError() error { return status(C.hipGetLastError()).error("GetLastError") }

type Dim3 C.dim3

func (d Dim3) c() C.dim3 { return (C.dim3)(d) }
func (d *Dim3) Get() (x, y, z uint32) {
	x, y, z = (uint32)(d.x), (uint32)(d.y), (uint32)(d.z)
	return x, y, z
}

func (d *Dim3) Set(x, y, z uint32) {
	d.x, d.y, d.z = (C.uint32_t)(x), (C.uint32_t)(y), (C.uint32_t)(z)
}

/*
typedef struct hipPointerAttribute_t {
    enum hipMemoryType memoryType;
    int device;
    void* devicePointer;
    void* hostPointer;
    int isManaged;
    unsigned allocationFlags;
   	} hipPointerAttribute_t;

//func  hipPointerGetAttributes(hipPointerAttribute_t* attributes, const void* ptr)error{return status(C.hipPointerGetAttributes(hipPointerAttribute_t* attributes, const void* ptr)).error("hipPointerGetAttributes")}



typedef struct hipLaunchParams_t {
    void* func;             ///< Device function symbol
    dim3 gridDim;           ///< Grid dimentions
    dim3 blockDim;          ///< Block dimentions
    void **args;            ///< Arguments
    size_t sharedMem;       ///< Shared memory
    hipStream_t stream;     ///< Stream identifier
} hipLaunchParams;
*/

/*
const char* hipGetErrorName(hipError_t hip_error);
const char* hipGetErrorString(hipError_t hipError);
typedef void (*hipStreamCallback_t)(hipStream_t stream, hipError_t status, void* userData);
hipError_t hipStreamAddCallback(hipStream_t stream, hipStreamCallback_t callback, void* userData, unsigned int flags);
*/
