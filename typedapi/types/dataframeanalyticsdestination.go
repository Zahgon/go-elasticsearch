package types

type DataframeAnalyticsDestination struct {
	Index string `json:"index"`

	ResultsField *string `json:"results_field,omitempty"`
}

func (s *DataframeAnalyticsDestination) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsDestination() *DataframeAnalyticsDestination {
	_ = "STUB: not implemented"
	return nil
}

type DataframeAnalyticsDestinationVariant interface {
	DataframeAnalyticsDestinationCaster() *DataframeAnalyticsDestination
}

func (s *DataframeAnalyticsDestination) DataframeAnalyticsDestinationCaster() *DataframeAnalyticsDestination {
	_ = "STUB: not implemented"
	return nil
}
