package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/responsecontenttype"
)

type _httpInput struct {
	v *types.HttpInput
}

func NewHttpInput() *_httpInput { _ = "STUB: not implemented"; return nil }

func (s *_httpInput) Extract(extracts ...string) *_httpInput { _ = "STUB: not implemented"; return nil }

func (s *_httpInput) Request(request types.HttpInputRequestDefinitionVariant) *_httpInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInput) ResponseContentType(responsecontenttype responsecontenttype.ResponseContentType) *_httpInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInput) WatcherInputCaster() *types.WatcherInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpInput) HttpInputCaster() *types.HttpInput { _ = "STUB: not implemented"; return nil }
