package hip

//#include <hip/hip_runtime_api.h>
import "C"

//DeviceProp is a binding for C.hipDeviceProp_t.  Subvalues are accessed through DeviceProp's methods
type DeviceProp C.hipDeviceProp_t

func (d DeviceProp) c() C.hipDeviceProp_t      { return (C.hipDeviceProp_t)(d) }
func (d *DeviceProp) cptr() *C.hipDeviceProp_t { return (*C.hipDeviceProp_t)(d) }

//Name - Device name
func (d *DeviceProp) Name() string {
	return C.GoString(&d.name[0])
}

//Arch returns the DeviceArch struct.
func (d *DeviceProp) Arch() DeviceArch {
	return (DeviceArch)(d.arch)
}

//TotalGlobalMem - size of globalmemory region in bytes
func (d *DeviceProp) TotalGlobalMem() uint {
	return (uint)(d.totalGlobalMem)
}

//SharedMemPerBlock -Size of shared memory region (in bytes)
func (d *DeviceProp) SharedMemPerBlock() uint {
	return (uint)(d.sharedMemPerBlock)
}

//MaxSharedMemoryPerMultiProcessor -  Maximum Shared Memory Per Multiprocessor.
func (d *DeviceProp) MaxSharedMemoryPerMultiProcessor() uint {
	return (uint)(d.maxSharedMemoryPerMultiProcessor)
}

//TotalConstMem - Size of shared memory region (in bytes).
func (d *DeviceProp) TotalConstMem() uint {
	return (uint)(d.totalConstMem)
}

//RegsPerBlock - Registers per block.
func (d *DeviceProp) RegsPerBlock() int32 {
	return (int32)(d.regsPerBlock)
}

//WarpSize -  Warp size.
func (d *DeviceProp) WarpSize() int32 {
	return (int32)(d.warpSize)
}

//MaxThreadsPerBlock - Max work items per work group or workgroup max size.
func (d *DeviceProp) MaxThreadsPerBlock() int32 {
	return (int32)(d.maxThreadsPerBlock)
}

//MaxThreadsDim - Max number of threads in each dimension (XYZ) of a block.
func (d *DeviceProp) MaxThreadsDim() []int32 {
	x := make([]int32, 3)
	for i := 0; i < 3; i++ {
		x[0] = (int32)(d.maxThreadsDim[i])
	}
	return x
}

//MaxGridSize - Max grid dimensions (XYZ).
func (d *DeviceProp) MaxGridSize() []int32 {
	x := make([]int32, 3)
	for i := 0; i < 3; i++ {
		x[0] = (int32)(d.maxGridSize[i])
	}
	return x
}

//ClockRate - Max clock frequency of the multiProcessors in khz.
func (d *DeviceProp) ClockRate() int32 {
	return (int32)(d.clockRate)
}

//MemoryClockRate -Max global memory clock frequency in khz.
func (d *DeviceProp) MemoryClockRate() int32 {
	return (int32)(d.memoryClockRate)
}

//MemoryBusWidth -G lobal memory bus width in bits.
func (d *DeviceProp) MemoryBusWidth() int32 {
	return (int32)(d.memoryBusWidth)
}

//Major - Major compute capability.  On HCC, this is an approximation and features may
//differ from CUDA CC.  See the arch feature flags for portable ways to query
//feature caps.
func (d *DeviceProp) Major() int32 {
	return (int32)(d.major)
}

//Minor - Minor compute capability.  On HCC, this is an approximation and features may
//differ from CUDA CC.  See the arch feature flags for portable ways to query
//feature caps.
func (d *DeviceProp) Minor() int32 {
	return (int32)(d.minor)
}

//MultiProcessorCount - Number of multi-processors (compute units).
func (d *DeviceProp) MultiProcessorCount() int32 {
	return (int32)(d.multiProcessorCount)
}

//L2CacheSize - L2 cache size.
func (d *DeviceProp) L2CacheSize() int32 {
	return (int32)(d.l2CacheSize)
}

//MaxThreadsPerMultiProcessor - Maximum resident threads per multi-processor.
func (d *DeviceProp) MaxThreadsPerMultiProcessor() int32 {
	return (int32)(d.maxThreadsPerMultiProcessor)
}

//ComputeMode - Compute mode.
func (d *DeviceProp) ComputeMode() int32 {
	return (int32)(d.computeMode)
}

