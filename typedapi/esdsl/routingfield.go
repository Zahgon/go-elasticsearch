package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _routingField struct {
	v *types.RoutingField
}

func NewRoutingField(required bool) *_routingField { _ = "STUB: not implemented"; return nil }

func (s *_routingField) Required(required bool) *_routingField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_routingField) RoutingFieldCaster() *types.RoutingField {
	_ = "STUB: not implemented"
	return nil
}
