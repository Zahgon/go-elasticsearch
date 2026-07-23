package types

type RemoveAction struct {
	Alias *string `json:"alias,omitempty"`

	Aliases []string `json:"aliases,omitempty"`

	Index *string `json:"index,omitempty"`

	Indices []string `json:"indices,omitempty"`

	MustExist *bool `json:"must_exist,omitempty"`
}

func (s *RemoveAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRemoveAction() *RemoveAction { _ = "STUB: not implemented"; return nil }

type RemoveActionVariant interface {
	RemoveActionCaster() *RemoveAction
}

func (s *RemoveAction) RemoveActionCaster() *RemoveAction { _ = "STUB: not implemented"; return nil }
