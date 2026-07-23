package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _wktGeoBounds struct {
	v *types.WktGeoBounds
}

func NewWktGeoBounds(wkt string) *_wktGeoBounds { _ = "STUB: not implemented"; return nil }

func (s *_wktGeoBounds) Wkt(wkt string) *_wktGeoBounds { _ = "STUB: not implemented"; return nil }

func (s *_wktGeoBounds) WktGeoBoundsCaster() *types.WktGeoBounds {
	_ = "STUB: not implemented"
	return nil
}
