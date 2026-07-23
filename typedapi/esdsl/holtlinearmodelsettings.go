package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _holtLinearModelSettings struct {
	v *types.HoltLinearModelSettings
}

func NewHoltLinearModelSettings() *_holtLinearModelSettings { _ = "STUB: not implemented"; return nil }

func (s *_holtLinearModelSettings) Alpha(alpha float32) *_holtLinearModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtLinearModelSettings) Beta(beta float32) *_holtLinearModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtLinearModelSettings) HoltLinearModelSettingsCaster() *types.HoltLinearModelSettings {
	_ = "STUB: not implemented"
	return nil
}
