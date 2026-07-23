package types

type Retention struct {
	ExpireAfter Duration `json:"expire_after"`

	MaxCount int `json:"max_count"`

	MinCount int `json:"min_count"`
}

func (s *Retention) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRetention() *Retention { _ = "STUB: not implemented"; return nil }

type RetentionVariant interface {
	RetentionCaster() *Retention
}

func (s *Retention) RetentionCaster() *Retention { _ = "STUB: not implemented"; return nil }
