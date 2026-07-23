package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/holtwinterstype"
)

type _holtWintersModelSettings struct {
	v *types.HoltWintersModelSettings
}

func NewHoltWintersModelSettings() *_holtWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersModelSettings) Alpha(alpha float32) *_holtWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersModelSettings) Beta(beta float32) *_holtWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersModelSettings) Gamma(gamma float32) *_holtWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersModelSettings) Pad(pad bool) *_holtWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersModelSettings) Period(period int) *_holtWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersModelSettings) Type(type_ holtwinterstype.HoltWintersType) *_holtWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersModelSettings) HoltWintersModelSettingsCaster() *types.HoltWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}
