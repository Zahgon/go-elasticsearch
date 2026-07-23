package types

type Float64 float64

func (f Float64) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Float64) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
