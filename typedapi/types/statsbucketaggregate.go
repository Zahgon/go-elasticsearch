package types

type StatsBucketAggregate struct {
	Avg         *Float64 `json:"avg,omitempty"`
	AvgAsString *string  `json:"avg_as_string,omitempty"`
	Count       int64    `json:"count"`
	Max         *Float64 `json:"max,omitempty"`
	MaxAsString *string  `json:"max_as_string,omitempty"`
	Meta        Metadata `json:"meta,omitempty"`
	Min         *Float64 `json:"min,omitempty"`
	MinAsString *string  `json:"min_as_string,omitempty"`
	Sum         Float64  `json:"sum"`
	SumAsString *string  `json:"sum_as_string,omitempty"`
}

func (s *StatsBucketAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStatsBucketAggregate() *StatsBucketAggregate { _ = "STUB: not implemented"; return nil }
