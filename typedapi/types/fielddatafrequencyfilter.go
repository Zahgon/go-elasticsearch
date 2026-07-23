package types

type FielddataFrequencyFilter struct {
	Max            Float64 `json:"max"`
	Min            Float64 `json:"min"`
	MinSegmentSize int     `json:"min_segment_size"`
}

func (s *FielddataFrequencyFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFielddataFrequencyFilter() *FielddataFrequencyFilter { _ = "STUB: not implemented"; return nil }

type FielddataFrequencyFilterVariant interface {
	FielddataFrequencyFilterCaster() *FielddataFrequencyFilter
}

func (s *FielddataFrequencyFilter) FielddataFrequencyFilterCaster() *FielddataFrequencyFilter {
	_ = "STUB: not implemented"
	return nil
}
