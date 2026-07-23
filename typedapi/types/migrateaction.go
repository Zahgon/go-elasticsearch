package types

type MigrateAction struct {
	Enabled *bool `json:"enabled,omitempty"`
}

func (s *MigrateAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMigrateAction() *MigrateAction { _ = "STUB: not implemented"; return nil }

type MigrateActionVariant interface {
	MigrateActionCaster() *MigrateAction
}

func (s *MigrateAction) MigrateActionCaster() *MigrateAction { _ = "STUB: not implemented"; return nil }
