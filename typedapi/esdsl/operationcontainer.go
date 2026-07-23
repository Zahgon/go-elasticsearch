package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _operationContainer struct {
	v *types.OperationContainer
}

func NewOperationContainer() *_operationContainer { _ = "STUB: not implemented"; return nil }

func (s *_operationContainer) Create(create types.CreateOperationVariant) *_operationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_operationContainer) Delete(delete types.DeleteOperationVariant) *_operationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_operationContainer) Index(index types.IndexOperationVariant) *_operationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_operationContainer) Update(update types.UpdateOperationVariant) *_operationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_operationContainer) OperationContainerCaster() *types.OperationContainer {
	_ = "STUB: not implemented"
	return nil
}
