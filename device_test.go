package hip

import (
	"fmt"
	"testing"
)

func Test_hipGetDevice(t *testing.T) {
	d, err := GetDevice()
	if err != nil {
		t.Error(err)
	}
	prop, err := GetDeviceProperties(d)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(prop.Name())
	fmt.Println(prop.ConcurrentKernels())
	fmt.Println(prop.GcnArch())
	fmt.Println(prop.TotalConstMem())
	fmt.Println(prop.TotalGlobalMem())
	fmt.Println(prop.ComputeMode())
	t.Error(prop)

}
