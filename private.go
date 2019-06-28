package hip

import "C"

//func  hipExtGetLinkTypeAndHopCount(int device1, int device2, uint32_t* linktype, uint32_t* hopcount)error{return status(C.hipExtGetLinkTypeAndHopCount(int device1, int device2, uint32_t* linktype, uint32_t* hopcount)).error("hipExtGetLinkTypeAndHopCount")}

//func  hipPeekAtLastError(void)error{return status(C.hipPeekAtLastError(void)).error("hipPeekAtLastError")}
//func  hipDeviceCanAccessPeer(int* canAccessPeer, int deviceId, int peerDeviceId)error{return status(C.hipDeviceCanAccessPeer(int* canAccessPeer, int deviceId, int peerDeviceId)).error("hipDeviceCanAccessPeer")}
//func  hipDeviceEnablePeerAccess(int peerDeviceId, unsigned int flags)error{return status(C.hipDeviceEnablePeerAccess(int peerDeviceId, unsigned int flags)).error("hipDeviceEnablePeerAccess")}
//func  hipDeviceDisablePeerAccess(int peerDeviceId)error{return status(C.hipDeviceDisablePeerAccess(int peerDeviceId)).error("hipDeviceDisablePeerAccess")}
//func  hipMemGetAddressRange(hipDeviceptr_t* pbase, size_t* psize, hipDeviceptr_t dptr)error{return status(C.hipMemGetAddressRange(hipDeviceptr_t* pbase, size_t* psize, hipDeviceptr_t dptr)).error("hipMemGetAddressRange")}
//func  hipMemcpyPeer(void* dst, int dstDeviceId, const void* src, int srcDeviceId, size_t sizeBytes)error{return status(C.hipMemcpyPeer(void* dst, int dstDeviceId, const void* src, int srcDeviceId, size_t sizeBytes)).error("hipMemcpyPeer")}
//func  hipMemcpyPeerAsync(void* dst, int dstDeviceId, const void* src, int srcDevice,size_t sizeBytes, hipStream_t stream __dparm(0))error{return status(C.hipMemcpyPeerAsync(void* dst, int dstDeviceId, const void* src, int srcDevice,size_t sizeBytes, hipStream_t stream __dparm(0))).error("hipMemcpyPeerAsync")}
//func  hipInit(unsigned int flags)error{return status(C.hipInit(unsigned int flags)).error("hipInit")}

//func  hipDriverGetVersion(int* driverVersion)error{return status(C.hipDriverGetVersion(int* driverVersion)).error("hipDriverGetVersion")}
//func  hipRuntimeGetVersion(int* runtimeVersion)error{return status(C.hipRuntimeGetVersion(int* runtimeVersion)).error("hipRuntimeGetVersion")}

//func  hipProfilerStart()error{return status(C.hipProfilerStart()).error("hipProfilerStart")}
//func  hipProfilerStop()error{return status(C.hipProfilerStop()).error("hipProfilerStop")}
//func  hipIpcGetMemHandle(hipIpcMemHandle_t* handle, void* devPtr)error{return status(C.hipIpcGetMemHandle(hipIpcMemHandle_t* handle, void* devPtr)).error("hipIpcGetMemHandle")}
//func  hipIpcOpenMemHandle(void** devPtr, hipIpcMemHandle_t handle, unsigned int flags)error{return status(C.hipIpcOpenMemHandle(void** devPtr, hipIpcMemHandle_t handle, unsigned int flags)).error("hipIpcOpenMemHandle")}
//func  hipIpcCloseMemHandle(void* devPtr)error{return status(C.hipIpcCloseMemHandle(void* devPtr)).error("hipIpcCloseMemHandle")}
//func  hipConfigureCall(dim3 gridDim, dim3 blockDim, size_t sharedMem __dparm(0), hipStream_t stream __dparm(0))error{return status(C.hipConfigureCall(dim3 gridDim, dim3 blockDim, size_t sharedMem __dparm(0), hipStream_t stream __dparm(0))).error("hipConfigureCall")}
//func  hipSetupArgument(const void* arg, size_t size, size_t offset)error{return status(C.hipSetupArgument(const void* arg, size_t size, size_t offset)).error("hipSetupArgument")}
