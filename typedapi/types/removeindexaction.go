package types

type RemoveIndexAction struct {
	Index *string `json:"index,omitempty"`

	Indices []string `json:"indices,omitempty"`

	MustExist *bool `json:"must_exist,omitempty"`
}

func (s *RemoveIndexAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRemoveIndexAction() *RemoveIndexAction { _ = "STUB: not implemented"; return nil }

type RemoveIndexActionVariant interface {
	RemoveIndexActionCaster() *RemoveIndexAction
}

func (s *RemoveIndexAction) RemoveIndexActionCaster() *RemoveIndexAction {
	_ = "STUB: not implemented"
	return nil
}
