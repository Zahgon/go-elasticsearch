package types

type CgroupCpu struct {
	CfsPeriodMicros *int `json:"cfs_period_micros,omitempty"`

	CfsQuotaMicros *int `json:"cfs_quota_micros,omitempty"`

	ControlGroup *string `json:"control_group,omitempty"`

	Stat *CgroupCpuStat `json:"stat,omitempty"`
}

func (s *CgroupCpu) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCgroupCpu() *CgroupCpu { _ = "STUB: not implemented"; return nil }
