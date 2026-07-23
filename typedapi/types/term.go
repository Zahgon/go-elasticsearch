package types

type Term struct {
	DocFreq  *int               `json:"doc_freq,omitempty"`
	Score    *Float64           `json:"score,omitempty"`
	TermFreq int                `json:"term_freq"`
	Tokens   []TermVectorsToken `json:"tokens,omitempty"`
	Ttf      *int               `json:"ttf,omitempty"`
}

func (s *Term) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTerm() *Term { _ = "STUB: not implemented"; return nil }
