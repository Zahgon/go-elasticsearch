package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sourcemode"
)

type _mappingLimitSettingsSourceFields struct {
	v *types.MappingLimitSettingsSourceFields
}

func NewMappingLimitSettingsSourceFields(mode sourcemode.SourceMode) *_mappingLimitSettingsSourceFields {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettingsSourceFields) Mode(mode sourcemode.SourceMode) *_mappingLimitSettingsSourceFields {
	_ = "STUB: not implemented"
	return nil
}

func (s *_mappingLimitSettingsSourceFields) MappingLimitSettingsSourceFieldsCaster() *types.MappingLimitSettingsSourceFields {
	_ = "STUB: not implemented"
	return nil
}
