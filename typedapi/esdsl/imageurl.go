package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/imageurldetail"
)

type _imageUrl struct {
	v *types.ImageUrl
}

func NewImageUrl(url string) *_imageUrl { _ = "STUB: not implemented"; return nil }

func (s *_imageUrl) Detail(detail imageurldetail.ImageUrlDetail) *_imageUrl {
	_ = "STUB: not implemented"
	return nil
}

func (s *_imageUrl) Url(url string) *_imageUrl { _ = "STUB: not implemented"; return nil }

func (s *_imageUrl) ImageUrlCaster() *types.ImageUrl { _ = "STUB: not implemented"; return nil }
