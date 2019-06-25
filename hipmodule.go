package hip

/*
#include <string.h>
#include <hip/hip_runtime_api.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/dereklstinson/cutil"
)

type Module struct {
	m      C.hipModule_t
	loaded bool
}

/*
type Argument struct {
	arg interface{}
	sib uint
}

func CreateArgument(arg interface{}, unitsib uint) *Argument {
	return &Argument{arg: arg, sib: unitsib}
}
*/
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
	f = new(Function)
	kname := C.CString(kernelname)
	err = status(C.hipModuleGetFunction(&f.f, m.m, kname)).error("(m *Module)GetFunction()")
	return f, err
}

type Function struct {
	f          C.hipFunction_t
	args       []unsafe.Pointer
	sizeofargs uintptr
}

func (f *Function) Launch(gridDimx, gridDimy, gridDimz uint32,
	blockDimx, blockDimy, blockDImz uint32,
	sharedMemBytes uint32,
	s *Stream,
	extra bool,
	args ...interface{}) error {
		var shold unsafe.Pointer

		if s == nil {
			shold = nil
		} else {
			shold = s.Ptr()
		}
	//C.HIP_LAUNCH_PARAM_BUFFER_POINTER
	//C.HIP_LAUNCH_PARAM_BUFFER_SIZE
	//C.HIP_LAUNCH_PARAM_END
config:=(unsafe.Pointer)({C.HIP_LAUNCH_PARAM_BUFFER_POINTER,f.args[0],
	C.HIP_LAUNCH_PARAM_BUFFER_SIZE, f.sizeofargs,C.HIP_LAUNCH_PARAM_END})
return status(C.hipModuleLaunchKernel(f.f,(C.uint)(gridDimX),(C.uint)(gridDimy),(C.uint)(gridDimz), 
(C.uint)(blockDimx),(C.uint)(blockDimy),(C.uint)(blockDImz), 
(C.int)( sharedMemBytes),
 shold,nil,&config)).error("hipModuleLaunchKernel")}
	
}
func (f *Function) interface2unsafePointercomplete(args []interface{}) error {
	if len(args) != len(f.args) {
		f.firstrunInterface2UnsafePointer(args)
	}
	if f.args == nil {
		f.firstrunInterface2UnsafePointer(args)
	}
	for i := range args {

		switch x := args[i].(type) {
		case nil:
			y := reflect.TypeOf(x)
			C.memcpy(f.args[i], unsafe.Pointer(&x), (C.size_t)(y.Size()))

		case cutil.Mem:
			y := reflect.TypeOf(x)
			C.memcpy(f.args[i], unsafe.Pointer(x.DPtr()), (C.size_t)(y.Size()))
		case bool:
			if x {
				val := C.int(255)
				C.memcpy(f.args[i], unsafe.Pointer(&val), C.size_t(4))
			} else {
				val := C.int(0)
				C.memcpy(f.args[i], unsafe.Pointer(&val), C.size_t(4))
			}
		case int:
			val := C.int(x)
			C.memcpy(f.args[i], unsafe.Pointer(&val), C.size_t(4))
		case uint:
			val := C.uint(x)
			C.memcpy(f.args[i], unsafe.Pointer(&val), C.size_t(4))
		default:
			/*
					val := reflect.ValueOf(x)
					sizeof := reflect.TypeOf(x).Size()
					y := unsafe.Pointer(val.Pointer())

					C.memcpy(k.args[i], y, (C.size_t)(sizeof))

				}
			*/
			scalar := cutil.CScalarConversion(x)
			if scalar == nil {
				return fmt.Errorf("Kernel Launch - type %T not supported .. %+v", x, x)
			}

			C.memcpy(f.args[i], scalar.CPtr(), C.size_t(scalar.SIB()))
		}

	}
	return nil
}
func (f *Function) firstrunInterface2UnsafePointer(args []interface{}) error {
	f.args = make([]unsafe.Pointer, len(args))
	var argsizes uintptr
	for i := range args {
		switch x := args[i].(type) {
		case nil:
			y := reflect.TypeOf(x)
			f.args[i] = (unsafe.Pointer)(C.malloc((C.ulong)(y.Size())))
			C.memcpy(f.args[i], unsafe.Pointer(&x), (C.size_t)(y.Size()))
			argsizes += y.Size()
		case cutil.Mem:
			if x == nil {
				y := reflect.TypeOf(x)
				f.args[i] = (unsafe.Pointer)(C.malloc((C.ulong)(y.Size())))
				C.memcpy(f.args[i], unsafe.Pointer(&x), (C.size_t)(y.Size()))
				argsizes += y.Size()
			}
			y := reflect.TypeOf(x)
			f.args[i] = (unsafe.Pointer)(C.malloc((C.ulong)(y.Size())))
			C.memcpy(f.args[i], unsafe.Pointer(x.DPtr()), (C.size_t)(y.Size()))
			argsizes += y.Size()
		case bool:
			f.args[i] = (unsafe.Pointer)(C.malloc((C.size_t)(4)))
			if x {
				val := C.int(255)
				C.memcpy(f.args[i], unsafe.Pointer(&val), C.size_t(4))
			} else {
				val := C.int(0)
				C.memcpy(f.args[i], unsafe.Pointer(&val), C.size_t(4))
			}
		case int:
			val := C.int(x)
			f.args[i] = (unsafe.Pointer)(C.malloc((C.size_t)(4)))
			C.memcpy(f.args[i], unsafe.Pointer(&val), C.size_t(4))
		case uint:
			val := C.uint(x)
			f.args[i] = (unsafe.Pointer)(C.malloc((C.size_t)(4)))
			C.memcpy(f.args[i], unsafe.Pointer(&val), C.size_t(4))

		default:
			scalar := cutil.CScalarConversion(x)
			if scalar == nil {
				return fmt.Errorf("Kernel Launch - type %T not supported .. %+v", x, x)
			}
			f.args[i] = (unsafe.Pointer)(C.malloc((C.size_t)(scalar.SIB())))
			C.memcpy(f.args[i], scalar.CPtr(), C.size_t(scalar.SIB()))
		}

	}

	f.sizeofargs = argsizes
	runtime.SetFinalizer(f.args, freeargs)
	return nil
}

