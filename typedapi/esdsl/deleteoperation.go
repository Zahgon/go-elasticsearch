package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _deleteOperation struct {
	v *types.DeleteOperation
}

func NewDeleteOperation() *_deleteOperation { _ = "STUB: not implemented"; return nil }

func (s *_deleteOperation) Id_(id string) *_deleteOperation { _ = "STUB: not implemented"; return nil }

func (s *_deleteOperation) IfPrimaryTerm(ifprimaryterm int64) *_deleteOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_deleteOperation) IfSeqNo(sequencenumber int64) *_deleteOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_deleteOperation) Index_(indexname string) *_deleteOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_deleteOperation) Routing(routing string) *_deleteOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_deleteOperation) Version(versionnumber int64) *_deleteOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_deleteOperation) VersionType(versiontype versiontype.VersionType) *_deleteOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_deleteOperation) OperationContainerCaster() *types.OperationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_deleteOperation) DeleteOperationCaster() *types.DeleteOperation {
	_ = "STUB: not implemented"
	return nil
}
