package hip

//#include <hip/hip_runtime_api.h>
import "C"

type MallocFlags C.uint

func (m *MallocFlags) Default() MallocFlags {
	*m = (C.hipHostMallocDefault)
	return *m
}
func (m *MallocFlags) Portable() MallocFlags {
	*m = (C.hipHostMallocPortable)
	return *m
}

func (m *MallocFlags) Mapped() MallocFlags {
	*m = (C.hipHostMallocMapped)
	return *m
}

func (m *MallocFlags) WriteCombined() MallocFlags {
	*m = (C.hipHostMallocWriteCombined)
	return *m
}

func (m *MallocFlags) Coherent() MallocFlags {
	*m = (C.hipHostMallocCoherent)
	return *m
}

func (m *MallocFlags) NonCoherent() MallocFlags {
	*m = (C.hipHostMallocNonCoherent)
	return *m
}

func (m *MallocFlags) Global() MallocFlags {
	*m = (C.hipMemAttachGlobal)
	return *m
}

func (m *MallocFlags) AttachHost() MallocFlags {
	*m = (C.hipMemAttachHost)
	return *m
}

func (m *MallocFlags) DeviceDefault() MallocFlags {
	*m = (C.hipDeviceMallocDefault)
	return *m
}

func (m *MallocFlags) DeviceFinegrained() MallocFlags {
	*m = (C.hipDeviceMallocFinegrained)
	return *m
}
