package types

type MinHashTokenFilter struct {
	BucketCount *int `json:"bucket_count,omitempty"`

	HashCount *int `json:"hash_count,omitempty"`

	HashSetSize *int    `json:"hash_set_size,omitempty"`
	Type        string  `json:"type,omitempty"`
	Version     *string `json:"version,omitempty"`

	WithRotation *bool `json:"with_rotation,omitempty"`
}

func (s *MinHashTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s MinHashTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMinHashTokenFilter() *MinHashTokenFilter { _ = "STUB: not implemented"; return nil }

type MinHashTokenFilterVariant interface {
	MinHashTokenFilterCaster() *MinHashTokenFilter
}

func (s *MinHashTokenFilter) MinHashTokenFilterCaster() *MinHashTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *MinHashTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
