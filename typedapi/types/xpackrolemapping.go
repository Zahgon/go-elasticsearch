package types

type XpackRoleMapping struct {
	Enabled int `json:"enabled"`
	Size    int `json:"size"`
}

func (s *XpackRoleMapping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewXpackRoleMapping() *XpackRoleMapping { _ = "STUB: not implemented"; return nil }
