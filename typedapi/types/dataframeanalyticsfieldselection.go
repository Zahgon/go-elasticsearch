package types

type DataframeAnalyticsFieldSelection struct {
	FeatureType *string `json:"feature_type,omitempty"`

	IsIncluded bool `json:"is_included"`

	IsRequired bool `json:"is_required"`

	MappingTypes []string `json:"mapping_types"`

	Name string `json:"name"`

	Reason *string `json:"reason,omitempty"`
}

func (s *DataframeAnalyticsFieldSelection) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsFieldSelection() *DataframeAnalyticsFieldSelection {
	_ = "STUB: not implemented"
	return nil
}
