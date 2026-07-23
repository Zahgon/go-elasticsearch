package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textIndexPrefixes struct {
	v *types.TextIndexPrefixes
}

func NewTextIndexPrefixes() *_textIndexPrefixes { _ = "STUB: not implemented"; return nil }

func (s *_textIndexPrefixes) MaxChars(maxchars int) *_textIndexPrefixes {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textIndexPrefixes) MinChars(minchars int) *_textIndexPrefixes {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textIndexPrefixes) TextIndexPrefixesCaster() *types.TextIndexPrefixes {
	_ = "STUB: not implemented"
	return nil
}
