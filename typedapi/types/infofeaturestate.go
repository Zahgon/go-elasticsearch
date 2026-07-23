package types

type InfoFeatureState struct {
	FeatureName string   `json:"feature_name"`
	Indices     []string `json:"indices"`
}

func (s *InfoFeatureState) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInfoFeatureState() *InfoFeatureState { _ = "STUB: not implemented"; return nil }
