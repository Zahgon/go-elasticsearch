package types

type TermVectorsToken struct {
	EndOffset   *int    `json:"end_offset,omitempty"`
	Payload     *string `json:"payload,omitempty"`
	Position    int     `json:"position"`
	StartOffset *int    `json:"start_offset,omitempty"`
}

func (s *TermVectorsToken) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermVectorsToken() *TermVectorsToken { _ = "STUB: not implemented"; return nil }
