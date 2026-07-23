package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheretruncatetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/nvidiainputtype"
)

type NvidiaTaskSettings struct {
	InputType *nvidiainputtype.NvidiaInputType `json:"input_type,omitempty"`

	Truncate *coheretruncatetype.CohereTruncateType `json:"truncate,omitempty"`
}

func NewNvidiaTaskSettings() *NvidiaTaskSettings { _ = "STUB: not implemented"; return nil }

type NvidiaTaskSettingsVariant interface {
	NvidiaTaskSettingsCaster() *NvidiaTaskSettings
}

func (s *NvidiaTaskSettings) NvidiaTaskSettingsCaster() *NvidiaTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
