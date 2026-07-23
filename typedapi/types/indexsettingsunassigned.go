package types

type IndexSettingsUnassigned struct {
	NodeLeft *IndexSettingsUnassignedNodeLeft `json:"node_left,omitempty"`
}

func NewIndexSettingsUnassigned() *IndexSettingsUnassigned { _ = "STUB: not implemented"; return nil }

type IndexSettingsUnassignedVariant interface {
	IndexSettingsUnassignedCaster() *IndexSettingsUnassigned
}

func (s *IndexSettingsUnassigned) IndexSettingsUnassignedCaster() *IndexSettingsUnassigned {
	_ = "STUB: not implemented"
	return nil
}
