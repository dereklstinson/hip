
#ifndef __HIP_GO_KERNAL_EXAMPLE__
#define __HIP_GO_KERNAL_EXAMPLE__

 
#include <hip/hip_runtime_api.h>
#ifdef __cplusplus
extern "C"{
#endif
void KernAdd(float *A, float *B, float *C, int N); 


#ifdef __cplusplus
}
#endif

#endif //__HIP_GO_KERNAL_EXAMPLE__  