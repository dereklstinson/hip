package hip

//#include <hip/hip_runtime_api.h>
import "C"
import "unsafe"
import "github.com/dereklstinson/cutil"

//Device is a device
type Device C.hipDevice_t

func (d Device) c() C.hipDevice_t      { return (C.hipDevice_t)(d) }
func (d *Device) cptr() *C.hipDevice_t { return (*C.hipDevice_t)(d) }
//DeviceSynchronize - Waits on all active streams on current device
//
//When this command is invoked, the host thread gets blocked until all the commands associated
//with streams associated with the device. HIP does not support multiple blocking modes (yet!).
func DeviceSynchronize() error {
	return status(C.hipDeviceSynchronize()).error("hipDeviceSynchronize")
}
//DeviceReset - The state of current device is discarded and updated to a fresh state.
//
//Calling this function deletes all streams created, memory allocated, kernels running, events
//created. Make sure that no other thread is using the device or streams, memory, kernels, events
//associated with the current device.
func DeviceReset() error { return status(C.hipDeviceReset()).error("hipDeviceReset") }
//SetDevice - Set default device to be used for subsequent hip API calls from this thread.
//
//Sets device as the default device for the calling host thread. Valid device id's are 0... (hipGetDeviceCount()-1).
//
//Many HIP APIs implicitly use the "default device" :
//	Any device memory subsequently allocated from this host thread (using hipMalloc) will be allocated on device.
//	Any streams or events created from this host thread will be associated with device.
//	Any kernels launched from this host thread (using hipLaunchKernel) will be executed on device (unless a specific stream is specified, in which case the device associated with that stream will be used).
//
//This function may be called from any host thread. Multiple host threads may use the same device. This function does no synchronization with the previous or new device, and has very little runtime overhead. Applications can use hipSetDevice to quickly switch the default device before making a HIP runtime call which uses the default device.
//
//The default device is stored in thread-local-storage for each thread. Thread-pool implementations may inherit the default device of the previous thread. A good practice is to always call hipSetDevice at the start of HIP coding sequency to establish a known standard device.
func SetDevice(deviceID int32) error {
	return status(C.hipSetDevice((C.int)(deviceID))).error("hipSetDevice")
}

//GetDevice - Return the default device id for the calling host thread. 
//
//HIP maintains an default device for each thread using thread-local-storage. This device is used implicitly for HIP runtime APIs called by this thread.
//GetDevice returns the default device for the calling host thread.
func GetDevice() (deviceID int32, err error) {
	err = status(C.hipGetDevice((*C.int)(&deviceID))).error("hipGetDevice")
	return deviceID, err
}
//GetDeviceCount - Return number of compute-capable devices.
func GetDeviceCount() (count int32, err error) {
	err = status(C.hipGetDeviceCount((*C.int)(&count))).error("hipGetDeviceCount")
	return count, err
}

func DeviceGetAttribute(attr DeviceAttribute, deviceID int32) (val int32, err error) {
	err = status(C.hipDeviceGetAttribute((*C.int)(&val), attr.c(), (C.int)(deviceID))).error("hipDeviceGetAttribute")
	return val, err
}

func GetDeviceProperties(deviceID int32) (prop *DeviceProp, err error) {
	prop = new(DeviceProp)
	err = status(C.hipGetDeviceProperties(prop.cptr(), (C.int)(deviceID))).error("hipGetDeviceProperties")
	return prop, err
}

func DeviceSetCacheConfig(cacheConfig FuncCache) error {
	return status(C.hipDeviceSetCacheConfig(cacheConfig.c())).error("hipDeviceSetCacheConfig")
}

func DeviceGetCacheConfig() (cacheConfig FuncCache, err error) {
	err = status(C.hipDeviceGetCacheConfig(cacheConfig.cptr())).error("hipDeviceGetCacheConfig")
	return cacheConfig, err
}

func DeviceGetLimit(limit Limit) (pValue uint, err error) {
	var sizet C.size_t
	err = status(C.hipDeviceGetLimit(&sizet, limit.c())).error("hipDeviceGetLimit")
	pValue = (uint)(sizet)
	return pValue, err
}

//for hipgo this has to be used with .so functions
func FuncSetCacheConfig(sofunc unsafe.Pointer, config FuncCache) error {
	return status(C.hipFuncSetCacheConfig(sofunc, config.c())).error("FuncSetCacheConfig")
}
func DeviceGetSharedMemConfig() (pConfig SharedMemConfig, err error) {
	err = status(C.hipDeviceGetSharedMemConfig(pConfig.cptr())).error("DeviceGetSharedMemConfig")
	return pConfig, err
}
func DeviceSetSharedMemConfig(pConfig SharedMemConfig) error {
	return status(C.hipDeviceSetSharedMemConfig(pConfig.c())).error("DeviceSetSharedMemConfig")
}
func SetDeviceFlags(dflag DeviceFlag) error {
	return status(C.hipSetDeviceFlags(dflag.c())).error("SetDeviceFlags")
}

