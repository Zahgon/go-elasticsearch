package types

type CharFilterTypes struct {
	AnalyzerTypes []FieldTypes `json:"analyzer_types"`

	BuiltInAnalyzers []FieldTypes `json:"built_in_analyzers"`

	BuiltInCharFilters []FieldTypes `json:"built_in_char_filters"`

	BuiltInFilters []FieldTypes `json:"built_in_filters"`

	BuiltInTokenizers []FieldTypes `json:"built_in_tokenizers"`

	CharFilterTypes []FieldTypes `json:"char_filter_types"`

	FilterTypes                 []FieldTypes                `json:"filter_types"`
	MultipleSynonymGraphFilters *MultipleSynonymGraphFilter `json:"multiple_synonym_graph_filters,omitempty"`

	Synonyms map[string]SynonymsStats `json:"synonyms"`

	TokenizerTypes []FieldTypes `json:"tokenizer_types"`
}

func NewCharFilterTypes() *CharFilterTypes { _ = "STUB: not implemented"; return nil }
