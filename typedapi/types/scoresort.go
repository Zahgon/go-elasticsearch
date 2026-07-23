package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type ScoreSort struct {
	Order *sortorder.SortOrder `json:"order,omitempty"`
}

func NewScoreSort() *ScoreSort { _ = "STUB: not implemented"; return nil }

type ScoreSortVariant interface {
	ScoreSortCaster() *ScoreSort
}

func (s *ScoreSort) ScoreSortCaster() *ScoreSort { _ = "STUB: not implemented"; return nil }
