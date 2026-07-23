package types

type Cgroup struct {
	Cpu *CgroupCpu `json:"cpu,omitempty"`

	Cpuacct *CpuAcct `json:"cpuacct,omitempty"`

	Memory *CgroupMemory `json:"memory,omitempty"`
}

func NewCgroup() *Cgroup { _ = "STUB: not implemented"; return nil }
