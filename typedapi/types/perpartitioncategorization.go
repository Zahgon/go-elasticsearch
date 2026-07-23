package types

type PerPartitionCategorization struct {
	Enabled *bool `json:"enabled,omitempty"`

	StopOnWarn *bool `json:"stop_on_warn,omitempty"`
}

func (s *PerPartitionCategorization) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPerPartitionCategorization() *PerPartitionCategorization {
	_ = "STUB: not implemented"
	return nil
}

type PerPartitionCategorizationVariant interface {
	PerPartitionCategorizationCaster() *PerPartitionCategorization
}

func (s *PerPartitionCategorization) PerPartitionCategorizationCaster() *PerPartitionCategorization {
	_ = "STUB: not implemented"
	return nil
}
