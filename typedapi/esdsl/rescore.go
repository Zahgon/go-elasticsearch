package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _rescore struct {
	v *types.Rescore
}

func NewRescore() *_rescore { _ = "STUB: not implemented"; return nil }

func (s *_rescore) AdditionalRescoreProperty(key string, value json.RawMessage) *_rescore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescore) LearningToRank(learningtorank types.LearningToRankVariant) *_rescore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescore) Query(query types.RescoreQueryVariant) *_rescore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescore) Script(script types.ScriptRescoreVariant) *_rescore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescore) WindowSize(windowsize int) *_rescore { _ = "STUB: not implemented"; return nil }

func (s *_rescore) RescoreCaster() *types.Rescore { _ = "STUB: not implemented"; return nil }
