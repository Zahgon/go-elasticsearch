package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/llamasimilaritytype"
)

type _llamaServiceSettings struct {
	v *types.LlamaServiceSettings
}

func NewLlamaServiceSettings(modelid string, url string) *_llamaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_llamaServiceSettings) MaxInputTokens(maxinputtokens int) *_llamaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_llamaServiceSettings) ModelId(modelid string) *_llamaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_llamaServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_llamaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_llamaServiceSettings) Similarity(similarity llamasimilaritytype.LlamaSimilarityType) *_llamaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_llamaServiceSettings) Url(url string) *_llamaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_llamaServiceSettings) LlamaServiceSettingsCaster() *types.LlamaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
