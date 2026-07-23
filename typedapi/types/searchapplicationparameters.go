package types

type SearchApplicationParameters struct {
	AnalyticsCollectionName *string `json:"analytics_collection_name,omitempty"`

	Indices []string `json:"indices"`

	Template *SearchApplicationTemplate `json:"template,omitempty"`
}

func (s *SearchApplicationParameters) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSearchApplicationParameters() *SearchApplicationParameters {
	_ = "STUB: not implemented"
	return nil
}

type SearchApplicationParametersVariant interface {
	SearchApplicationParametersCaster() *SearchApplicationParameters
}

func (s *SearchApplicationParameters) SearchApplicationParametersCaster() *SearchApplicationParameters {
	_ = "STUB: not implemented"
	return nil
}
