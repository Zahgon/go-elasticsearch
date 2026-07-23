package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheretruncatetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/nvidiainputtype"
)

type _nvidiaTaskSettings struct {
	v *types.NvidiaTaskSettings
}

func NewNvidiaTaskSettings() *_nvidiaTaskSettings { _ = "STUB: not implemented"; return nil }

func (s *_nvidiaTaskSettings) InputType(inputtype nvidiainputtype.NvidiaInputType) *_nvidiaTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaTaskSettings) Truncate(truncate coheretruncatetype.CohereTruncateType) *_nvidiaTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nvidiaTaskSettings) NvidiaTaskSettingsCaster() *types.NvidiaTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
