package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/holtwinterstype"
)

type HoltWintersModelSettings struct {
	Alpha  *float32                         `json:"alpha,omitempty"`
	Beta   *float32                         `json:"beta,omitempty"`
	Gamma  *float32                         `json:"gamma,omitempty"`
	Pad    *bool                            `json:"pad,omitempty"`
	Period *int                             `json:"period,omitempty"`
	Type   *holtwinterstype.HoltWintersType `json:"type,omitempty"`
}

func (s *HoltWintersModelSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHoltWintersModelSettings() *HoltWintersModelSettings { _ = "STUB: not implemented"; return nil }

type HoltWintersModelSettingsVariant interface {
	HoltWintersModelSettingsCaster() *HoltWintersModelSettings
}

func (s *HoltWintersModelSettings) HoltWintersModelSettingsCaster() *HoltWintersModelSettings {
	_ = "STUB: not implemented"
	return nil
}
