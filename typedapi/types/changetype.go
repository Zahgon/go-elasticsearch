package types

type ChangeType struct {
	Dip                *Dip                `json:"dip,omitempty"`
	DistributionChange *DistributionChange `json:"distribution_change,omitempty"`
	Indeterminable     *Indeterminable     `json:"indeterminable,omitempty"`
	NonStationary      *NonStationary      `json:"non_stationary,omitempty"`
	Spike              *Spike              `json:"spike,omitempty"`
	Stationary         *Stationary         `json:"stationary,omitempty"`
	StepChange         *StepChange         `json:"step_change,omitempty"`
	TrendChange        *TrendChange        `json:"trend_change,omitempty"`
}

func NewChangeType() *ChangeType { _ = "STUB: not implemented"; return nil }
