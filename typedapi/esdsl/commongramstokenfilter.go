package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _commonGramsTokenFilter struct {
	v *types.CommonGramsTokenFilter
}

func NewCommonGramsTokenFilter() *_commonGramsTokenFilter { _ = "STUB: not implemented"; return nil }

func (s *_commonGramsTokenFilter) CommonWords(commonwords ...string) *_commonGramsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonGramsTokenFilter) CommonWordsPath(commonwordspath string) *_commonGramsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonGramsTokenFilter) IgnoreCase(ignorecase bool) *_commonGramsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonGramsTokenFilter) QueryMode(querymode bool) *_commonGramsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonGramsTokenFilter) Version(versionstring string) *_commonGramsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_commonGramsTokenFilter) CommonGramsTokenFilterCaster() *types.CommonGramsTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
