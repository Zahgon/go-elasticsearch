package types

type XpackQuery struct {
	Count  *int `json:"count,omitempty"`
	Failed *int `json:"failed,omitempty"`
	Paging *int `json:"paging,omitempty"`
	Total  *int `json:"total,omitempty"`
}

func (s *XpackQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewXpackQuery() *XpackQuery { _ = "STUB: not implemented"; return nil }
