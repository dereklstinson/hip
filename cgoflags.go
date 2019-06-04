package hipgo

/*
#cgo CFLAGS: -D__HIP_PLATFORM_HCC__ -D__HIP_VDI__ -D__HCC_C__
#cgo CFLAGS: -I/home/derek/amd/hip/include
#cgo CXXFLAGS: -D__HIP_PLATFORM_HCC__ -D__HIP_VDI__
#cgo CXXFLAGS: -I/home/derek/amd/hip/include
#cgo LDFLAGS: -L/home/derek/amd/hip/lib -lhip_hcc
*/
import "C"
