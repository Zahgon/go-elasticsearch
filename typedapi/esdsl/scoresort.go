package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _scoreSort struct {
	v *types.ScoreSort
}

func NewScoreSort() *_scoreSort { _ = "STUB: not implemented"; return nil }

func (s *_scoreSort) Order(order sortorder.SortOrder) *_scoreSort {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scoreSort) SortOptionsCaster() *types.SortOptions { _ = "STUB: not implemented"; return nil }

func (s *_scoreSort) ScoreSortCaster() *types.ScoreSort { _ = "STUB: not implemented"; return nil }
