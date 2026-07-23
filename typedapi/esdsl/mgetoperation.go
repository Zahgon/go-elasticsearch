package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _mgetOperation struct {
	v *types.MgetOperation
}

func NewMgetOperation() *_mgetOperation { _ = "STUB: not implemented"; return nil }

func (s *_mgetOperation) Id_(id string) *_mgetOperation { _ = "STUB: not implemented"; return nil }

func (s *_mgetOperation) Index_(indexname string) *_mgetOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mgetOperation) Routing(routings ...string) *_mgetOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mgetOperation) Source_(sourceconfig types.SourceConfigVariant) *_mgetOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mgetOperation) StoredFields(fields ...string) *_mgetOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mgetOperation) Version(versionnumber int64) *_mgetOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mgetOperation) VersionType(versiontype versiontype.VersionType) *_mgetOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mgetOperation) MgetOperationCaster() *types.MgetOperation {
	_ = "STUB: not implemented"
	return nil
}
