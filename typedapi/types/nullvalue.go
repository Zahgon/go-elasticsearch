package types

type NullValue struct{}

func (n NullValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
