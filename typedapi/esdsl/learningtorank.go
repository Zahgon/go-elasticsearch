package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _learningToRank struct {
	v *types.LearningToRank
}

func NewLearningToRank(modelid string) *_learningToRank { _ = "STUB: not implemented"; return nil }

func (s *_learningToRank) ModelId(modelid string) *_learningToRank {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRank) Params(params map[string]json.RawMessage) *_learningToRank {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRank) AddParam(key string, value json.RawMessage) *_learningToRank {
	_ = "STUB: not implemented"
	return nil
}

func (s *_learningToRank) RescoreCaster() *types.Rescore { _ = "STUB: not implemented"; return nil }

func (s *_learningToRank) LearningToRankCaster() *types.LearningToRank {
	_ = "STUB: not implemented"
	return nil
}
