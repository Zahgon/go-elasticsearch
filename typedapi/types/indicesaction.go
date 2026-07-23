package types

type IndicesAction struct {
	Add *AddAction `json:"add,omitempty"`

	Remove *RemoveAction `json:"remove,omitempty"`

	RemoveIndex *RemoveIndexAction `json:"remove_index,omitempty"`
}

func NewIndicesAction() *IndicesAction { _ = "STUB: not implemented"; return nil }

type IndicesActionVariant interface {
	IndicesActionCaster() *IndicesAction
}

func (s *IndicesAction) IndicesActionCaster() *IndicesAction { _ = "STUB: not implemented"; return nil }
