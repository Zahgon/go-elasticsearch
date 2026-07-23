package types

type CloseIndexResult struct {
	Closed bool                        `json:"closed"`
	Shards map[string]CloseShardResult `json:"shards,omitempty"`
}

func (s *CloseIndexResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCloseIndexResult() *CloseIndexResult { _ = "STUB: not implemented"; return nil }
