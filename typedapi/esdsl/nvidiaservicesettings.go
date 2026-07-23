package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/nvidiasimilaritytype"
)

type _nvidiaServiceSettings struct {
	v *types.NvidiaServiceSettings
}

func NewNvidiaServiceSettings(apikey string, modelid string) *_nvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaServiceSettings) ApiKey(apikey string) *_nvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaServiceSettings) MaxInputTokens(maxinputtokens int) *_nvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaServiceSettings) ModelId(modelid string) *_nvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_nvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaServiceSettings) Similarity(similarity nvidiasimilaritytype.NvidiaSimilarityType) *_nvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaServiceSettings) Url(url string) *_nvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaServiceSettings) NvidiaServiceSettingsCaster() *types.NvidiaServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
