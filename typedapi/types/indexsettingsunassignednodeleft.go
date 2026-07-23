package types

type IndexSettingsUnassignedNodeLeft struct {
	DelayedTimeout Duration `json:"delayed_timeout,omitempty"`
}

func (s *IndexSettingsUnassignedNodeLeft) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexSettingsUnassignedNodeLeft() *IndexSettingsUnassignedNodeLeft {
	_ = "STUB: not implemented"
	return nil
}

type IndexSettingsUnassignedNodeLeftVariant interface {
	IndexSettingsUnassignedNodeLeftCaster() *IndexSettingsUnassignedNodeLeft
}

func (s *IndexSettingsUnassignedNodeLeft) IndexSettingsUnassignedNodeLeftCaster() *IndexSettingsUnassignedNodeLeft {
	_ = "STUB: not implemented"
	return nil
}
