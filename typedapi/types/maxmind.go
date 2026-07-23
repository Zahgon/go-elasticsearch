package types

type Maxmind struct {
	AccountId string `json:"account_id"`
}

func (s *Maxmind) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMaxmind() *Maxmind { _ = "STUB: not implemented"; return nil }

type MaxmindVariant interface {
	MaxmindCaster() *Maxmind
}

func (s *Maxmind) MaxmindCaster() *Maxmind { _ = "STUB: not implemented"; return nil }
