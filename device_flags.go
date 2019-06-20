package hip

//#include <hip/hip_runtime_api.h>
import "C"

//FuncCache - On AMD devices and some Nvidia devices, these hints and controls are ignored.
type FuncCache C.hipFuncCache_t

func (f FuncCache) c() C.hipFuncCache_t {
	return (C.hipFuncCache_t)(f)
}
func (f *FuncCache) cptr() *C.hipFuncCache_t {
	return (*C.hipFuncCache_t)(f)
}

//PreferNone -  no preference for shared memory or L1 (default)
func (f *FuncCache) PreferNone() FuncCache { *f = (FuncCache)(C.hipFuncCachePreferNone); return *f }

//PreferShared - prefer larger shared memory and smaller L1 cache
func (f *FuncCache) PreferShared() FuncCache { *f = (FuncCache)(C.hipFuncCachePreferShared); return *f }

//PreferL1 -  prefer larger L1 cache and smaller shared memory
func (f *FuncCache) PreferL1() FuncCache { *f = (FuncCache)(C.hipFuncCachePreferL1); return *f }

//PreferEqual -  prefer equal size L1 cache and shared memory
func (f *FuncCache) PreferEqual() FuncCache { *f = (FuncCache)(C.hipFuncCachePreferEqual); return *f }

//Limit holds flags for device limits and are exposed through its methods
type Limit uint32

func (l Limit) c() uint32 {
	return (uint32)(l)
}

//MallocHeapSize sets l to MallocHeapSize and returns value of l
func (l *Limit) MallocHeapSize() Limit { *l = (Limit)(C.hipLimitMallocHeapSize); return *l }

//SharedMemConfig holds flags that can be set through its methods
//
//On AMD devices and some Nvidia devices, these hints and controls are ignored.
type SharedMemConfig C.hipSharedMemConfig

func (s SharedMemConfig) c() C.hipSharedMemConfig      { return (C.hipSharedMemConfig)(s) }
func (s *SharedMemConfig) cptr() *C.hipSharedMemConfig { return (*C.hipSharedMemConfig)(s) }

//Default changeds s to Default and returns value of s
//
//The compiler selects a device-specific value for the banking.
func (s *SharedMemConfig) Default() SharedMemConfig {
	*s = (SharedMemConfig)(C.hipSharedMemBankSizeDefault)
	return *s
}

//FourByte changeds s to FourByte and returns value of s
//
//Shared mem is banked at 4-bytes intervals and performs best
//when adjacent threads access data 4 bytes apart.
func (s *SharedMemConfig) FourByte() SharedMemConfig {
	*s = (SharedMemConfig)(C.hipSharedMemBankSizeFourByte)
	return *s
}

//EightByte changeds s to EightByte and returns value of s
//
//Shared mem is banked at 8-byte intervals and performs best
//when adjacent threads access data 4 bytes apart.
func (s *SharedMemConfig) EightByte() SharedMemConfig {
	*s = (SharedMemConfig)(C.hipSharedMemBankSizeEightByte)
	return *s
}

//DeviceFlag contains deviceflags that are exposed through its methods
type DeviceFlag C.uint

func (d DeviceFlag) c() C.uint      { return (C.uint)(d) }
func (d *DeviceFlag) cptr() *C.uint { return (*C.uint)(d) }

//ScheduleAuto changes d to ScheduleAuto and returns value of d.
func (d *DeviceFlag) ScheduleAuto() DeviceFlag { *d = (DeviceFlag)(C.hipDeviceScheduleAuto); return *d }

//ScheduleSpin changes d to ScheduleSpin and returns value of d.
func (d *DeviceFlag) ScheduleSpin() DeviceFlag { *d = (DeviceFlag)(C.hipDeviceScheduleSpin); return *d }

//ScheduleYield changes d to ScheduleYield and returns value of d.
func (d *DeviceFlag) ScheduleYield() DeviceFlag {
	*d = (DeviceFlag)(C.hipDeviceScheduleYield)
	return *d
}

//ScheduleBlockingSync changes d to ScheduleBlockingSync and returns value of d.
func (d *DeviceFlag) ScheduleBlockingSync() DeviceFlag {
	*d = (DeviceFlag)(C.hipDeviceScheduleBlockingSync)
	return *d
}

//ScheduleMask changes d to ScheduleMask and returns value of d.
func (d *DeviceFlag) ScheduleMask() DeviceFlag { *d = (DeviceFlag)(C.hipDeviceScheduleMask); return *d }

//MapHost changes d to MapHost and returns value of d.
func (d *DeviceFlag) MapHost() DeviceFlag { *d = (DeviceFlag)(C.hipDeviceMapHost); return *d }

//LmemResizeToMax changes d to LmemResizeToMax and returns value of d.
func (d *DeviceFlag) LmemResizeToMax() DeviceFlag {
	*d = (DeviceFlag)(C.hipDeviceLmemResizeToMax)
	return *d
}

//DeviceAttribute are flags that are changed through its methods
type DeviceAttribute C.hipDeviceAttribute_t

func (d DeviceAttribute) c() C.hipDeviceAttribute_t {
	return (C.hipDeviceAttribute_t)(d)
}
func (d *DeviceAttribute) cptr() *C.hipDeviceAttribute_t {
	return (*C.hipDeviceAttribute_t)(d)
}

