
#include <hip/hip_runtime_api.h>
#include <stdio.h>
#define LEN 4096
#define SIZE LEN << 2

#define fileName "example.co"
#define kernel_name "KernAddFloat"

#define HIP_CHECK(status)                                                                          \
    if (status != hipSuccess) {                                                                    \
        printf("GotStatus: %d",status);                                                            \
        return -1;                                                                                   \
    }

int main() {
   // float *A, *B, *C;
    hipDeviceptr_t Ad, Bd, Cd;
    float A[LEN];
    float B[LEN];
    float C[LEN];

    for (uint32_t i = 0; i < LEN; i++) {
        A[i] = 1.0f;
        B[i] = 2.0f;
        C[i] = 0.0f;
    }

    HIP_CHECK(hipInit(0));

    hipDevice_t device;

    HIP_CHECK(hipDeviceGet(&device, 0));
    HIP_CHECK(hipMalloc((void**)&Ad, SIZE));
    HIP_CHECK(hipMalloc((void**)&Bd, SIZE));
    HIP_CHECK(hipMalloc((void**)&Cd, SIZE));
    HIP_CHECK(hipMemcpyHtoD(Ad, A, SIZE));
    HIP_CHECK(hipMemcpyHtoD(Bd, B, SIZE));

    hipModule_t Module;
    hipFunction_t Function;
    HIP_CHECK(hipModuleLoad(&Module, fileName));
    HIP_CHECK(hipModuleGetFunction(&Function, Module, kernel_name));

    hipStream_t stream;
    HIP_CHECK(hipStreamCreate(&stream));

    struct {
        void* _Ad;
        void* _Bd;
        void* _Cd;
        int _elements;
    } args;
   int elements=(int)(LEN);
    args._Ad = (void*) Ad;
    args._Bd = (void*) Bd;
    args._Cd = (void*) Cd;
    args._elements= elements;  
    size_t size = sizeof(args);
    printf("size of size %d", (int)size);
    void* config[] = {HIP_LAUNCH_PARAM_BUFFER_POINTER, &args, HIP_LAUNCH_PARAM_BUFFER_SIZE, &size,
                      HIP_LAUNCH_PARAM_END};
    HIP_CHECK(hipModuleLaunchKernel(Function, 4, 1, 1, 1024, 1, 1, 0, stream, NULL, (void**)&config));

    HIP_CHECK(hipStreamDestroy(stream));

    HIP_CHECK(hipMemcpyDtoH(C, Cd, SIZE));
int counter;
    for (uint32_t i = 0; i < LEN; i++) {
        if (C[i]!=3.0){
          counter++;
        }
    }
 //   hipDeviceReset();
printf("\nGot %d out of %d\n",LEN-counter,LEN);

 return 0;
}