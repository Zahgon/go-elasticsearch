package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _createOperation struct {
	v *types.CreateOperation
}

func NewCreateOperation() *_createOperation { _ = "STUB: not implemented"; return nil }

func (s *_createOperation) DynamicTemplates(dynamictemplates map[string]string) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) AddDynamicTemplate(key string, value string) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) Id_(id string) *_createOperation { _ = "STUB: not implemented"; return nil }

func (s *_createOperation) IfPrimaryTerm(ifprimaryterm int64) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) IfSeqNo(sequencenumber int64) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) Index_(indexname string) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) Pipeline(pipeline string) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) RequireAlias(requirealias bool) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) Routing(routing string) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) Version(versionnumber int64) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) VersionType(versiontype versiontype.VersionType) *_createOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) OperationContainerCaster() *types.OperationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_createOperation) CreateOperationCaster() *types.CreateOperation {
	_ = "STUB: not implemented"
	return nil
}
