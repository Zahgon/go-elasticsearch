package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _indexOperation struct {
	v *types.IndexOperation
}

func NewIndexOperation() *_indexOperation { _ = "STUB: not implemented"; return nil }

func (s *_indexOperation) DynamicTemplates(dynamictemplates map[string]string) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) AddDynamicTemplate(key string, value string) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) Id_(id string) *_indexOperation { _ = "STUB: not implemented"; return nil }

func (s *_indexOperation) IfPrimaryTerm(ifprimaryterm int64) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) IfSeqNo(sequencenumber int64) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) Index_(indexname string) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) Pipeline(pipeline string) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) RequireAlias(requirealias bool) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) Routing(routing string) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) Version(versionnumber int64) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) VersionType(versiontype versiontype.VersionType) *_indexOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) OperationContainerCaster() *types.OperationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexOperation) IndexOperationCaster() *types.IndexOperation {
	_ = "STUB: not implemented"
	return nil
}
