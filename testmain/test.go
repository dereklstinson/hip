package main

import (
	"fmt"
	"runtime"

	"github.com/dereklstinson/cutil"
	hip "github.com/dereklstinson/hipgo"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	runtime.LockOSThread()
	dev, err := hip.GetDevice()
	check(err)
	check(hip.SetDevice(dev))

	s, err := hip.CreateStream()
	check(err)
	size := uint(64 * 64)
	Ahalf := make([]float32, size)
	Bhalf := make([]float32, size)
	Chalf := make([]float32, size)

	for i := range Ahalf {
		Ahalf[i] = 1
		Bhalf[i] = 2
	}
	AhalfW, err := cutil.WrapGoMem(Ahalf)
	check(err)
	BhalfW, err := cutil.WrapGoMem(Bhalf)
	check(err)
	ChalfW, err := cutil.WrapGoMem(Chalf)
	check(err)
	gpuA := new(hip.DevicePtr)
	gpuB := new(hip.DevicePtr)
	gpuC := new(hip.DevicePtr)
	check(hip.Malloc(gpuA, size*4))

	check(hip.Malloc(gpuB, size*4))

	check(hip.Malloc(gpuC, size*4))

	check(hip.MemcpyHtoD(gpuA, AhalfW, size*4))

	check(hip.MemcpyHtoD(gpuB, BhalfW, size*4))

	check(hip.MemcpyHtoD(gpuC, ChalfW, size*4))

	check(hip.MemcpyDtoD(gpuC, gpuA, size*4))
	mod := new(hip.Module)
	///home/derek/go/src/github.com/dereklstinson/hipgo/kernel/example.cpp
	check(mod.Load("/home/derek/go/src/github.com/dereklstinson/hipgo/kernel/example.co"))

	kernelSingleAdd, err := mod.GetFunction("KernAddFloat")
	check(err)
	check(s.Sync())

	conf, err := hip.DeviceConfig1d(dev, (int32)(size))
	check(err)
	fmt.Println("Config:", "GridDim", conf.GridDim, "BlockDim", conf.BlockDim, "Elements", conf.Elements)
	check(kernelSingleAdd.Launch(conf.GridDim, 1, 1, conf.BlockDim, 1, 1, 0, s, gpuA, gpuB, gpuC, (conf.Elements)))
	check(s.Sync())

	check(hip.MemcpyDtoH(ChalfW, gpuC, size*4))

	check(s.Sync())
	check(hip.GetLastError())

	fmt.Println(Chalf)

}
