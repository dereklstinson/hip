package hip

//#include <hip/hip_runtime_api.h>
import "C"

/*
//func  hipEventCreateWithFlags(hipEvent_t* event, unsigned flags)error{return status(C.hipEventCreateWithFlags(hipEvent_t* event, unsigned flags)).error("hipEventCreateWithFlags")}
//func  hipEventCreate(hipEvent_t* event)error{return status(C.hipEventCreate(hipEvent_t* event)).error("hipEventCreate")}
//func  hipEventRecord(hipEvent_t event, hipStream_t stream)error{return status(C.hipEventRecord(hipEvent_t event, hipStream_t stream)).error("hipEventRecord")}
//func  hipEventDestroy(hipEvent_t event)error{return status(C.hipEventDestroy(hipEvent_t event)).error("hipEventDestroy")}
//func  hipEventSynchronize(hipEvent_t event)error{return status(C.hipEventSynchronize(hipEvent_t event)).error("hipEventSynchronize")}
//func  hipEventElapsedTime(float* ms, hipEvent_t start, hipEvent_t stop)error{return status(C.hipEventElapsedTime(float* ms, hipEvent_t start, hipEvent_t stop)).error("hipEventElapsedTime")}
//func  hipEventQuery(hipEvent_t event)error{return status(C.hipEventQuery(hipEvent_t event)).error("hipEventQuery")}


hipEventDefault
hipEventBlockingSync
hipEventDisableTiming
hipEventInterprocess
hipEventReleaseToDevice
hipEventReleaseToSystem



*/
