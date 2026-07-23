package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _allField struct {
	v *types.AllField
}

func NewAllField(analyzer string, enabled bool, omitnorms bool, searchanalyzer string, similarity string, store bool, storetermvectoroffsets bool, storetermvectorpayloads bool, storetermvectorpositions bool, storetermvectors bool) *_allField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allField) Analyzer(analyzer string) *_allField { _ = "STUB: not implemented"; return nil }

func (s *_allField) Enabled(enabled bool) *_allField { _ = "STUB: not implemented"; return nil }

func (s *_allField) OmitNorms(omitnorms bool) *_allField { _ = "STUB: not implemented"; return nil }

func (s *_allField) SearchAnalyzer(searchanalyzer string) *_allField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allField) Similarity(similarity string) *_allField { _ = "STUB: not implemented"; return nil }

func (s *_allField) Store(store bool) *_allField { _ = "STUB: not implemented"; return nil }

func (s *_allField) StoreTermVectorOffsets(storetermvectoroffsets bool) *_allField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allField) StoreTermVectorPayloads(storetermvectorpayloads bool) *_allField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allField) StoreTermVectorPositions(storetermvectorpositions bool) *_allField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allField) StoreTermVectors(storetermvectors bool) *_allField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_allField) AllFieldCaster() *types.AllField { _ = "STUB: not implemented"; return nil }
