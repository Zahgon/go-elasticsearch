package types

type TranslogRetention struct {
	Age Duration `json:"age,omitempty"`

	Size ByteSize `json:"size,omitempty"`
}

func (s *TranslogRetention) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTranslogRetention() *TranslogRetention { _ = "STUB: not implemented"; return nil }

type TranslogRetentionVariant interface {
	TranslogRetentionCaster() *TranslogRetention
}

func (s *TranslogRetention) TranslogRetentionCaster() *TranslogRetention {
	_ = "STUB: not implemented"
	return nil
}
