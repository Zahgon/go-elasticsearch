package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/imageurldetail"
)

type ImageUrl struct {
	Detail *imageurldetail.ImageUrlDetail `json:"detail,omitempty"`

	Url string `json:"url"`
}

func (s *ImageUrl) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewImageUrl() *ImageUrl { _ = "STUB: not implemented"; return nil }

type ImageUrlVariant interface {
	ImageUrlCaster() *ImageUrl
}

func (s *ImageUrl) ImageUrlCaster() *ImageUrl { _ = "STUB: not implemented"; return nil }
