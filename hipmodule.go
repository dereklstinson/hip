package hip

/*
#include <string.h>
#include <hip/hip_runtime_api.h>
#include <stdlib.h>
#include <stdio.h>


const void ** voiddptrnull = NULL;
const size_t ptrSize = sizeof(void *);
const size_t maxArgSize = 8;

hipError_t golangLaunchKernelwithcharbuffer(hipFunction_t f, unsigned int gridDimX, unsigned int gridDimY, unsigned int gridDimZ,
								 unsigned int blockDimX, unsigned int blockDimY, unsigned int blockDimZ,
                                 unsigned int sharedMemBytes, hipStream_t stream,
                                 unsigned char *args, size_t sib){
						     	 void *config[] = {HIP_LAUNCH_PARAM_BUFFER_POINTER,args,HIP_LAUNCH_PARAM_BUFFER_SIZE,&sib,HIP_LAUNCH_PARAM_END};
  return hipModuleLaunchKernel(f,gridDimX, gridDimY,gridDimZ, blockDimX,blockDimY,blockDimZ, sharedMemBytes, stream, NULL,(void**)&config);
								 }

*/
import "C"
import (
	"errors"
	"fmt"
	"github.com/dereklstinson/cutil"
	"unsafe"
)

type Module struct {
	m      C.hipModule_t
	loaded bool
}

func (m *Module) Load(filename string) error {
	if m.loaded {
		return errors.New("(m *Modlue)Load(): Goside: Module already loaded")
	}
	fname := C.CString(filename)
	return status(C.hipModuleLoad(&m.m, fname)).error("(m *Modlue)Load()")
}
func (m *Module) UnLoad() error {
	if !m.loaded {
		return errors.New("(m *Modlue)Unload(): Goside: Module Not loaded")
	}
	return status(C.hipModuleUnload(m.m)).error("(m *Modlue)Unload()")
}

func (m *Module) GetFunction(kernelname string) (f *Function, err error) {
	var fun C.hipFunction_t
	kname := C.CString(kernelname)
	err = status(C.hipModuleGetFunction(&fun, m.m, kname)).error("(m *Module)GetFunction()")
	f = &Function{
		f: fun,
	}
	return f, err
}
func (m *Module) GetGlobal(name string) (d *DevicePtr, sib uint, err error) {
	d = new(DevicePtr)
	var sizet C.size_t
	err = status(C.hipModuleGetGlobal(&d.d, &sizet, m.m, C.CString(name))).error(" (m *Module) GetGlobal()")
	sib = (uint)(sizet)
	return d, sib, err
}
func (m *Module) LoadData(image cutil.Mem) error {
	return status(C.hipModuleLoadData(&m.m, image.Ptr())).error("hipModuleLoadData")
}

//LoadDataEx options and vals not used for now
func (m *Module) LoadDataEx(image cutil.Mem, options []JitOption, vals []interface{}) error {
	//	ops,	numOptions:=jithiparray(options)
	return status(C.hipModuleLoadDataEx(&m.m, image.Ptr(), 0, nil, nil)).error("hipModuleLoadDataEx")
}

type Function struct {
	f                 C.hipFunction_t
	sizeofargs        C.size_t
	argsbuffer        [255]C.uchar
	argsbuffer2       [255]C.uchar
	unsafeargsbuffer  unsafe.Pointer
	unsafeargsbuffer2 unsafe.Pointer
}

//Launch launches a kernel.
func (f *Function) Launch(gridDimx, gridDimy, gridDimz uint32,
	blockDimx, blockDimy, blockDimz uint32,
	sharedMemBytes uint32,
	s *Stream,
	args ...interface{}) error {
	//var shold unsafe.Pointer
	//	f.interface2uchararray(args)
	//f.interface2unsafePointercomplete(args)
	f.bufferparams(args)

	return status(C.golangLaunchKernelwithcharbuffer(f.f,
		(C.uint)(gridDimx), (C.uint)(gridDimy), (C.uint)(gridDimz),
		(C.uint)(blockDimx), (C.uint)(blockDimy), (C.uint)(blockDimz),
		(C.uint)(sharedMemBytes),
		s.s,
		(&f.argsbuffer2[0]), f.sizeofargs)).error("golangLaunchKernel")

}

func offsetelement(ptr unsafe.Pointer, element int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) + 8*uintptr(element))
}

