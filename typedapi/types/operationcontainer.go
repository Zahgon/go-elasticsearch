package types

type OperationContainer struct {
	Create *CreateOperation `json:"create,omitempty"`

	Delete *DeleteOperation `json:"delete,omitempty"`

	Index *IndexOperation `json:"index,omitempty"`

	Update *UpdateOperation `json:"update,omitempty"`
}

func NewOperationContainer() *OperationContainer { _ = "STUB: not implemented"; return nil }

type OperationContainerVariant interface {
	OperationContainerCaster() *OperationContainer
}

func (s *OperationContainer) OperationContainerCaster() *OperationContainer {
	_ = "STUB: not implemented"
	return nil
}
