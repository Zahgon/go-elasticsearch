package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/responsecontenttype"
)

type HttpInput struct {
	Extract             []string                                 `json:"extract,omitempty"`
	Request             *HttpInputRequestDefinition              `json:"request,omitempty"`
	ResponseContentType *responsecontenttype.ResponseContentType `json:"response_content_type,omitempty"`
}

func NewHttpInput() *HttpInput { _ = "STUB: not implemented"; return nil }

type HttpInputVariant interface {
	HttpInputCaster() *HttpInput
}

func (s *HttpInput) HttpInputCaster() *HttpInput { _ = "STUB: not implemented"; return nil }
