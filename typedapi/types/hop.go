package types

type Hop struct {
	Connections *Hop `json:"connections,omitempty"`

	Query *Query `json:"query,omitempty"`

	Vertices []VertexDefinition `json:"vertices"`
}

func NewHop() *Hop { _ = "STUB: not implemented"; return nil }

type HopVariant interface {
	HopCaster() *Hop
}

func (s *Hop) HopCaster() *Hop { _ = "STUB: not implemented"; return nil }
