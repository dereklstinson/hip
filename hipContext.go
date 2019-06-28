package hip

//#include <hip/hip_runtime_api.h>
import "C"

type Context struct {
	c C.hipCtx_t
}

/*
func CreateContext(flags uint32, d *Device) (c *Context, err error) {
	c = new(Context)
	err = status(C.hipCtxCreate(&c.c, (C.uint)(flags), d.c())).error("hipCtxCreate")
	return c, err
}
*/
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
