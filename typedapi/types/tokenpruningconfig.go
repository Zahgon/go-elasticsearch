package types

type TokenPruningConfig struct {
	OnlyScorePrunedTokens *bool `json:"only_score_pruned_tokens,omitempty"`

	TokensFreqRatioThreshold *int `json:"tokens_freq_ratio_threshold,omitempty"`

	TokensWeightThreshold *float32 `json:"tokens_weight_threshold,omitempty"`
}

func (s *TokenPruningConfig) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTokenPruningConfig() *TokenPruningConfig { _ = "STUB: not implemented"; return nil }

type TokenPruningConfigVariant interface {
	TokenPruningConfigCaster() *TokenPruningConfig
}

func (s *TokenPruningConfig) TokenPruningConfigCaster() *TokenPruningConfig {
	_ = "STUB: not implemented"
	return nil
}
