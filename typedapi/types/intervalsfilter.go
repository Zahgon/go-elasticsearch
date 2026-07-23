package types

type IntervalsFilter struct {
	After *Intervals `json:"after,omitempty"`

	Before *Intervals `json:"before,omitempty"`

	ContainedBy *Intervals `json:"contained_by,omitempty"`

	Containing *Intervals `json:"containing,omitempty"`

	NotContainedBy *Intervals `json:"not_contained_by,omitempty"`

	NotContaining *Intervals `json:"not_containing,omitempty"`

	NotOverlapping *Intervals `json:"not_overlapping,omitempty"`

	Overlapping *Intervals `json:"overlapping,omitempty"`

	Script *Script `json:"script,omitempty"`
}

func NewIntervalsFilter() *IntervalsFilter { _ = "STUB: not implemented"; return nil }

type IntervalsFilterVariant interface {
	IntervalsFilterCaster() *IntervalsFilter
}

func (s *IntervalsFilter) IntervalsFilterCaster() *IntervalsFilter {
	_ = "STUB: not implemented"
	return nil
}
