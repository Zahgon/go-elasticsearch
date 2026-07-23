package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/connectorfieldtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/displaytype"
)

type _connectorConfigProperties struct {
	v *types.ConnectorConfigProperties
}

func NewConnectorConfigProperties(display displaytype.DisplayType, label string, required bool, sensitive bool, value json.RawMessage) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Category(category string) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) DefaultValue(scalarvalue types.ScalarValueVariant) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) DependsOn(dependsons ...types.DependencyVariant) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) DependsOnValues(dependsonvalues []types.Dependency) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Display(display displaytype.DisplayType) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Label(label string) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Options(options ...types.SelectOptionVariant) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) OptionsValues(optionsvalues []types.SelectOption) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Order(order int) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Placeholder(placeholder string) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Required(required bool) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Sensitive(sensitive bool) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Tooltip(tooltip string) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Type(type_ connectorfieldtype.ConnectorFieldType) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) UiRestrictions(uirestrictions ...string) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Validations(validations ...types.ValidationVariant) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) ValidationsValues(validationsvalues []types.Validation) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) Value(value json.RawMessage) *_connectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *_connectorConfigProperties) ConnectorConfigPropertiesCaster() *types.ConnectorConfigProperties {
	_ = "STUB: not implemented"
	return nil
}
