package types

type StringStatsAggregate struct {
	AvgLength         *Float64            `json:"avg_length,omitempty"`
	AvgLengthAsString *string             `json:"avg_length_as_string,omitempty"`
	Count             int64               `json:"count"`
	Distribution      *map[string]Float64 `json:"distribution,omitempty"`
	Entropy           *Float64            `json:"entropy,omitempty"`
	MaxLength         *int                `json:"max_length,omitempty"`
	MaxLengthAsString *string             `json:"max_length_as_string,omitempty"`
	Meta              Metadata            `json:"meta,omitempty"`
	MinLength         *int                `json:"min_length,omitempty"`
	MinLengthAsString *string             `json:"min_length_as_string,omitempty"`
}

func (s *StringStatsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStringStatsAggregate() *StringStatsAggregate { _ = "STUB: not implemented"; return nil }
