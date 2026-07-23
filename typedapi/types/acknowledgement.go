package types

type Acknowledgement struct {
	License []string `json:"license"`
	Message string   `json:"message"`
}

func (s *Acknowledgement) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAcknowledgement() *Acknowledgement { _ = "STUB: not implemented"; return nil }
