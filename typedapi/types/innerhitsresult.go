package types

type InnerHitsResult struct {
	Hits HitsMetadata `json:"hits"`
}

func NewInnerHitsResult() *InnerHitsResult { _ = "STUB: not implemented"; return nil }
