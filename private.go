package hip

import "C"

//func  hipExtGetLinkTypeAndHopCount(int device1, int device2, uint32_t* linktype, uint32_t* hopcount)error{return status(C.hipExtGetLinkTypeAndHopCount(int device1, int device2, uint32_t* linktype, uint32_t* hopcount)).error("hipExtGetLinkTypeAndHopCount")}
//func  hipGetLastError(void)error{return status(C.hipGetLastError(void)).error("hipGetLastError")}
//func  hipPeekAtLastError(void)error{return status(C.hipPeekAtLastError(void)).error("hipPeekAtLastError")}
//func  hipDeviceCanAccessPeer(int* canAccessPeer, int deviceId, int peerDeviceId)error{return status(C.hipDeviceCanAccessPeer(int* canAccessPeer, int deviceId, int peerDeviceId)).error("hipDeviceCanAccessPeer")}
//func  hipDeviceEnablePeerAccess(int peerDeviceId, unsigned int flags)error{return status(C.hipDeviceEnablePeerAccess(int peerDeviceId, unsigned int flags)).error("hipDeviceEnablePeerAccess")}
//func  hipDeviceDisablePeerAccess(int peerDeviceId)error{return status(C.hipDeviceDisablePeerAccess(int peerDeviceId)).error("hipDeviceDisablePeerAccess")}
//func  hipMemGetAddressRange(hipDeviceptr_t* pbase, size_t* psize, hipDeviceptr_t dptr)error{return status(C.hipMemGetAddressRange(hipDeviceptr_t* pbase, size_t* psize, hipDeviceptr_t dptr)).error("hipMemGetAddressRange")}
//func  hipMemcpyPeer(void* dst, int dstDeviceId, const void* src, int srcDeviceId, size_t sizeBytes)error{return status(C.hipMemcpyPeer(void* dst, int dstDeviceId, const void* src, int srcDeviceId, size_t sizeBytes)).error("hipMemcpyPeer")}
//func  hipMemcpyPeerAsync(void* dst, int dstDeviceId, const void* src, int srcDevice,size_t sizeBytes, hipStream_t stream __dparm(0))error{return status(C.hipMemcpyPeerAsync(void* dst, int dstDeviceId, const void* src, int srcDevice,size_t sizeBytes, hipStream_t stream __dparm(0))).error("hipMemcpyPeerAsync")}
//func  hipInit(unsigned int flags)error{return status(C.hipInit(unsigned int flags)).error("hipInit")}
//func  hipCtxCreate(hipCtx_t* ctx, unsigned int flags, hipDevice_t device)error{return status(C.hipCtxCreate(hipCtx_t* ctx, unsigned int flags, hipDevice_t device)).error("hipCtxCreate")}
//func  hipCtxDestroy(hipCtx_t ctx)error{return status(C.hipCtxDestroy(hipCtx_t ctx)).error("hipCtxDestroy")}
//func  hipCtxPopCurrent(hipCtx_t* ctx)error{return status(C.hipCtxPopCurrent(hipCtx_t* ctx)).error("hipCtxPopCurrent")}
//func  hipCtxPushCurrent(hipCtx_t ctx)error{return status(C.hipCtxPushCurrent(hipCtx_t ctx)).error("hipCtxPushCurrent")}
//func  hipCtxSetCurrent(hipCtx_t ctx)error{return status(C.hipCtxSetCurrent(hipCtx_t ctx)).error("hipCtxSetCurrent")}
//func  hipCtxGetCurrent(hipCtx_t* ctx)error{return status(C.hipCtxGetCurrent(hipCtx_t* ctx)).error("hipCtxGetCurrent")}
//func  hipCtxGetDevice(hipDevice_t* device)error{return status(C.hipCtxGetDevice(hipDevice_t* device)).error("hipCtxGetDevice")}
//func  hipCtxGetApiVersion(hipCtx_t ctx, int* apiVersion)error{return status(C.hipCtxGetApiVersion(hipCtx_t ctx, int* apiVersion)).error("hipCtxGetApiVersion")}
//func  hipCtxGetCacheConfig(hipFuncCache_t* cacheConfig)error{return status(C.hipCtxGetCacheConfig(hipFuncCache_t* cacheConfig)).error("hipCtxGetCacheConfig")}
//func  hipCtxSetCacheConfig(hipFuncCache_t cacheConfig)error{return status(C.hipCtxSetCacheConfig(hipFuncCache_t cacheConfig)).error("hipCtxSetCacheConfig")}
//func  hipCtxSetSharedMemConfig(hipSharedMemConfig config)error{return status(C.hipCtxSetSharedMemConfig(hipSharedMemConfig config)).error("hipCtxSetSharedMemConfig")}
//func  hipCtxGetSharedMemConfig(hipSharedMemConfig* pConfig)error{return status(C.hipCtxGetSharedMemConfig(hipSharedMemConfig* pConfig)).error("hipCtxGetSharedMemConfig")}
//func  hipCtxSynchronize(void)error{return status(C.hipCtxSynchronize(void)).error("hipCtxSynchronize")}
//func  hipCtxGetFlags(unsigned int* flags)error{return status(C.hipCtxGetFlags(unsigned int* flags)).error("hipCtxGetFlags")}
//func  hipCtxEnablePeerAccess(hipCtx_t peerCtx, unsigned int flags)error{return status(C.hipCtxEnablePeerAccess(hipCtx_t peerCtx, unsigned int flags)).error("hipCtxEnablePeerAccess")}
//func  hipCtxDisablePeerAccess(hipCtx_t peerCtx)error{return status(C.hipCtxDisablePeerAccess(hipCtx_t peerCtx)).error("hipCtxDisablePeerAccess")}
//func  hipDriverGetVersion(int* driverVersion)error{return status(C.hipDriverGetVersion(int* driverVersion)).error("hipDriverGetVersion")}
//func  hipRuntimeGetVersion(int* runtimeVersion)error{return status(C.hipRuntimeGetVersion(int* runtimeVersion)).error("hipRuntimeGetVersion")}
//func  hipModuleLoad(hipModule_t* module, const char* fname)error{return status(C.hipModuleLoad(hipModule_t* module, const char* fname)).error("hipModuleLoad")}
//func  hipModuleUnload(hipModule_t module)error{return status(C.hipModuleUnload(hipModule_t module)).error("hipModuleUnload")}
//func  hipModuleGetFunction(hipFunction_t* function, hipModule_t module, const char* kname)error{return status(C.hipModuleGetFunction(hipFunction_t* function, hipModule_t module, const char* kname)).error("hipModuleGetFunction")}
//func  hipFuncGetAttributes(struct hipFuncAttributes* attr, const void* func)error{return status(C.hipFuncGetAttributes(struct hipFuncAttributes* attr, const void* func)).error("hipFuncGetAttributes")}
//func  hipModuleGetGlobal(hipDeviceptr_t* dptr, size_t* bytes, hipModule_t hmod, const char* name)error{return status(C.hipModuleGetGlobal(hipDeviceptr_t* dptr, size_t* bytes, hipModule_t hmod, const char* name)).error("hipModuleGetGlobal")}
//func  hipModuleGetTexRef(textureReference** texRef, hipModule_t hmod, const char* name)error{return status(C.hipModuleGetTexRef(textureReference** texRef, hipModule_t hmod, const char* name)).error("hipModuleGetTexRef")}
//func  hipModuleLoadData(hipModule_t* module, const void* image)error{return status(C.hipModuleLoadData(hipModule_t* module, const void* image)).error("hipModuleLoadData")}
//func  hipModuleLoadDataEx(hipModule_t* module, const void* image, unsigned int numOptions,hipJitOption* options, void** optionValues)error{return status(C.hipModuleLoadDataEx(hipModule_t* module, const void* image, unsigned int numOptions,hipJitOption* options, void** optionValues)).error("hipModuleLoadDataEx")}
//func  hipModuleLaunchKernel(hipFunction_t f, unsigned int gridDimX, unsigned int gridDimY,unsigned int gridDimZ, unsigned int blockDimX,unsigned int blockDimY, unsigned int blockDimZ,unsigned int sharedMemBytes, hipStream_t stream,void** kernelParams, void** extra)error{return status(C.hipModuleLaunchKernel(hipFunction_t f, unsigned int gridDimX, unsigned int gridDimY,unsigned int gridDimZ, unsigned int blockDimX,unsigned int blockDimY, unsigned int blockDimZ,unsigned int sharedMemBytes, hipStream_t stream,void** kernelParams, void** extra)).error("hipModuleLaunchKernel")}
//func  hipLaunchCooperativeKernel(const void* f, dim3 gridDim, dim3 blockDimX,void** kernelParams, unsigned int sharedMemBytes,hipStream_t stream)error{return status(C.hipLaunchCooperativeKernel(const void* f, dim3 gridDim, dim3 blockDimX,void** kernelParams, unsigned int sharedMemBytes,hipStream_t stream)).error("hipLaunchCooperativeKernel")}
//func  hipLaunchCooperativeKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)error{return status(C.hipLaunchCooperativeKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)).error("hipLaunchCooperativeKernelMultiDevice")}
//func  hipOccupancyMaxActiveBlocksPerMultiprocessor(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize)error{return status(C.hipOccupancyMaxActiveBlocksPerMultiprocessor(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize)).error("hipOccupancyMaxActiveBlocksPerMultiprocessor")}
//func  hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize, unsigned int flags)error{return status(C.hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize, unsigned int flags)).error("hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags")}
//func  hipExtLaunchMultiKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)error{return status(C.hipExtLaunchMultiKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)).error("hipExtLaunchMultiKernelMultiDevice")}
//func  hipProfilerStart()error{return status(C.hipProfilerStart()).error("hipProfilerStart")}
//func  hipProfilerStop()error{return status(C.hipProfilerStop()).error("hipProfilerStop")}
//func  hipIpcGetMemHandle(hipIpcMemHandle_t* handle, void* devPtr)error{return status(C.hipIpcGetMemHandle(hipIpcMemHandle_t* handle, void* devPtr)).error("hipIpcGetMemHandle")}
//func  hipIpcOpenMemHandle(void** devPtr, hipIpcMemHandle_t handle, unsigned int flags)error{return status(C.hipIpcOpenMemHandle(void** devPtr, hipIpcMemHandle_t handle, unsigned int flags)).error("hipIpcOpenMemHandle")}
//func  hipIpcCloseMemHandle(void* devPtr)error{return status(C.hipIpcCloseMemHandle(void* devPtr)).error("hipIpcCloseMemHandle")}
//func  hipConfigureCall(dim3 gridDim, dim3 blockDim, size_t sharedMem __dparm(0), hipStream_t stream __dparm(0))error{return status(C.hipConfigureCall(dim3 gridDim, dim3 blockDim, size_t sharedMem __dparm(0), hipStream_t stream __dparm(0))).error("hipConfigureCall")}
//func  hipSetupArgument(const void* arg, size_t size, size_t offset)error{return status(C.hipSetupArgument(const void* arg, size_t size, size_t offset)).error("hipSetupArgument")}
//func  hipLaunchByPtr(const void* func)error{return status(C.hipLaunchByPtr(const void* func)).error("hipLaunchByPtr")}
