package hip

//#include <hip/hip_runtime_api.h>
import "C"

type ArrayFlag C.uint

func (a ArrayFlag) c() C.uint      { return (C.uint)(a) }
func (a *ArrayFlag) cptr() *C.uint { return (*C.uint)(a) }

func (a *ArrayFlag) Default() ArrayFlag { *a = (ArrayFlag)(C.hipArrayDefault); return *a }
func (a *ArrayFlag) Layered() ArrayFlag { *a = (ArrayFlag)(C.hipArrayLayered); return *a }
func (a *ArrayFlag) SurfaceLoadStore() ArrayFlag {
	*a = (ArrayFlag)(C.hipArraySurfaceLoadStore)
	return *a
}
func (a *ArrayFlag) Cubemap() ArrayFlag       { *a = (ArrayFlag)(C.hipArrayCubemap); return *a }
func (a *ArrayFlag) TextureGather() ArrayFlag { *a = (ArrayFlag)(C.hipArrayTextureGather); return *a }