//ClockInstructionRate - Frequency in khz of the timer used by the device-side "clock*"
//instructions.  New for HIP.
func (d *DeviceProp) ClockInstructionRate() int32 {
	return (int32)(d.clockInstructionRate)
}

//ConcurrentKernels - Device can possibly execute multiple kernels concurrently.
func (d *DeviceProp) ConcurrentKernels() int32 {
	return (int32)(d.concurrentKernels)
}

//PciDomainID - PCI Domain ID
func (d *DeviceProp) PciDomainID() int32 {
	return (int32)(d.pciDomainID)
}

//PciBusID - PCI Bus ID.
func (d *DeviceProp) PciBusID() int32 {
	return (int32)(d.pciBusID)
}

//PciDeviceID - PCI Device ID.
func (d *DeviceProp) PciDeviceID() int32 {
	return (int32)(d.pciDeviceID)
}

//IsMultiGpuBoard -  1 if device is on a multi-GPU board, 0 if not.
func (d *DeviceProp) IsMultiGpuBoard() int32 {
	return (int32)(d.isMultiGpuBoard)
}

//CanMapHostMemory - Check whether HIP can map host memory
func (d *DeviceProp) CanMapHostMemory() int32 {
	return (int32)(d.canMapHostMemory)
}

//GcnArch - AMD GCN Arch Value. Eg: 803, 701
func (d *DeviceProp) GcnArch() int32 {
	return (int32)(d.gcnArch)
}

//Integrated - APU vs dGPU
func (d *DeviceProp) Integrated() int32 {
	return (int32)(d.integrated)
}

//CooperativeLaunch - HIP device supports cooperative launch
func (d *DeviceProp) CooperativeLaunch() int32 {
	return (int32)(d.cooperativeLaunch)
}

//CooperativeMultiDeviceLaunch - HIP device supports cooperative launch on multiple devices
func (d *DeviceProp) CooperativeMultiDeviceLaunch() int32 {
	return (int32)(d.cooperativeMultiDeviceLaunch)
}

//DeviceArch contains bit flags.
//
//TODO: Will have to write a tiny c function to access
type DeviceArch C.hipDeviceArch_t

/*
//GlobalInt32Atomics - 32-bit integer atomics for global memory.
func (d DeviceArch) GlobalInt32Atomics() bool { return false }

//GlobalFloatAtomicExch - 32-bit float atomic exch for global memory.
func (d DeviceArch) GlobalFloatAtomicExch() bool { return false }

//SharedInt32Atomics - 32-bit integer atomics for shared memory.
func (d DeviceArch) SharedInt32Atomics() bool { return false }

//SharedFloatAtomicExch - 32-bit float atomic exch for shared memory.
func (d DeviceArch) SharedFloatAtomicExch() bool { return false }

//FloatAtomicAdd - 32-bit float atomic add in global and shared memory.
func (d DeviceArch) FloatAtomicAdd() bool { return false }

//GlobalInt64Atomics - 64-bit integer atomics for global memory.
func (d DeviceArch) GlobalInt64Atomics() bool { return false }

//SharedInt64Atomics - 64-bit integer atomics for shared memory.
func (d DeviceArch) SharedInt64Atomics() bool { return false }

//Doubles - Double-precision floating point.
func (d DeviceArch) Doubles() bool { return false }

//WarpVote - Warp vote instructions (__any, __all).
func (d DeviceArch) WarpVote() bool { return false }

//WarpBallot - Warp ballot instructions (__ballot).
func (d DeviceArch) WarpBallot() bool { return false }

//WarpShuffle - Warp shuffle operations. (__shfl_*).
func (d DeviceArch) WarpShuffle() bool { return false }

//FunnelShift - Funnel two words into one with shift&mask caps.
func (d DeviceArch) FunnelShift() bool { return false }

//ThreadFenceSystem - __threadfence_system.
func (d DeviceArch) ThreadFenceSystem() bool { return false }

//SyncThreadsExt - __syncthreads_count, syncthreads_and, syncthreads_or.
func (d DeviceArch) SyncThreadsExt() bool { return false }

//SurfaceFuncs - Surface functions.
func (d DeviceArch) SurfaceFuncs() bool { return false }

//Grid3D - Grid and group dims are 3D (rather than 2D).
func (d DeviceArch) Grid3D() bool { return false }

//DynamicParallelism - Dynamic parallelism.
func (d DeviceArch) DynamicParallelism() bool { return false }
*/
