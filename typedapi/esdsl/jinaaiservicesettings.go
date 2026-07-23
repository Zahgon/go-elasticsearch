package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaielementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaisimilaritytype"
)

type _jinaAIServiceSettings struct {
	v *types.JinaAIServiceSettings
}

func NewJinaAIServiceSettings(apikey string, modelid string) *_jinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAIServiceSettings) ApiKey(apikey string) *_jinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAIServiceSettings) Dimensions(dimensions int) *_jinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAIServiceSettings) EmbeddingType(embeddingtype jinaaielementtype.JinaAIElementType) *_jinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAIServiceSettings) ModelId(modelid string) *_jinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAIServiceSettings) MultimodalModel(multimodalmodel bool) *_jinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_jinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAIServiceSettings) Similarity(similarity jinaaisimilaritytype.JinaAISimilarityType) *_jinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jinaAIServiceSettings) JinaAIServiceSettingsCaster() *types.JinaAIServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
