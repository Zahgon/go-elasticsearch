package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _innerHits struct {
	v *types.InnerHits
}

func NewInnerHits() *_innerHits { _ = "STUB: not implemented"; return nil }

func (s *_innerHits) Collapse(collapse types.FieldCollapseVariant) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) DocvalueFieldsValues(docvaluefieldsvalues []types.FieldAndFormat) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) Explain(explain bool) *_innerHits { _ = "STUB: not implemented"; return nil }

func (s *_innerHits) Field(fields ...string) *_innerHits { _ = "STUB: not implemented"; return nil }

func (s *_innerHits) Fields(fields ...types.FieldAndFormatVariant) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) FieldsValues(fieldsvalues []types.FieldAndFormat) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) From(from int) *_innerHits { _ = "STUB: not implemented"; return nil }

func (s *_innerHits) Highlight(highlight types.HighlightVariant) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) IgnoreUnmapped(ignoreunmapped bool) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) Name(name string) *_innerHits { _ = "STUB: not implemented"; return nil }

func (s *_innerHits) ScriptFields(scriptfields map[string]types.ScriptField) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) AddScriptField(key string, value types.ScriptFieldVariant) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) SeqNoPrimaryTerm(seqnoprimaryterm bool) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) Size(size int) *_innerHits { _ = "STUB: not implemented"; return nil }

func (s *_innerHits) Sort(sorts ...types.SortCombinationsVariant) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) SortValues(sortvalues []types.SortCombinations) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) Source_(sourceconfig types.SourceConfigVariant) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) StoredFields(fields ...string) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) TrackScores(trackscores bool) *_innerHits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerHits) Version(version bool) *_innerHits { _ = "STUB: not implemented"; return nil }

func (s *_innerHits) InnerHitsCaster() *types.InnerHits { _ = "STUB: not implemented"; return nil }
