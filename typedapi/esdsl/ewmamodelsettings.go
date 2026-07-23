package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ewmaModelSettings struct {
	v *types.EwmaModelSettings
}

func NewEwmaModelSettings() *_ewmaModelSettings { _ = "STUB: not implemented"; return nil }

func (s *_ewmaModelSettings) Alpha(alpha float32) *_ewmaModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaModelSettings) EwmaModelSettingsCaster() *types.EwmaModelSettings {
	_ = "STUB: not implemented"
	return nil
}
