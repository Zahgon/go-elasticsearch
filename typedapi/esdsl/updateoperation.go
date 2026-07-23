package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _updateOperation struct {
	v *types.UpdateOperation
}

func NewUpdateOperation() *_updateOperation { _ = "STUB: not implemented"; return nil }

func (s *_updateOperation) RequireAlias(requirealias bool) *_updateOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) RetryOnConflict(retryonconflict int) *_updateOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) Id_(id string) *_updateOperation { _ = "STUB: not implemented"; return nil }

func (s *_updateOperation) IfPrimaryTerm(ifprimaryterm int64) *_updateOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) IfSeqNo(sequencenumber int64) *_updateOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) Index_(indexname string) *_updateOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) Routing(routing string) *_updateOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) Version(versionnumber int64) *_updateOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) VersionType(versiontype versiontype.VersionType) *_updateOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) OperationContainerCaster() *types.OperationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateOperation) UpdateOperationCaster() *types.UpdateOperation {
	_ = "STUB: not implemented"
	return nil
}
