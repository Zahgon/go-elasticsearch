package types

type DeleteAction struct {
	DeleteSearchableSnapshot *bool `json:"delete_searchable_snapshot,omitempty"`
}

func (s *DeleteAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDeleteAction() *DeleteAction { _ = "STUB: not implemented"; return nil }

type DeleteActionVariant interface {
	DeleteActionCaster() *DeleteAction
}

func (s *DeleteAction) DeleteActionCaster() *DeleteAction { _ = "STUB: not implemented"; return nil }
