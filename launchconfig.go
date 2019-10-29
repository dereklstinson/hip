package hip

import (
	"errors"

//	"github.com/dereklstinson/GoCudnn/kernels"
)

//Config is for a 1d kernel launch
type Config struct {
	Elements int32
	BlockDim uint32
	GridDim  uint32
}

//Config2d are parameters for the kernel launch
type Config2d struct {
	Dimx            int32
	Dimy            int32
	ThreadPerBlockx uint32
	ThreadPerBlocky uint32
	BlockCountx     uint32
	BlockCounty     uint32
}

//Config3d are parameters for the kernel launch
type Config3d struct {
	Dimx            int32
	Dimy            int32
	Dimz            int32
	ThreadPerBlockx uint32
	ThreadPerBlocky uint32
	ThreadPerBlockz uint32
	BlockCountx     uint32
	BlockCounty     uint32
	BlockCountz     uint32
}
func divup(a, b int32) int32 {
	return (a + b - 1) / b
}

func DeviceConfig1d(device int32, elements int32) (Config, error) {
	prop, err := GetDeviceProperties(device)
	if err != nil {
		return Config{}, err
	}
	ptc := min((prop.MultiProcessorCount() * prop.MaxThreadsPerMultiProcessor()), elements)
	threadperblock := min(1024, prop.MaxThreadsPerBlock())
	innerbcount := divup(ptc, threadperblock)
	bcount := min(innerbcount, prop.MultiProcessorCount())
	return Config{
		Elements: elements,
		BlockDim: uint32(threadperblock),
		GridDim:  uint32(bcount),
	}, nil
}
func DeviceConfig2d(device int32, xdim, ydim int32) (Config2d, error) {
	if xdim < 1 || ydim < 1 {
		return Config2d{}, errors.New("DeviceConfig2d: xdim < 1 || ydim < 1")
	}
	prop, err := GetDeviceProperties(device)
	if err != nil {
		return Config2d{}, err
	}
	kthreadsperblock := int32(256)
	gputhreads := (prop.MultiProcessorCount() * prop.MaxThreadsPerMultiProcessor())
	blockx := min(xdim, kthreadsperblock)
	blocky := max(kthreadsperblock/blockx, 1)
	maxblocks := max(gputhreads/kthreadsperblock, 1)
	ratiox := divideandroundup(xdim, blockx)
	gridx := uint32(min(int32(ratiox), maxblocks))
	return Config2d{
		Dimx:            xdim,
		Dimy:            ydim,
		ThreadPerBlockx: (uint32)(blockx),
		ThreadPerBlocky: (uint32)(blocky),
		BlockCountx:     gridx,
		BlockCounty:     uint32(min(maxblocks/int32(gridx), max(ydim/blocky, 1))),
	}, nil

}

//DeviceConfig3d returns configs for the kernel launch
func DeviceConfig3d(device int32, xdim, ydim, zdim int32) (Config3d, error) {
	if xdim < 1 || ydim < 1 || zdim < 1 {
		return Config3d{}, errors.New("xdim < 1 || ydim < 1 ||zdim<1")
	}
	prop, err := GetDeviceProperties(device)
	if err != nil {
		return Config3d{}, err
	}
	kthreadsperblock := int32(256)
	gputhreads := (prop.MultiProcessorCount() * prop.MaxThreadsPerMultiProcessor())
	//Blocks
	maxthreadsperdim := prop.MaxThreadsDim()
	blockx := min3(xdim, kthreadsperblock, maxthreadsperdim[0])
	blocky := min3(ydim, max(kthreadsperblock/blockx, 1), maxthreadsperdim[1])
	blockz := min3(zdim, max(kthreadsperblock/(blockx*blocky), 1), maxthreadsperdim[2])
	maxblocks := max(gputhreads/kthreadsperblock, 1)
	ratiox := divideandroundup(xdim, blockx)
	//Grids
	blocksperdim := prop.MaxGridSize()
	gridx := uint32(min3(maxblocks, int32(ratiox), blocksperdim[0]))
	ratioy := divideandroundup(ydim, blocky)
	ratioy2 := divideandroundup(maxblocks, int32(gridx))
	gridy := uint32(min3(int32(ratioy), int32(ratioy2), blocksperdim[1]))
	ratioz := divideandroundup(maxblocks, int32(gridx*gridy))
	ratioz2 := divideandroundup(zdim, blockz)
	gridz := uint32(min3(int32(ratioz), int32(ratioz2), blocksperdim[2]))
	return Config3d{
		Dimx:            xdim,
		Dimy:            ydim,
		Dimz:            zdim,
		ThreadPerBlockx: (uint32)(blockx),
		ThreadPerBlocky: (uint32)(blocky),
		ThreadPerBlockz: (uint32)(blockz),
		BlockCountx:     gridx,
		BlockCounty:     gridy,
		BlockCountz:     gridz,
	}, nil

}
func divideandroundup(den, num int32) uint32 {

	return uint32(((den - 1) / num) + 1)

}
func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
func max3(a, b, c int32) int32 {
	return max(max(a, b), c)
}
func min3(a, b, c int32) int32 {
	return min(min(a, b), c)
}
func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
