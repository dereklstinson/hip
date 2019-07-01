package hip

import (
	"runtime"
	"testing"

	"github.com/dereklstinson/cutil"
	"github.com/dereklstinson/half"
)

func TestFunction_Launch(t *testing.T) {
	runtime.LockOSThread()
	dev, err := GetDevice()
	if err != nil {
		t.Error(err)
	}
	err = SetDevice(dev)
	if err != nil {
		t.Error(err)
	}
	s, err := CreateStream()
	if err != nil {
		t.Error(err)
	}
	size := uint(64 * 64)
	Ahalf := make([]half.Float16, size)
	Bhalf := make([]half.Float16, size)
	Chalf := make([]half.Float16, size)
	ahalf1 := half.NewFloat16(1.0)
	bhalf2 := half.NewFloat16(2.0)
	for i := range Ahalf {
		Ahalf[i] = ahalf1
		Bhalf[i] = bhalf2
	}
	AhalfW, err := cutil.WrapGoMem(Ahalf)
	if err != nil {
		t.Error(err)
	}
	BhalfW, err := cutil.WrapGoMem(Bhalf)
	if err != nil {
		t.Error(err)
	}
	ChalfW, err := cutil.WrapGoMem(Chalf)
	if err != nil {
		t.Error(err)
	}
	gpuA := new(DevicePtr)
	gpuB := new(DevicePtr)
	gpuC := new(DevicePtr)
	err = Malloc(gpuA, size*2)
	if err != nil {
		t.Error(err)
	}
	err = Malloc(gpuB, size*2)
	if err != nil {
		t.Error(err)
	}
	err = Malloc(gpuC, size*2)
	if err != nil {
		t.Error(err)
	}

	//var kind MemCpyKind
	//kind.HtoD()
	err = MemcpyHtoD(gpuA, AhalfW, size*2)
	if err != nil {
		t.Error(err)
	}
	err = MemcpyHtoD(gpuB, BhalfW, size*2)
	if err != nil {
		t.Error(err)
	}
	err = MemcpyHtoD(gpuC, ChalfW, size*2)
	if err != nil {
		t.Error(err)
	}
	err = MemcpyDtoD(gpuC, gpuA, size*2)
	mod := new(Module)
	///home/derek/go/src/github.com/dereklstinson/hipgo/kernel/example.cpp
	err = mod.Load("/home/derek/go/src/github.com/dereklstinson/hipgo/kernel/example.co")
	if err != nil {
		t.Fatal(err)
	}
	kernelHalfAdd, err := mod.GetFunction("KernAddHalf")
	if err != nil {
		t.Error(err)
	}
	err = s.Sync()
	if err != nil {
		t.Error(err)
	}
	conf, err := DeviceConfig1d(dev, (int32)(size))
	if err != nil {
		t.Error(err)
	}
	err = kernelHalfAdd.Launch(conf.GridDim, 1, 1, conf.BlockDim, 1, 1, 0, s, gpuA, gpuB, gpuC, (conf.Elements))
	if err != nil {
		t.Error(err)
	}
	err = s.Sync()
	if err != nil {
		t.Error(err)
	}
	//kind.DtoH()
	err = MemcpyDtoH(ChalfW, gpuC, size*2)
	if err != nil {
		t.Error(err)
	}
	s.Sync()
	if err != nil {
		t.Error(err)
	}
	t.Error(Chalf)
	//fmt.Println(Chalf)
}
func TestFunction_Launch2(t *testing.T) {
	runtime.LockOSThread()
	dev, err := GetDevice()
	if err != nil {
		t.Error(err)
	}
	err = SetDevice(dev)
	if err != nil {
		t.Error(err)
	}
	s, err := CreateStream()
	if err != nil {
		t.Error(err)
	}
	size := uint(64 * 64)
	Ahalf := make([]float32, size)
	Bhalf := make([]float32, size)
	Chalf := make([]float32, size)

	for i := range Ahalf {
		Ahalf[i] = 1
		Bhalf[i] = 2
	}
	AhalfW, err := cutil.WrapGoMem(Ahalf)
	if err != nil {
		t.Error(err)
	}
	BhalfW, err := cutil.WrapGoMem(Bhalf)
	if err != nil {
		t.Error(err)
	}
	ChalfW, err := cutil.WrapGoMem(Chalf)
	if err != nil {
		t.Error(err)
	}
	gpuA := new(DevicePtr)
	gpuB := new(DevicePtr)
	gpuC := new(DevicePtr)
	err = Malloc(gpuA, size*4)
	if err != nil {
		t.Error(err)
	}
	err = Malloc(gpuB, size*4)
	if err != nil {
		t.Error(err)
	}
	err = Malloc(gpuC, size*4)
	if err != nil {
		t.Error(err)
	}
	//var kind MemCpyKind
	//kind.HtoD()
	err = MemcpyHtoD(gpuA, AhalfW, size*4)
	if err != nil {
		t.Error(err)
	}
	err = MemcpyHtoD(gpuB, BhalfW, size*4)
	if err != nil {
		t.Error(err)
	}
	err = MemcpyHtoD(gpuC, ChalfW, size*4)
	if err != nil {
		t.Error(err)
	}
	err = MemcpyDtoD(gpuC, gpuA, size*4)
	mod := new(Module)
	///home/derek/go/src/github.com/dereklstinson/hipgo/kernel/example.cpp
	err = mod.Load("/home/derek/go/src/github.com/dereklstinson/hipgo/kernel/example.co")
	if err != nil {
		t.Fatal(err)
	}
	kernelSingleAdd, err := mod.GetFunction("KernAddFloat")
	if err != nil {
		t.Error(err)
	}
	err = s.Sync()
	if err != nil {
		t.Error(err)
	}
	conf, err := DeviceConfig1d(dev, (int32)(size))
	if err != nil {
		t.Error(err)
	}
	err = kernelSingleAdd.Launch(conf.GridDim, 1, 1, conf.BlockDim, 1, 1, 0, s, gpuA, gpuB, gpuC, (conf.Elements))
	if err != nil {
		t.Error(err)
	}
	err = s.Sync()
	if err != nil {
		t.Error(err)
	}
	//kind.DtoH()
	err = MemcpyDtoH(ChalfW, gpuC, size*4)
	if err != nil {
		t.Error(err)
	}
	s.Sync()
	if err != nil {
		t.Error(err)
	}
	err = GetLastError()
	if err != nil {
		t.Fatal(err)
	}
	t.Error("check printing")
	//t.Error(Chalf)
	//fmt.Println(Chalf)
}
