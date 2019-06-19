package hip

//#include <hip/hip_runtime_api.h>
import "C"

func hipDeviceSynchronize() error {
	return status(C.hipDeviceSynchronize()).error("hipDeviceSynchronize")
}
func hipDeviceReset() error { return status(C.hipDeviceReset()).error("hipDeviceReset") }

func hipSetDevice(deviceID int32) error {
	return status(C.hipSetDevice((C.int)(deviceID))).error("hipSetDevice")
}
func hipGetDevice() (deviceID int32, err error) {
	err = status(C.hipGetDevice((*C.int)(&deviceID))).error("hipGetDevice")
	return deviceID, err
}
func hipGetDeviceCount() (count int32, err error) {
	err = status(C.hipGetDeviceCount((*C.int)(&count))).error("hipGetDeviceCount")
	return count, err
}

//func  hipDeviceGetAttribute(int* pi, hipDeviceAttribute_t attr, int deviceId)error{return status(C.hipDeviceGetAttribute(int* pi, hipDeviceAttribute_t attr, int deviceId)).error("hipDeviceGetAttribute")}
//func  hipGetDeviceProperties(hipDeviceProp_t* prop, int deviceId)error{return status(C.hipGetDeviceProperties(hipDeviceProp_t* prop, int deviceId)).error("hipGetDeviceProperties")}
//func  hipDeviceSetCacheConfig(hipFuncCache_t cacheConfig)error{return status(C.hipDeviceSetCacheConfig(hipFuncCache_t cacheConfig)).error("hipDeviceSetCacheConfig")}
//func  hipDeviceGetCacheConfig(hipFuncCache_t* cacheConfig)error{return status(C.hipDeviceGetCacheConfig(hipFuncCache_t* cacheConfig)).error("hipDeviceGetCacheConfig")}
//func  hipDeviceGetLimit(size_t* pValue, enum hipLimit_t limit)error{return status(C.hipDeviceGetLimit(size_t* pValue, enum hipLimit_t limit)).error("hipDeviceGetLimit")}
//func  hipFuncSetCacheConfig(const void* func, hipFuncCache_t config)error{return status(C.hipFuncSetCacheConfig(const void* func, hipFuncCache_t config)).error("hipFuncSetCacheConfig")}
//func  hipDeviceGetSharedMemConfig(hipSharedMemConfig* pConfig)error{return status(C.hipDeviceGetSharedMemConfig(hipSharedMemConfig* pConfig)).error("hipDeviceGetSharedMemConfig")}
//func  hipDeviceSetSharedMemConfig(hipSharedMemConfig config)error{return status(C.hipDeviceSetSharedMemConfig(hipSharedMemConfig config)).error("hipDeviceSetSharedMemConfig")}
//func  hipSetDeviceFlags(unsigned flags)error{return status(C.hipSetDeviceFlags(unsigned flags)).error("hipSetDeviceFlags")}
//func  hipChooseDevice(int* device, const hipDeviceProp_t* prop)error{return status(C.hipChooseDevice(int* device, const hipDeviceProp_t* prop)).error("hipChooseDevice")}
/*


device flags?
hipDeviceScheduleAuto
hipDeviceScheduleSpin
hipDeviceScheduleYield
hipDeviceScheduleBlockingSync
hipDeviceScheduleMask
hipDeviceMapHost
hipDeviceLmemResizeToMax
*/
