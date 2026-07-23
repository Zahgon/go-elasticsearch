package types

type IndicesModifyAction struct {
	AddBackingIndex *IndexAndDataStreamAction `json:"add_backing_index,omitempty"`

	RemoveBackingIndex *IndexAndDataStreamAction `json:"remove_backing_index,omitempty"`
}

func NewIndicesModifyAction() *IndicesModifyAction { _ = "STUB: not implemented"; return nil }

type IndicesModifyActionVariant interface {
	IndicesModifyActionCaster() *IndicesModifyAction
}

func (s *IndicesModifyAction) IndicesModifyActionCaster() *IndicesModifyAction {
	_ = "STUB: not implemented"
	return nil
}
