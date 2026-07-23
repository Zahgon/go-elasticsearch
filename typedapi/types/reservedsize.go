package types

type ReservedSize struct {
	NodeId string   `json:"node_id"`
	Path   string   `json:"path"`
	Shards []string `json:"shards"`
	Total  int64    `json:"total"`
}

func (s *ReservedSize) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReservedSize() *ReservedSize { _ = "STUB: not implemented"; return nil }
