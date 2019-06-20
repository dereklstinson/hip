package hip

import "C"

//func  hipExtGetLinkTypeAndHopCount(int device1, int device2, uint32_t* linktype, uint32_t* hopcount)error{return status(C.hipExtGetLinkTypeAndHopCount(int device1, int device2, uint32_t* linktype, uint32_t* hopcount)).error("hipExtGetLinkTypeAndHopCount")}
//func  hipGetLastError(void)error{return status(C.hipGetLastError(void)).error("hipGetLastError")}
//func  hipPeekAtLastError(void)error{return status(C.hipPeekAtLastError(void)).error("hipPeekAtLastError")}
//func  hipStreamCreate(hipStream_t* stream)error{return status(C.hipStreamCreate(hipStream_t* stream)).error("hipStreamCreate")}
//func  hipStreamCreateWithFlags(hipStream_t* stream, unsigned int flags)error{return status(C.hipStreamCreateWithFlags(hipStream_t* stream, unsigned int flags)).error("hipStreamCreateWithFlags")}
//func  hipStreamCreateWithPriority(hipStream_t* stream, unsigned int flags, int priority)error{return status(C.hipStreamCreateWithPriority(hipStream_t* stream, unsigned int flags, int priority)).error("hipStreamCreateWithPriority")}
//func  hipDeviceGetStreamPriorityRange(int* leastPriority, int* greatestPriority)error{return status(C.hipDeviceGetStreamPriorityRange(int* leastPriority, int* greatestPriority)).error("hipDeviceGetStreamPriorityRange")}
//func  hipStreamDestroy(hipStream_t stream)error{return status(C.hipStreamDestroy(hipStream_t stream)).error("hipStreamDestroy")}
//func  hipStreamQuery(hipStream_t stream)error{return status(C.hipStreamQuery(hipStream_t stream)).error("hipStreamQuery")}
//func  hipStreamSynchronize(hipStream_t stream)error{return status(C.hipStreamSynchronize(hipStream_t stream)).error("hipStreamSynchronize")}
//func  hipStreamWaitEvent(hipStream_t stream, hipEvent_t event, unsigned int flags)error{return status(C.hipStreamWaitEvent(hipStream_t stream, hipEvent_t event, unsigned int flags)).error("hipStreamWaitEvent")}
//func  hipStreamGetFlags(hipStream_t stream, unsigned int* flags)error{return status(C.hipStreamGetFlags(hipStream_t stream, unsigned int* flags)).error("hipStreamGetFlags")}
//func  hipStreamGetPriority(hipStream_t stream, int* priority)error{return status(C.hipStreamGetPriority(hipStream_t stream, int* priority)).error("hipStreamGetPriority")}
//func  hipEventCreateWithFlags(hipEvent_t* event, unsigned flags)error{return status(C.hipEventCreateWithFlags(hipEvent_t* event, unsigned flags)).error("hipEventCreateWithFlags")}
//func  hipEventCreate(hipEvent_t* event)error{return status(C.hipEventCreate(hipEvent_t* event)).error("hipEventCreate")}
//func  hipEventRecord(hipEvent_t event, hipStream_t stream)error{return status(C.hipEventRecord(hipEvent_t event, hipStream_t stream)).error("hipEventRecord")}
//func  hipEventDestroy(hipEvent_t event)error{return status(C.hipEventDestroy(hipEvent_t event)).error("hipEventDestroy")}
//func  hipEventSynchronize(hipEvent_t event)error{return status(C.hipEventSynchronize(hipEvent_t event)).error("hipEventSynchronize")}
//func  hipEventElapsedTime(float* ms, hipEvent_t start, hipEvent_t stop)error{return status(C.hipEventElapsedTime(float* ms, hipEvent_t start, hipEvent_t stop)).error("hipEventElapsedTime")}
//func  hipEventQuery(hipEvent_t event)error{return status(C.hipEventQuery(hipEvent_t event)).error("hipEventQuery")}
//func  hipPointerGetAttributes(hipPointerAttribute_t* attributes, const void* ptr)error{return status(C.hipPointerGetAttributes(hipPointerAttribute_t* attributes, const void* ptr)).error("hipPointerGetAttributes")}
//func  hipMalloc(void** ptr, size_t size)error{return status(C.hipMalloc(void** ptr, size_t size)).error("hipMalloc")}
//func  hipExtMallocWithFlags(void** ptr, size_t sizeBytes, unsigned int flags)error{return status(C.hipExtMallocWithFlags(void** ptr, size_t sizeBytes, unsigned int flags)).error("hipExtMallocWithFlags")}
//func  hipHostMalloc(void** ptr, size_t size, unsigned int flags)error{return status(C.hipHostMalloc(void** ptr, size_t size, unsigned int flags)).error("hipHostMalloc")}
//func  hipMallocManaged(void** devPtr, size_t size, unsigned int flags __dparm(0))error{return status(C.hipMallocManaged(void** devPtr, size_t size, unsigned int flags __dparm(0))).error("hipMallocManaged")}
//func  hipHostGetDevicePointer(void** devPtr, void* hstPtr, unsigned int flags)error{return status(C.hipHostGetDevicePointer(void** devPtr, void* hstPtr, unsigned int flags)).error("hipHostGetDevicePointer")}
//func  hipHostGetFlags(unsigned int* flagsPtr, void* hostPtr)error{return status(C.hipHostGetFlags(unsigned int* flagsPtr, void* hostPtr)).error("hipHostGetFlags")}
//func  hipHostRegister(void* hostPtr, size_t sizeBytes, unsigned int flags)error{return status(C.hipHostRegister(void* hostPtr, size_t sizeBytes, unsigned int flags)).error("hipHostRegister")}
//func  hipHostUnregister(void* hostPtr)error{return status(C.hipHostUnregister(void* hostPtr)).error("hipHostUnregister")}
//func  hipMallocPitch(void** ptr, size_t* pitch, size_t width, size_t height)error{return status(C.hipMallocPitch(void** ptr, size_t* pitch, size_t width, size_t height)).error("hipMallocPitch")}
//func  hipFree(void* ptr)error{return status(C.hipFree(void* ptr)).error("hipFree")}
//func  hipHostFree(void* ptr)error{return status(C.hipHostFree(void* ptr)).error("hipHostFree")}
//func  hipMemcpy(void* dst, const void* src, size_t sizeBytes, hipMemcpyKind kind)error{return status(C.hipMemcpy(void* dst, const void* src, size_t sizeBytes, hipMemcpyKind kind)).error("hipMemcpy")}
//func  hipMemcpyHtoD(hipDeviceptr_t dst, void* src, size_t sizeBytes)error{return status(C.hipMemcpyHtoD(hipDeviceptr_t dst, void* src, size_t sizeBytes)).error("hipMemcpyHtoD")}
//func  hipMemcpyDtoH(void* dst, hipDeviceptr_t src, size_t sizeBytes)error{return status(C.hipMemcpyDtoH(void* dst, hipDeviceptr_t src, size_t sizeBytes)).error("hipMemcpyDtoH")}
//func  hipMemcpyDtoD(hipDeviceptr_t dst, hipDeviceptr_t src, size_t sizeBytes)error{return status(C.hipMemcpyDtoD(hipDeviceptr_t dst, hipDeviceptr_t src, size_t sizeBytes)).error("hipMemcpyDtoD")}
//func  hipMemcpyHtoDAsync(hipDeviceptr_t dst, void* src, size_t sizeBytes, hipStream_t stream)error{return status(C.hipMemcpyHtoDAsync(hipDeviceptr_t dst, void* src, size_t sizeBytes, hipStream_t stream)).error("hipMemcpyHtoDAsync")}
//func  hipMemcpyDtoHAsync(void* dst, hipDeviceptr_t src, size_t sizeBytes, hipStream_t stream)error{return status(C.hipMemcpyDtoHAsync(void* dst, hipDeviceptr_t src, size_t sizeBytes, hipStream_t stream)).error("hipMemcpyDtoHAsync")}
//func  hipMemcpyDtoDAsync(hipDeviceptr_t dst, hipDeviceptr_t src, size_t sizeBytes, hipStream_t stream)error{return status(C.hipMemcpyDtoDAsync(hipDeviceptr_t dst, hipDeviceptr_t src, size_t sizeBytes, hipStream_t stream)).error("hipMemcpyDtoDAsync")}
//func  hipModuleGetGlobal(hipDeviceptr_t* dptr, size_t* bytes, hipModule_t hmod, const char* name)error{return status(C.hipModuleGetGlobal(hipDeviceptr_t* dptr, size_t* bytes, hipModule_t hmod, const char* name)).error("hipModuleGetGlobal")}
//func  hipGetSymbolAddress(void** devPtr, const void* symbolName)error{return status(C.hipGetSymbolAddress(void** devPtr, const void* symbolName)).error("hipGetSymbolAddress")}
//func  hipGetSymbolSize(size_t* size, const void* symbolName)error{return status(C.hipGetSymbolSize(size_t* size, const void* symbolName)).error("hipGetSymbolSize")}
//func  hipMemcpyToSymbol(const void* symbolName, const void* src, size_t sizeBytes, size_t offset __dparm(0),hipMemcpyKind kind __dparm(hipMemcpyHostToDevice))error{return status(C.hipMemcpyToSymbol(const void* symbolName, const void* src, size_t sizeBytes, size_t offset __dparm(0),hipMemcpyKind kind __dparm(hipMemcpyHostToDevice))).error("hipMemcpyToSymbol")}
//func  hipMemcpyToSymbolAsync(const void* symbolName, const void* src,size_t sizeBytes, size_t offset,hipMemcpyKind kind, hipStream_t stream __dparm(0))error{return status(C.hipMemcpyToSymbolAsync(const void* symbolName, const void* src,size_t sizeBytes, size_t offset,hipMemcpyKind kind, hipStream_t stream __dparm(0))).error("hipMemcpyToSymbolAsync")}
//func  hipMemcpyFromSymbol(void* dst, const void* symbolName,size_t sizeBytes, size_t offset __dparm(0),hipMemcpyKind kind __dparm(hipMemcpyDeviceToHost))error{return status(C.hipMemcpyFromSymbol(void* dst, const void* symbolName,size_t sizeBytes, size_t offset __dparm(0),hipMemcpyKind kind __dparm(hipMemcpyDeviceToHost))).error("hipMemcpyFromSymbol")}
//func  hipMemcpyFromSymbolAsync(void* dst, const void* symbolName,size_t sizeBytes, size_t offset,hipMemcpyKind kind,hipStream_t stream __dparm(0))error{return status(C.hipMemcpyFromSymbolAsync(void* dst, const void* symbolName,size_t sizeBytes, size_t offset,hipMemcpyKind kind,hipStream_t stream __dparm(0))).error("hipMemcpyFromSymbolAsync")}
//func  hipMemcpyAsync(void* dst, const void* src, size_t sizeBytes, hipMemcpyKind kind,hipStream_t stream __dparm(0))error{return status(C.hipMemcpyAsync(void* dst, const void* src, size_t sizeBytes, hipMemcpyKind kind,hipStream_t stream __dparm(0))).error("hipMemcpyAsync")}
//func  hipMemset(void* dst, int value, size_t sizeBytes)error{return status(C.hipMemset(void* dst, int value, size_t sizeBytes)).error("hipMemset")}
//func  hipMemsetD8(hipDeviceptr_t dest, unsigned char value, size_t sizeBytes)error{return status(C.hipMemsetD8(hipDeviceptr_t dest, unsigned char value, size_t sizeBytes)).error("hipMemsetD8")}
//func  hipMemsetD32(hipDeviceptr_t dest, int value, size_t count)error{return status(C.hipMemsetD32(hipDeviceptr_t dest, int value, size_t count)).error("hipMemsetD32")}
//func  hipMemsetAsync(void* dst, int value, size_t sizeBytes, hipStream_t stream __dparm(0))error{return status(C.hipMemsetAsync(void* dst, int value, size_t sizeBytes, hipStream_t stream __dparm(0))).error("hipMemsetAsync")}
//func  hipMemsetD32Async(hipDeviceptr_t dst, int value, size_t count, hipStream_t stream __dparm(0))error{return status(C.hipMemsetD32Async(hipDeviceptr_t dst, int value, size_t count, hipStream_t stream __dparm(0))).error("hipMemsetD32Async")}
//func  hipMemset2D(void* dst, size_t pitch, int value, size_t width, size_t height)error{return status(C.hipMemset2D(void* dst, size_t pitch, int value, size_t width, size_t height)).error("hipMemset2D")}
//func  hipMemset2DAsync(void* dst, size_t pitch, int value, size_t width, size_t height,hipStream_t stream __dparm(0))error{return status(C.hipMemset2DAsync(void* dst, size_t pitch, int value, size_t width, size_t height,hipStream_t stream __dparm(0))).error("hipMemset2DAsync")}
//func  hipMemset3D(hipPitchedPtr pitchedDevPtr, int  value, hipExtent extent )error{return status(C.hipMemset3D(hipPitchedPtr pitchedDevPtr, int  value, hipExtent extent )).error("hipMemset3D")}
//func  hipMemset3DAsync(hipPitchedPtr pitchedDevPtr, int  value, hipExtent extent ,hipStream_t stream __dparm(0))error{return status(C.hipMemset3DAsync(hipPitchedPtr pitchedDevPtr, int  value, hipExtent extent ,hipStream_t stream __dparm(0))).error("hipMemset3DAsync")}
//func  hipMemGetInfo(size_t* free, size_t* total)error{return status(C.hipMemGetInfo(size_t* free, size_t* total)).error("hipMemGetInfo")}
//func  hipMemPtrGetInfo(void* ptr, size_t* size)error{return status(C.hipMemPtrGetInfo(void* ptr, size_t* size)).error("hipMemPtrGetInfo")}
//func  hipMallocArray(hipArray** array, const hipChannelFormatDesc* desc, size_t width, size_t height __dparm(0), unsigned int flags __dparm(hipArrayDefault))error{return status(C.hipMallocArray(hipArray** array, const hipChannelFormatDesc* desc, size_t width, size_t height __dparm(0), unsigned int flags __dparm(hipArrayDefault))).error("hipMallocArray")}
//func  hipArrayCreate(hipArray** pHandle, const HIP_ARRAY_DESCRIPTOR* pAllocateArray)error{return status(C.hipArrayCreate(hipArray** pHandle, const HIP_ARRAY_DESCRIPTOR* pAllocateArray)).error("hipArrayCreate")}
//func  hipArray3DCreate(hipArray** array, const HIP_ARRAY_DESCRIPTOR* pAllocateArray)error{return status(C.hipArray3DCreate(hipArray** array, const HIP_ARRAY_DESCRIPTOR* pAllocateArray)).error("hipArray3DCreate")}
//func  hipMalloc3D(hipPitchedPtr* pitchedDevPtr, hipExtent extent)error{return status(C.hipMalloc3D(hipPitchedPtr* pitchedDevPtr, hipExtent extent)).error("hipMalloc3D")}
//func  hipFreeArray(hipArray* array)error{return status(C.hipFreeArray(hipArray* array)).error("hipFreeArray")}
//func  hipMalloc3DArray(hipArray** array, const struct hipChannelFormatDesc* desc, struct hipExtent extent, unsigned int flags)error{return status(C.hipMalloc3DArray(hipArray** array, const struct hipChannelFormatDesc* desc, struct hipExtent extent, unsigned int flags)).error("hipMalloc3DArray")}
//func  hipMemcpy2D(void* dst, size_t dpitch, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind)error{return status(C.hipMemcpy2D(void* dst, size_t dpitch, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind)).error("hipMemcpy2D")}
//func  hipMemcpyParam2D(const hip_Memcpy2D* pCopy)error{return status(C.hipMemcpyParam2D(const hip_Memcpy2D* pCopy)).error("hipMemcpyParam2D")}
//func  hipMemcpy2DAsync(void* dst, size_t dpitch, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind, hipStream_t stream __dparm(0))error{return status(C.hipMemcpy2DAsync(void* dst, size_t dpitch, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind, hipStream_t stream __dparm(0))).error("hipMemcpy2DAsync")}
//func  hipMemcpy2DToArray(hipArray* dst, size_t wOffset, size_t hOffset, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind)error{return status(C.hipMemcpy2DToArray(hipArray* dst, size_t wOffset, size_t hOffset, const void* src, size_t spitch, size_t width, size_t height, hipMemcpyKind kind)).error("hipMemcpy2DToArray")}
//func  hipMemcpyToArray(hipArray* dst, size_t wOffset, size_t hOffset, const void* src, size_t count, hipMemcpyKind kind)error{return status(C.hipMemcpyToArray(hipArray* dst, size_t wOffset, size_t hOffset, const void* src, size_t count, hipMemcpyKind kind)).error("hipMemcpyToArray")}
//func  hipMemcpyFromArray(void* dst, hipArray_const_t srcArray, size_t wOffset, size_t hOffset, size_t count, hipMemcpyKind kind)error{return status(C.hipMemcpyFromArray(void* dst, hipArray_const_t srcArray, size_t wOffset, size_t hOffset, size_t count, hipMemcpyKind kind)).error("hipMemcpyFromArray")}
//func  hipMemcpyAtoH(void* dst, hipArray* srcArray, size_t srcOffset, size_t count)error{return status(C.hipMemcpyAtoH(void* dst, hipArray* srcArray, size_t srcOffset, size_t count)).error("hipMemcpyAtoH")}
//func  hipMemcpyHtoA(hipArray* dstArray, size_t dstOffset, const void* srcHost, size_t count)error{return status(C.hipMemcpyHtoA(hipArray* dstArray, size_t dstOffset, const void* srcHost, size_t count)).error("hipMemcpyHtoA")}
//func  hipMemcpy3D(const struct hipMemcpy3DParms* p)error{return status(C.hipMemcpy3D(const struct hipMemcpy3DParms* p)).error("hipMemcpy3D")}
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