func ChooseDevice(prop *DeviceProp) (device int32, err error) {
	err = status(C.hipChooseDevice((*C.int)(&device), prop.cptr())).error("ChooseDevice")
	return device, err
}

func DeviceGetPCIBusId(device int32) (pciBusId string, err error) {
	length := C.int(255)
	bid := make([]C.char, length)
	err = status(C.hipDeviceGetPCIBusId(&bid[0], length, (C.int)(device))).error("DeviceGetPCIBusId")
	pciBusId = C.GoString(&bid[0])
	return pciBusId, err

}

func DeviceGetByPCIBusId(pciBusId string) (device int32, err error) {
	bid := []byte(pciBusId)
	cbid := make([]C.char, len(bid))
	for i := range bid {
		cbid[i] = (C.char)(bid[i])
	}
	err = status(C.hipDeviceGetByPCIBusId((*C.int)(&device), &cbid[0])).error("DeviceGetByPCIBusId")
	return device, err
}

func DeviceGet(ordinal int32) (d Device, err error) {
	err = status(C.hipDeviceGet(d.cptr(), (C.int)(ordinal))).error("DeviceGet")
	return d, err
}
func (d Device) ComputeCapability() (major, minor int32, err error) {
	err = status(C.hipDeviceComputeCapability((*C.int)(&major), (*C.int)(&minor), d.c())).error("(d Device) ComputeCapability()")
	return major, minor, err
}
func (d Device) GetName() (name string, err error) {
	length := C.int(255)
	cname := make([]C.char, length)
	err = status(C.hipDeviceGetName(&cname[0], length, d.c())).error("(d Device) GetName()")
	name = C.GoString(&cname[0])
	return name, err
}
func (d Device) GetTotalMem() (totalSIB uint, err error) {
	var sib C.size_t
	err = status(C.hipDeviceTotalMem(&sib, d.c())).error("(d Device) GetTotalMem()")
	totalSIB = (uint)(sib)
	return totalSIB, err
}

func (d Device) PrimaryCtxGetState() (flags uint32, active bool, err error) {
	var act C.int
	err = status(C.hipDevicePrimaryCtxGetState(d.c(), (*C.uint)(&flags), &act)).error("(d Device) PrimaryCtxGetState()")
	if act > 0 {
		active = true
	}
	return flags, active, err
}
func (d Device) PrimaryCtxRelease() error {
	return status(C.hipDevicePrimaryCtxRelease(d.c())).error("(d Device)PrimaryCtxRelease()")
}
func (d Device) PrimaryCtxRetain() (pctx *Context, err error) {
	pctx = new(Context)
	err = status(C.hipDevicePrimaryCtxRetain(&pctx.c, d.c())).error("hipDevicePrimaryCtxRetain")
	return pctx, err
}
func (d Device) PrimaryCtxReset() error {
	return status(C.hipDevicePrimaryCtxReset(d.c())).error("hipDevicePrimaryCtxReset")
}

func (d Device) PrimaryCtxSetFlags(flags uint32) error {
	return status(C.hipDevicePrimaryCtxSetFlags(d.c(), (C.uint)(flags))).error("hipDevicePrimaryCtxSetFlags")
}

func DeviceCanAccessPeer(deviceId, peerDeviceId int32) (bool, error) {
	var canAccessPeer C.int
	err := status(C.hipDeviceCanAccessPeer(&canAccessPeer, (C.int)(deviceId), (C.int)(peerDeviceId))).error("DeviceCanAccessPeer")
	if canAccessPeer > 0 {
		return true, err
	}
	return false, err
}

//PeerFlag is a holder for flag used in DeviceEnablePeer Access
type PeerFlag C.uint

//Default is the default flag
func (p *PeerFlag) Default() PeerFlag { *p = 0; return *p }
func (p PeerFlag) c() C.uint          { return (C.uint)(p) }
func DeviceEnablePeerAccess(peerDeviceId int32, p PeerFlag) error {
	return status(C.hipDeviceEnablePeerAccess((C.int)(peerDeviceId), p.c())).error("hipDeviceEnablePeerAccess")
}

func DeviceDisablePeerAccess(peerDeviceId int) error {
	return status(C.hipDeviceDisablePeerAccess((C.int)(peerDeviceId))).error("hipDeviceDisablePeerAccess")
}

func MemcpyPeer(dst cutil.Mem, dstDeviceId int32, src cutil.Mem, srcDeviceId int32, sib uint) error {
	return status(C.hipMemcpyPeer(dst.Ptr(), (C.int)(dstDeviceId), src.Ptr(), (C.int)(srcDeviceId), (C.size_t)(sib))).error("hipMemcpyPeer")
}

func MemcpyPeerAsync(dst cutil.Mem, dstDeviceId int32, src cutil.Mem, srcDeviceId int32, sib uint, s *Stream) error {
	return status(C.hipMemcpyPeerAsync(dst.Ptr(), (C.int)(dstDeviceId), src.Ptr(), (C.int)(srcDeviceId), (C.size_t)(sib), s.s)).error("hipMemcpyPeer")
}
