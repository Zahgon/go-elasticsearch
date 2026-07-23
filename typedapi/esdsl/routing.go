package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _routing struct {
	v types.Routing
}

func NewRouting() *_routing { _ = "STUB: not implemented"; return nil }

func (u *_routing) Strings(strings ...string) *_routing { _ = "STUB: not implemented"; return nil }

func (u *_routing) RoutingCaster() *types.Routing { _ = "STUB: not implemented"; return nil }
