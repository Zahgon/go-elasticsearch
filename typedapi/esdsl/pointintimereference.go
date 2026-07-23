package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pointInTimeReference struct {
	v *types.PointInTimeReference
}

func NewPointInTimeReference() *_pointInTimeReference { _ = "STUB: not implemented"; return nil }

func (s *_pointInTimeReference) Id(id string) *_pointInTimeReference {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointInTimeReference) KeepAlive(duration types.DurationVariant) *_pointInTimeReference {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pointInTimeReference) PointInTimeReferenceCaster() *types.PointInTimeReference {
	_ = "STUB: not implemented"
	return nil
}
