package types

type KeyedPercentiles map[string]string

func (s KeyedPercentiles) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
