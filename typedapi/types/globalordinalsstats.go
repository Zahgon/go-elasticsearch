package types

type GlobalOrdinalsStats struct {
	BuildTime         *string                            `json:"build_time,omitempty"`
	BuildTimeInMillis int64                              `json:"build_time_in_millis"`
	Fields            map[string]GlobalOrdinalFieldStats `json:"fields,omitempty"`
}

func (s *GlobalOrdinalsStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGlobalOrdinalsStats() *GlobalOrdinalsStats { _ = "STUB: not implemented"; return nil }
