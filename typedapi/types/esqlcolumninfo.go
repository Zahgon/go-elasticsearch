package types

type EsqlColumnInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s *EsqlColumnInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEsqlColumnInfo() *EsqlColumnInfo { _ = "STUB: not implemented"; return nil }
