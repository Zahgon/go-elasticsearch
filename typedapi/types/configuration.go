package types

type Configuration struct {
	FeatureStates []string `json:"feature_states,omitempty"`

	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`

	IncludeGlobalState *bool `json:"include_global_state,omitempty"`

	Indices []string `json:"indices,omitempty"`

	Metadata Metadata `json:"metadata,omitempty"`

	Partial *bool `json:"partial,omitempty"`
}

func (s *Configuration) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewConfiguration() *Configuration { _ = "STUB: not implemented"; return nil }

type ConfigurationVariant interface {
	ConfigurationCaster() *Configuration
}

func (s *Configuration) ConfigurationCaster() *Configuration { _ = "STUB: not implemented"; return nil }
