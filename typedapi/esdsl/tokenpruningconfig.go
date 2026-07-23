package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _tokenPruningConfig struct {
	v *types.TokenPruningConfig
}

func NewTokenPruningConfig() *_tokenPruningConfig { _ = "STUB: not implemented"; return nil }

func (s *_tokenPruningConfig) OnlyScorePrunedTokens(onlyscoreprunedtokens bool) *_tokenPruningConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenPruningConfig) TokensFreqRatioThreshold(tokensfreqratiothreshold int) *_tokenPruningConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenPruningConfig) TokensWeightThreshold(tokensweightthreshold float32) *_tokenPruningConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenPruningConfig) SparseVectorQueryCaster() *types.SparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenPruningConfig) TokenPruningConfigCaster() *types.TokenPruningConfig {
	_ = "STUB: not implemented"
	return nil
}
