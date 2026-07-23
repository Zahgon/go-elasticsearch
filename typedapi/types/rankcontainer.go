package types

type RankContainer struct {
	Rrf *RrfRank `json:"rrf,omitempty"`
}

func NewRankContainer() *RankContainer { _ = "STUB: not implemented"; return nil }

type RankContainerVariant interface {
	RankContainerCaster() *RankContainer
}

func (s *RankContainer) RankContainerCaster() *RankContainer { _ = "STUB: not implemented"; return nil }
