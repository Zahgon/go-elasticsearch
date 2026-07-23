package types

type Connection struct {
	DocCount int64   `json:"doc_count"`
	Source   int64   `json:"source"`
	Target   int64   `json:"target"`
	Weight   Float64 `json:"weight"`
}

func (s *Connection) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewConnection() *Connection { _ = "STUB: not implemented"; return nil }
