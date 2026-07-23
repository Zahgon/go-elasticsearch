package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _searchTemplateRequestBody struct {
	v *types.SearchTemplateRequestBody
}

func NewSearchTemplateRequestBody() *_searchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTemplateRequestBody) Explain(explain bool) *_searchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTemplateRequestBody) Id(id string) *_searchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTemplateRequestBody) Params(params map[string]json.RawMessage) *_searchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTemplateRequestBody) AddParam(key string, value json.RawMessage) *_searchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTemplateRequestBody) Profile(profile bool) *_searchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTemplateRequestBody) Source(source string) *_searchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchTemplateRequestBody) SearchTemplateRequestBodyCaster() *types.SearchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}
