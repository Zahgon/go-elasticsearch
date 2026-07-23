package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/contenttype"
)

type ContentObject struct {
	File FileContent `json:"file"`

	ImageUrl ImageUrl `json:"image_url"`

	Text string `json:"text"`

	Type contenttype.ContentType `json:"type"`
}

func (s *ContentObject) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewContentObject() *ContentObject { _ = "STUB: not implemented"; return nil }

type ContentObjectVariant interface {
	ContentObjectCaster() *ContentObject
}

func (s *ContentObject) ContentObjectCaster() *ContentObject { _ = "STUB: not implemented"; return nil }
