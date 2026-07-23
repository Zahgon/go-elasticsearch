package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectorfieldtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/displaytype"
)

type ConnectorConfigProperties struct {
	Category       *string                                `json:"category,omitempty"`
	DefaultValue   ScalarValue                            `json:"default_value"`
	DependsOn      []Dependency                           `json:"depends_on"`
	Display        displaytype.DisplayType                `json:"display"`
	Label          string                                 `json:"label"`
	Options        []SelectOption                         `json:"options"`
	Order          *int                                   `json:"order,omitempty"`
	Placeholder    *string                                `json:"placeholder,omitempty"`
	Required       bool                                   `json:"required"`
	Sensitive      bool                                   `json:"sensitive"`
	Tooltip        *string                                `json:"tooltip,omitempty"`
	Type           *connectorfieldtype.ConnectorFieldType `json:"type,omitempty"`
	UiRestrictions []string                               `json:"ui_restrictions,omitempty"`
	Validations    []Validation                           `json:"validations,omitempty"`
	Value          json.RawMessage                        `json:"value,omitempty"`
}

func (s *ConnectorConfigProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewConnectorConfigProperties() *ConnectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

type ConnectorConfigPropertiesVariant interface {
	ConnectorConfigPropertiesCaster() *ConnectorConfigProperties
}

func (s *ConnectorConfigProperties) ConnectorConfigPropertiesCaster() *ConnectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}
