#include <hip/hip_runtime.h>
#include <hip/hip_runtime_api.h>
#define GRID_LOOP_X(i, n)                                               \
    for (int i = hipBlockIdx_x * hipBlockDim_x + hipThreadIdx_x; i < n; \
         i += hipBlockDim_x * hipGridDim_x)

#define GRID_AXIS_LOOP(i, n, axis)                                 \
    for (int i = hipBlockIdx_axis * hipBlockDim_axis + hipThreadIdx_axis; i < n; \
         i += hipBlockDim_axis * hipGridDim_axis)

       
     
     
#define MAX 64*64 

extern "C" __global__ 
void KernAddFloat(const float *A,
                  const float *B,
                        float *C,
                  const int N){
        GRID_LOOP_X(i,N){
        C[i]=A[i]+B[i];
        }

}

extern "C" __global__ void KernAddHalf(
                        const _Float16 *A,
                        const  _Float16 *B,
                         _Float16  *C,
                         const int N){
GRID_LOOP_X(i,N){
C[i]=A[i]+B[i];

}
}