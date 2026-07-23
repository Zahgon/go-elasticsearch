package types

type CartesianPoint struct {
	X Float64 `json:"x"`
	Y Float64 `json:"y"`
}

func (s *CartesianPoint) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCartesianPoint() *CartesianPoint { _ = "STUB: not implemented"; return nil }