func freeargs(x []unsafe.Pointer) {
	for i := range x {
		C.free(x[i])
	}

}

//func  hipFuncGetAttributes(struct hipFuncAttributes* attr, const void* func)error{return status(C.hipFuncGetAttributes(struct hipFuncAttributes* attr, const void* func)).error("hipFuncGetAttributes")}
//func  hipModuleGetGlobal(hipDeviceptr_t* dptr, size_t* bytes, hipModule_t hmod, const char* name)error{return status(C.hipModuleGetGlobal(hipDeviceptr_t* dptr, size_t* bytes, hipModule_t hmod, const char* name)).error("hipModuleGetGlobal")}
//func  hipModuleGetTexRef(textureReference** texRef, hipModule_t hmod, const char* name)error{return status(C.hipModuleGetTexRef(textureReference** texRef, hipModule_t hmod, const char* name)).error("hipModuleGetTexRef")}
//func  hipModuleLoadData(hipModule_t* module, const void* image)error{return status(C.hipModuleLoadData(hipModule_t* module, const void* image)).error("hipModuleLoadData")}
//func  hipModuleLoadDataEx(hipModule_t* module, const void* image, unsigned int numOptions,hipJitOption* options, void** optionValues)error{return status(C.hipModuleLoadDataEx(hipModule_t* module, const void* image, unsigned int numOptions,hipJitOption* options, void** optionValues)).error("hipModuleLoadDataEx")}
//func  hipLaunchCooperativeKernel(const void* f, dim3 gridDim, dim3 blockDimX,void** kernelParams, unsigned int sharedMemBytes,hipStream_t stream)error{return status(C.hipLaunchCooperativeKernel(const void* f, dim3 gridDim, dim3 blockDimX,void** kernelParams, unsigned int sharedMemBytes,hipStream_t stream)).error("hipLaunchCooperativeKernel")}
//func  hipLaunchCooperativeKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)error{return status(C.hipLaunchCooperativeKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)).error("hipLaunchCooperativeKernelMultiDevice")}
//func  hipOccupancyMaxActiveBlocksPerMultiprocessor(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize)error{return status(C.hipOccupancyMaxActiveBlocksPerMultiprocessor(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize)).error("hipOccupancyMaxActiveBlocksPerMultiprocessor")}
//func  hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize, unsigned int flags)error{return status(C.hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(int* numBlocks, const void* f, int  blockSize, size_t dynamicSMemSize, unsigned int flags)).error("hipOccupancyMaxActiveBlocksPerMultiprocessorWithFlags")}
//func  hipExtLaunchMultiKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)error{return status(C.hipExtLaunchMultiKernelMultiDevice(hipLaunchParams* launchParamsList,int  numDevices, unsigned int  flags)).error("hipExtLaunchMultiKernelMultiDevice")}
//func  hipLaunchByPtr(const void* func)error{return status(C.hipLaunchByPtr(const void* func)).error("hipLaunchByPtr")}
