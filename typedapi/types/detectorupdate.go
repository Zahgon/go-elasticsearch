package types

type DetectorUpdate struct {
	CustomRules []DetectionRule `json:"custom_rules,omitempty"`

	Description *string `json:"description,omitempty"`

	DetectorIndex int `json:"detector_index"`
}

func (s *DetectorUpdate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDetectorUpdate() *DetectorUpdate { _ = "STUB: not implemented"; return nil }

type DetectorUpdateVariant interface {
	DetectorUpdateCaster() *DetectorUpdate
}

func (s *DetectorUpdate) DetectorUpdateCaster() *DetectorUpdate {
	_ = "STUB: not implemented"
	return nil
}
