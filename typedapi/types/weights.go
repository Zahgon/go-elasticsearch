package types

type Weights struct {
	Weights Float64 `json:"weights"`
}

func (s *Weights) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWeights() *Weights { _ = "STUB: not implemented"; return nil }

type WeightsVariant interface {
	WeightsCaster() *Weights
}

func (s *Weights) WeightsCaster() *Weights { _ = "STUB: not implemented"; return nil }
