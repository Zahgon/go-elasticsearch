package types

type InvertedIndex struct {
	Offsets         uint `json:"offsets"`
	Payloads        uint `json:"payloads"`
	Positions       uint `json:"positions"`
	Postings        uint `json:"postings"`
	Proximity       uint `json:"proximity"`
	TermFrequencies uint `json:"term_frequencies"`
	Terms           uint `json:"terms"`
}

func NewInvertedIndex() *InvertedIndex { _ = "STUB: not implemented"; return nil }
