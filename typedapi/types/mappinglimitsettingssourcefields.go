package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sourcemode"
)

type MappingLimitSettingsSourceFields struct {
	Mode sourcemode.SourceMode `json:"mode"`
}

func NewMappingLimitSettingsSourceFields() *MappingLimitSettingsSourceFields {
	_ = "STUB: not implemented"
	return nil
}

type MappingLimitSettingsSourceFieldsVariant interface {
	MappingLimitSettingsSourceFieldsCaster() *MappingLimitSettingsSourceFields
}

func (s *MappingLimitSettingsSourceFields) MappingLimitSettingsSourceFieldsCaster() *MappingLimitSettingsSourceFields {
	_ = "STUB: not implemented"
	return nil
}
