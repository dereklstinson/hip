package hip

//#include <hip/hip_runtime_api.h>
import "C"

/*
hipHostRegisterDefault
hipHostRegisterPortable
hipHostRegisterMapped
hipHostRegisterIoMemory
hipExtHostRegisterCoarseGrained
*/
//func  HostRegister(void* hostPtr, size_t sizeBytes, unsigned int flags)error{return status(C.hipHostRegister(void* hostPtr, size_t sizeBytes, unsigned int flags)).error("hipHostRegister")}
//func  HostUnregister(void* hostPtr)error{return status(C.hipHostUnregister(void* hostPtr)).error("hipHostUnregister")}
