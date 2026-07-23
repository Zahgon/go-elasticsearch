package types

type Routing []string

type RoutingVariant interface {
	RoutingCaster() *Routing
}