//MaxThreadsPerBlock changes d to MaxThreadsPerBlock and returns value of d.
func (d *DeviceAttribute) MaxThreadsPerBlock() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxThreadsPerBlock)
	return *d
}

//MaxBlockDimX changes d to MaxBlockDimX and returns value of d.
func (d *DeviceAttribute) MaxBlockDimX() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxBlockDimX)
	return *d
}

//MaxBlockDimY changes d to MaxBlockDimY and returns value of d.
func (d *DeviceAttribute) MaxBlockDimY() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxBlockDimY)
	return *d
}

//MaxBlockDimZ changes d to MaxBlockDimZ and returns value of d.
func (d *DeviceAttribute) MaxBlockDimZ() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxBlockDimZ)
	return *d
}

//MaxGridDimX changes d to MaxGridDimX and returns value of d.
func (d *DeviceAttribute) MaxGridDimX() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxGridDimX)
	return *d
}

//MaxGridDimY changes d to MaxGridDimY and returns value of d.
func (d *DeviceAttribute) MaxGridDimY() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxGridDimY)
	return *d
}

//MaxGridDimZ changes d to MaxGridDimZ and returns value of d.
func (d *DeviceAttribute) MaxGridDimZ() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxGridDimZ)
	return *d
}

//MaxSharedMemoryPerBlock changes d to d to MaxSharedMemoryPerBlock and returns value of d.
func (d *DeviceAttribute) MaxSharedMemoryPerBlock() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxSharedMemoryPerBlock)
	return *d
}

//TotalConstantMemory changes d to TotalConstantMemory and returns value of d.
func (d *DeviceAttribute) TotalConstantMemory() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeTotalConstantMemory)
	return *d
}

//WarpSize changes d to WarpSize and returns value of d.
func (d *DeviceAttribute) WarpSize() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeWarpSize)
	return *d
}

//MaxRegistersPerBlock changes d d to MaxRegistersPerBlock and returns value of d.
func (d *DeviceAttribute) MaxRegistersPerBlock() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxRegistersPerBlock)
	return *d
}

//ClockRate changes d to ClockRate and returns value of d.
func (d *DeviceAttribute) ClockRate() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeClockRate)
	return *d
}

//MemoryClockRate changes d to MemoryClockRate and returns value of d.
func (d *DeviceAttribute) MemoryClockRate() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMemoryClockRate)
	return *d
}

//MemoryBusWidth changes d to MemoryBusWidth and returns value of d.
func (d *DeviceAttribute) MemoryBusWidth() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMemoryBusWidth)
	return *d
}

//MultiprocessorCount changes  d to MultiprocessorCount and returns value of d.
func (d *DeviceAttribute) MultiprocessorCount() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMultiprocessorCount)
	return *d
}

//ComputeMode changes d to ComputeMode and returns value of d.
func (d *DeviceAttribute) ComputeMode() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeComputeMode)
	return *d
}

//L2CacheSize changes d to L2CacheSize and returns value of d.
func (d *DeviceAttribute) L2CacheSize() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeL2CacheSize)
	return *d
}

//MaxThreadsPerMultiProcessor changes d to Max d to MaxThreadsPerMultiProcessor and returns value of d.
func (d *DeviceAttribute) MaxThreadsPerMultiProcessor() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxThreadsPerMultiProcessor)
	return *d
}

//ComputeCapabilityMajor changes d t d to ComputeCapabilityMajor and returns value of d.
func (d *DeviceAttribute) ComputeCapabilityMajor() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeComputeCapabilityMajor)
	return *d
}

//ComputeCapabilityMinor changes d t d to ComputeCapabilityMinor and returns value of d.
func (d *DeviceAttribute) ComputeCapabilityMinor() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeComputeCapabilityMinor)
	return *d
}

//ConcurrentKernels changes d to ConcurrentKernels and returns value of d.
func (d *DeviceAttribute) ConcurrentKernels() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeConcurrentKernels)
	return *d
}

//PciBusID changes d to PciBusID and returns value of d.
func (d *DeviceAttribute) PciBusID() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributePciBusId)
	return *d
}

//PciDeviceID changes d to PciDeviceID and returns value of d.
func (d *DeviceAttribute) PciDeviceID() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributePciDeviceId)
	return *d
}

//MaxSharedMemoryPerMultiprocessor changes d to MaxShare d to MaxSharedMemoryPerMultiprocessor and returns value of d.
func (d *DeviceAttribute) MaxSharedMemoryPerMultiprocessor() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeMaxSharedMemoryPerMultiprocessor)
	return *d
}

//IsMultiGpuBoard changes d to IsMultiGpuBoard and returns value of d.
func (d *DeviceAttribute) IsMultiGpuBoard() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeIsMultiGpuBoard)
	return *d
}

//Integrated changes d to Integrated and returns value of d.
func (d *DeviceAttribute) Integrated() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeIntegrated)
	return *d
}

//CooperativeLaunch changes d to CooperativeLaunch and returns value of d.
func (d *DeviceAttribute) CooperativeLaunch() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeCooperativeLaunch)
	return *d
}

//CooperativeMultiDeviceLaunch changes d to Coop d to CooperativeMultiDeviceLaunch and returns value of d.
func (d *DeviceAttribute) CooperativeMultiDeviceLaunch() DeviceAttribute {
	*d = (DeviceAttribute)(C.hipDeviceAttributeCooperativeMultiDeviceLaunch)
	return *d
}
