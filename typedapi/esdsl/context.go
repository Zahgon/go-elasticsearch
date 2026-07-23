package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _context struct {
	v types.Context
}

func NewContext() *_context { _ = "STUB: not implemented"; return nil }

func (u *_context) String(string string) *_context { _ = "STUB: not implemented"; return nil }

func (u *_context) GeoLocation(geolocation types.GeoLocationVariant) *_context {
	_ = "STUB: not implemented"
	return nil
}

func (u *_geoLocation) ContextCaster() *types.Context { _ = "STUB: not implemented"; return nil }

func (u *_context) ContextCaster() *types.Context { _ = "STUB: not implemented"; return nil }
