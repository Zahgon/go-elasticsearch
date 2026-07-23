package types

type AnomalyCause struct {
	Actual                 []Float64   `json:"actual,omitempty"`
	ByFieldName            *string     `json:"by_field_name,omitempty"`
	ByFieldValue           *string     `json:"by_field_value,omitempty"`
	CorrelatedByFieldValue *string     `json:"correlated_by_field_value,omitempty"`
	FieldName              *string     `json:"field_name,omitempty"`
	Function               *string     `json:"function,omitempty"`
	FunctionDescription    *string     `json:"function_description,omitempty"`
	GeoResults             *GeoResults `json:"geo_results,omitempty"`
	Influencers            []Influence `json:"influencers,omitempty"`
	OverFieldName          *string     `json:"over_field_name,omitempty"`
	OverFieldValue         *string     `json:"over_field_value,omitempty"`
	PartitionFieldName     *string     `json:"partition_field_name,omitempty"`
	PartitionFieldValue    *string     `json:"partition_field_value,omitempty"`
	Probability            Float64     `json:"probability"`
	Typical                []Float64   `json:"typical,omitempty"`
}

func (s *AnomalyCause) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnomalyCause() *AnomalyCause { _ = "STUB: not implemented"; return nil }
