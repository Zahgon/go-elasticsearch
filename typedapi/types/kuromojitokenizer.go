package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/kuromojitokenizationmode"
)

type KuromojiTokenizer struct {
	DiscardCompoundToken *bool                                             `json:"discard_compound_token,omitempty"`
	DiscardPunctuation   *bool                                             `json:"discard_punctuation,omitempty"`
	Mode                 kuromojitokenizationmode.KuromojiTokenizationMode `json:"mode"`
	NbestCost            *int                                              `json:"nbest_cost,omitempty"`
	NbestExamples        *string                                           `json:"nbest_examples,omitempty"`
	Type                 string                                            `json:"type,omitempty"`
	UserDictionary       *string                                           `json:"user_dictionary,omitempty"`
	UserDictionaryRules  []string                                          `json:"user_dictionary_rules,omitempty"`
	Version              *string                                           `json:"version,omitempty"`
}

func (s *KuromojiTokenizer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s KuromojiTokenizer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKuromojiTokenizer() *KuromojiTokenizer { _ = "STUB: not implemented"; return nil }

type KuromojiTokenizerVariant interface {
	KuromojiTokenizerCaster() *KuromojiTokenizer
}

func (s *KuromojiTokenizer) KuromojiTokenizerCaster() *KuromojiTokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (s *KuromojiTokenizer) TokenizerDefinitionCaster() *TokenizerDefinition {
	_ = "STUB: not implemented"
	return nil
}
