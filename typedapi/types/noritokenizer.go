package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noridecompoundmode"
)

type NoriTokenizer struct {
	DecompoundMode      *noridecompoundmode.NoriDecompoundMode `json:"decompound_mode,omitempty"`
	DiscardPunctuation  *bool                                  `json:"discard_punctuation,omitempty"`
	Type                string                                 `json:"type,omitempty"`
	UserDictionary      *string                                `json:"user_dictionary,omitempty"`
	UserDictionaryRules []string                               `json:"user_dictionary_rules,omitempty"`
	Version             *string                                `json:"version,omitempty"`
}

func (s *NoriTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NoriTokenizer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNoriTokenizer() *NoriTokenizer { _ = "STUB: not implemented"; return nil }

type NoriTokenizerVariant interface {
	NoriTokenizerCaster() *NoriTokenizer
}

func (s *NoriTokenizer) NoriTokenizerCaster() *NoriTokenizer { _ = "STUB: not implemented"; return nil }

func (s *NoriTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
