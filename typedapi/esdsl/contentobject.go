package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/contenttype"
)

type _contentObject struct {
	v *types.ContentObject
}

func NewContentObject(file types.FileContentVariant, imageurl types.ImageUrlVariant, text string, type_ contenttype.ContentType) *_contentObject {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contentObject) File(file types.FileContentVariant) *_contentObject {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contentObject) ImageUrl(imageurl types.ImageUrlVariant) *_contentObject {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contentObject) Text(text string) *_contentObject { _ = "STUB: not implemented"; return nil }

func (s *_contentObject) Type(type_ contenttype.ContentType) *_contentObject {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contentObject) ContentObjectCaster() *types.ContentObject {
	_ = "STUB: not implemented"
	return nil
}
