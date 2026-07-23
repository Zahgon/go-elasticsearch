package types

type FieldStat struct {
	Cardinality int      `json:"cardinality"`
	Count       int      `json:"count"`
	Earliest    *string  `json:"earliest,omitempty"`
	Latest      *string  `json:"latest,omitempty"`
	MaxValue    *int     `json:"max_value,omitempty"`
	MeanValue   *int     `json:"mean_value,omitempty"`
	MedianValue *int     `json:"median_value,omitempty"`
	MinValue    *int     `json:"min_value,omitempty"`
	TopHits     []TopHit `json:"top_hits"`
}

func (s *FieldStat) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldStat() *FieldStat { _ = "STUB: not implemented"; return nil }
