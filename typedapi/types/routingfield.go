package types

type RoutingField struct {
	Required bool `json:"required"`
}

func (s *RoutingField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRoutingField() *RoutingField { _ = "STUB: not implemented"; return nil }

type RoutingFieldVariant interface {
	RoutingFieldCaster() *RoutingField
}

func (s *RoutingField) RoutingFieldCaster() *RoutingField { _ = "STUB: not implemented"; return nil }
