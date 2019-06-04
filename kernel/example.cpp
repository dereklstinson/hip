#include <hip/hip_runtime.h>
#include <hip/hip_runtime_api.h>


extern "C" __global__ void kernaddfloat(float *A, float *B, float *C, int N){
  int i = hipBlockDim_x * hipBlockIdx_x + hipThreadIdx_x;
  if (i<N){
C[i]=A[i]+B[i];
  }
}

    extern "C" __global__ void kernaddhalf(_Float16 *A, _Float16 *B, _Float16  *C, int N){
  int i = hipBlockDim_x * hipBlockIdx_x + hipThreadIdx_x;
  if (i<N){
C[i]=A[i]+B[i];
  }
}

  