func (f *Function) bufferparams(args []interface{}) error {
	/*
		if f.args == nil || len(args) != len(f.args) {
			f.args = make([]unsafe.Pointer, len(args))
		}
	*/
	for i := range f.argsbuffer {
		f.argsbuffer[i] = 0
		f.argsbuffer2[i] = 0
	}

	f.unsafeargsbuffer = unsafe.Pointer(&f.argsbuffer[0])
	f.unsafeargsbuffer2 = unsafe.Pointer(&f.argsbuffer2[0])
	for i := range args {
		fmt.Println(i)
		switch x := args[i].(type) {
		case nil:
			*((*unsafe.Pointer)(offsetelement(f.unsafeargsbuffer, i))) = offsetelement(f.unsafeargsbuffer2, i) // argp[i] = &argv[i]
			*((*uint64)(offsetelement(f.unsafeargsbuffer2, i))) = *((*uint64)(nil))                            // argv[i] = *f.args[i]
		case *DevicePtr:
			*((*unsafe.Pointer)(offsetelement(f.unsafeargsbuffer, i))) = offsetelement(f.unsafeargsbuffer2, i) // argp[i] = &argv[i]
			usptr := (unsafe.Pointer)(&x.d)
			*((*uint64)(offsetelement(f.unsafeargsbuffer2, i))) = *((*uint64)(usptr))
		case cutil.Mem:
			*((*unsafe.Pointer)(offsetelement(f.unsafeargsbuffer, i))) = offsetelement(f.unsafeargsbuffer2, i) // argp[i] = &argv[i]
			usptr := (unsafe.Pointer)(x.DPtr())
			*((*uint64)(offsetelement(f.unsafeargsbuffer2, i))) = *((*uint64)(usptr)) // argv[i] = *f.args[i]
		default:
			y := cutil.CScalarConversion(x)
			if y == nil {
				return errors.New("Unsupported arg passed")
			}
			*((*unsafe.Pointer)(offsetelement(f.unsafeargsbuffer, i))) = offsetelement(f.unsafeargsbuffer2, i) // argp[i] = &argv[i]
			*((*uint64)(offsetelement(f.unsafeargsbuffer2, i))) = *((*uint64)(y.CPtr()))                       // argv[i] = *f.args[i]

		}

	}
	f.sizeofargs = (C.size_t)(len(args) * 8)
	return nil

}

func freeargs(x []unsafe.Pointer) {
	for i := range x {
		C.free(x[i])
	}

}

type FuncAtributes C.hipFuncAttributes

func (f *FuncAtributes) cptr() *C.hipFuncAttributes {
	return (*C.hipFuncAttributes)(f)
}
func (f FuncAtributes) BinaryVersion() int32 {
	return (int32)(f.binaryVersion)
}
func (f FuncAtributes) CacheModeCA() int32 {
	return (int32)(f.cacheModeCA)
}
func (f FuncAtributes) MaxDynamicSharedSizeBytes() int32 {
	return (int32)(f.maxDynamicSharedSizeBytes)
}
func (f FuncAtributes) MaxThreadsPerBlock() int32 {
	return (int32)(f.maxThreadsPerBlock)
}
func (f FuncAtributes) NumRegs() int32 {
	return (int32)(f.numRegs)
}
func (f FuncAtributes) PreferredShmemCarveout() int32 {
	return (int32)(f.preferredShmemCarveout)
}
func (f FuncAtributes) PtxVersion() int32 {
	return (int32)(f.ptxVersion)
}
func (f FuncAtributes) ConstSIB() uint {
	return (uint)(f.constSizeBytes)
}
func (f FuncAtributes) LocalSIB() uint {
	return (uint)(f.localSizeBytes)
}
func (f FuncAtributes) SharedSIB() uint {
	return (uint)(f.sharedSizeBytes)
}

func FuncGetAttributes(function cutil.Mem) (attr FuncAtributes, err error) {
	err = status(C.hipFuncGetAttributes(attr.cptr(), function.Ptr())).error("hipFuncGetAttributes")
	return attr, err
}

//func  hipModuleGetTexRef(textureReference** texRef, hipModule_t hmod, const char* name)error{return status(C.hipModuleGetTexRef(textureReference** texRef, hipModule_t hmod, const char* name)).error("hipModuleGetTexRef")}
//func  hipLaunchCooperativeKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)error{return status(C.hipLaunchCooperativeKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)).error("hipLaunchCooperativeKernelMultiDevice")}
//func  hipOccupancyMaxActiveBlocksPerMultiprocessor(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize)error{return status(C.hipOccupancyMaxActiveBlocksPerMultiprocessor(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize)).error("hipOccupancyMaxActiveBlocksPerMultiprocessor")}
//func  hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize, unsigned int flags)error{return status(C.hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize, unsigned int flags)).error("hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags")}
//func  hipExtLaunchMultiKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)error{return status(C.hipExtLaunchMultiKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)).error("hipExtLaunchMultiKernelMultiDevice")}

