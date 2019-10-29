package hip

/*
#cgo CFLAGS: -D__HIP_PLATFORM_HCC__ -D__HIP_VDI__
#cgo CFLAGS: -I/opt/rocm/hip/include
#cgo CXXFLAGS: -D__HIP_PLATFORM_HCC__ -D__HIP_VDI__
#cgo CXXFLAGS:-I/opt/rocm/hip/include
#cgo LDFLAGS: -L/opt/rocm/hip/lib -lhip_hcc
*/
import "C"
