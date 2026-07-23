package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _charFilterDefinition struct {
	v types.CharFilterDefinition
}

func NewCharFilterDefinition() *_charFilterDefinition { _ = "STUB: not implemented"; return nil }

func (u *_charFilterDefinition) UnknownCharFilterDefinition(unknown json.RawMessage) *_charFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_charFilterDefinition) HtmlStripCharFilter(htmlstripcharfilter types.HtmlStripCharFilterVariant) *_charFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_htmlStripCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_charFilterDefinition) MappingCharFilter(mappingcharfilter types.MappingCharFilterVariant) *_charFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_mappingCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_charFilterDefinition) PatternReplaceCharFilter(patternreplacecharfilter types.PatternReplaceCharFilterVariant) *_charFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_patternReplaceCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_charFilterDefinition) IcuNormalizationCharFilter(icunormalizationcharfilter types.IcuNormalizationCharFilterVariant) *_charFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_icuNormalizationCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_charFilterDefinition) KuromojiIterationMarkCharFilter(kuromojiiterationmarkcharfilter types.KuromojiIterationMarkCharFilterVariant) *_charFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_kuromojiIterationMarkCharFilter) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (u *_charFilterDefinition) CharFilterDefinitionCaster() *types.CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
