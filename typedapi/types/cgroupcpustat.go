package types

type CgroupCpuStat struct {
	NumberOfElapsedPeriods *int64 `json:"number_of_elapsed_periods,omitempty"`

	NumberOfTimesThrottled *int64 `json:"number_of_times_throttled,omitempty"`

	TimeThrottledNanos *int64 `json:"time_throttled_nanos,omitempty"`
}

func (s *CgroupCpuStat) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCgroupCpuStat() *CgroupCpuStat { _ = "STUB: not implemented"; return nil }
