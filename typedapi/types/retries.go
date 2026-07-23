package types

type Retries struct {
	Bulk int64 `json:"bulk"`

	Search int64 `json:"search"`
}

func (s *Retries) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRetries() *Retries { _ = "STUB: not implemented"; return nil }
