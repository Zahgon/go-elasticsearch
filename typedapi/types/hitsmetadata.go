package types

type HitsMetadata struct {
	Hits     []Hit    `json:"hits"`
	MaxScore *Float64 `json:"max_score,omitempty"`

	Total *TotalHits `json:"total,omitempty"`
}

func (s *HitsMetadata) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHitsMetadata() *HitsMetadata { _ = "STUB: not implemented"; return nil }
