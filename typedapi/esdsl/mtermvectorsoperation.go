package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _mTermVectorsOperation struct {
	v *types.MTermVectorsOperation
}

func NewMTermVectorsOperation() *_mTermVectorsOperation { _ = "STUB: not implemented"; return nil }

func (s *_mTermVectorsOperation) Doc(doc json.RawMessage) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) FieldStatistics(fieldstatistics bool) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Fields(fields ...string) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Filter(filter types.TermVectorsFilterVariant) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Id_(id string) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Index_(indexname string) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Offsets(offsets bool) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Payloads(payloads bool) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Positions(positions bool) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Routing(routings ...string) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) TermStatistics(termstatistics bool) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) Version(versionnumber int64) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) VersionType(versiontype versiontype.VersionType) *_mTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mTermVectorsOperation) MTermVectorsOperationCaster() *types.MTermVectorsOperation {
	_ = "STUB: not implemented"
	return nil
}
