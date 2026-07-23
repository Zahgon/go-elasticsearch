package types

type BoxPlotAggregate struct {
	Lower         Float64  `json:"lower"`
	LowerAsString *string  `json:"lower_as_string,omitempty"`
	Max           Float64  `json:"max"`
	MaxAsString   *string  `json:"max_as_string,omitempty"`
	Meta          Metadata `json:"meta,omitempty"`
	Min           Float64  `json:"min"`
	MinAsString   *string  `json:"min_as_string,omitempty"`
	Q1            Float64  `json:"q1"`
	Q1AsString    *string  `json:"q1_as_string,omitempty"`
	Q2            Float64  `json:"q2"`
	Q2AsString    *string  `json:"q2_as_string,omitempty"`
	Q3            Float64  `json:"q3"`
	Q3AsString    *string  `json:"q3_as_string,omitempty"`
	Upper         Float64  `json:"upper"`
	UpperAsString *string  `json:"upper_as_string,omitempty"`
}

func (s *BoxPlotAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBoxPlotAggregate() *BoxPlotAggregate { _ = "STUB: not implemented"; return nil }
