package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticencoder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticlanguage"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticnametype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticruletype"
)

type PhoneticTokenFilter struct {
	Encoder     phoneticencoder.PhoneticEncoder     `json:"encoder"`
	Languageset []phoneticlanguage.PhoneticLanguage `json:"languageset,omitempty"`
	MaxCodeLen  *int                                `json:"max_code_len,omitempty"`
	NameType    *phoneticnametype.PhoneticNameType  `json:"name_type,omitempty"`
	Replace     *bool                               `json:"replace,omitempty"`
	RuleType    *phoneticruletype.PhoneticRuleType  `json:"rule_type,omitempty"`
	Type        string                              `json:"type,omitempty"`
	Version     *string                             `json:"version,omitempty"`
}

func (s *PhoneticTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PhoneticTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPhoneticTokenFilter() *PhoneticTokenFilter { _ = "STUB: not implemented"; return nil }

type PhoneticTokenFilterVariant interface {
	PhoneticTokenFilterCaster() *PhoneticTokenFilter
}

func (s *PhoneticTokenFilter) PhoneticTokenFilterCaster() *PhoneticTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PhoneticTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
