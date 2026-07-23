package types

type EqlHits struct {
	Events []HitsEvent `json:"events,omitempty"`

	Sequences []HitsSequence `json:"sequences,omitempty"`

	Total *TotalHits `json:"total,omitempty"`
}

func NewEqlHits() *EqlHits { _ = "STUB: not implemented"; return nil }