//func  hipLaunchByPtr(const void* func)error{return status(C.hipLaunchByPtr(const void* func)).error("hipLaunchByPtr")}

/*
func  LaunchCooperativeKernel(f cutil.Mem , gridDim, blockDim Dim3,sharedMemBytes uint32,s *Stream, kernelParams ...unsafe.Pointer)error{
arg1:=	(unsafe.Pointer)(C.malloc(C.size_t(8*len(kernelParams))))
arg2:=	(unsafe.Pointer)(C.malloc(C.size_t(8*len(kernelParams))))
defer C.free(arg1)
defer C.free(arg2)
for i := range kernelParams {
	*((*unsafe.Pointer)(offsetelement(arg1, i))) = offsetelement(arg2, i) // argp[i] = &argv[i]
			*((*uint64)(offsetelement(arg2, i))) = *((*uint64)(kernelParams[i]))                            // argv[i] = *f.args[i]

}
	return status(C.hipLaunchCooperativeKernel(f.Ptr(),
	 gridDim.c(),
	  blockDim.c(),
	 &arg2,
	   (C.uint) (sharedMemBytes),s.s)).error("hipLaunchCooperativeKernel")
}

*/
type JitOption C.hipJitOption

func (j JitOption) c() C.hipJitOption {
	return (C.hipJitOption)(j)
}
func jithiparray(slice []JitOption) (array []C.hipJitOption, length C.uint) {
	length = (C.uint)(len(slice))
	array = make([]C.hipJitOption, length)
	for i := range slice {
		array[i] = slice[i].c()
	}
	return array, length
}
func (j *JitOption) MaxRegisters() JitOption { *j = (JitOption)(C.hipJitOptionMaxRegisters); return *j }
func (j *JitOption) ThreadsPerBlock() JitOption {
	*j = (JitOption)(C.hipJitOptionThreadsPerBlock)
	return *j
}
func (j *JitOption) WallTime() JitOption      { *j = (JitOption)(C.hipJitOptionWallTime); return *j }
func (j *JitOption) InfoLogBuffer() JitOption { *j = (JitOption)(C.hipJitOptionInfoLogBuffer); return *j }
func (j *JitOption) InfoLogBufferSizeBytes() JitOption {
	*j = (JitOption)(C.hipJitOptionInfoLogBufferSizeBytes)
	return *j
}
func (j *JitOption) ErrorLogBuffer() JitOption {
	*j = (JitOption)(C.hipJitOptionErrorLogBuffer)
	return *j
}
func (j *JitOption) ErrorLogBufferSizeBytes() JitOption {
	*j = (JitOption)(C.hipJitOptionErrorLogBufferSizeBytes)
	return *j
}
func (j *JitOption) OptimizationLevel() JitOption {
	*j = (JitOption)(C.hipJitOptionOptimizationLevel)
	return *j
}
func (j *JitOption) TargetFromContext() JitOption {
	*j = (JitOption)(C.hipJitOptionTargetFromContext)
	return *j
}
func (j *JitOption) Target() JitOption { *j = (JitOption)(C.hipJitOptionTarget); return *j }
func (j *JitOption) FallbackStrategy() JitOption {
	*j = (JitOption)(C.hipJitOptionFallbackStrategy)
	return *j
}
func (j *JitOption) GenerateDebugInfo() JitOption {
	*j = (JitOption)(C.hipJitOptionGenerateDebugInfo)
	return *j
}
func (j *JitOption) LogVerbose() JitOption { *j = (JitOption)(C.hipJitOptionLogVerbose); return *j }
func (j *JitOption) GenerateLineInfo() JitOption {
	*j = (JitOption)(C.hipJitOptionGenerateLineInfo)
	return *j
}
func (j *JitOption) CacheMode() JitOption   { *j = (JitOption)(C.hipJitOptionCacheMode); return *j }
func (j *JitOption) Sm3xOpt() JitOption     { *j = (JitOption)(C.hipJitOptionSm3xOpt); return *j }
func (j *JitOption) FastCompile() JitOption { *j = (JitOption)(C.hipJitOptionFastCompile); return *j }
func (j *JitOption) NumOptions() JitOption  { *j = (JitOption)(C.hipJitOptionNumOptions); return *j }
