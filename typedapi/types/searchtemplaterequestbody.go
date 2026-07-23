package types

import (
	"encoding/json"
)

type SearchTemplateRequestBody struct {
	Explain *bool `json:"explain,omitempty"`

	Id      *string                    `json:"id,omitempty"`
	Params  map[string]json.RawMessage `json:"params,omitempty"`
	Profile *bool                      `json:"profile,omitempty"`

	Source *string `json:"source,omitempty"`
}

func (s *SearchTemplateRequestBody) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSearchTemplateRequestBody() *SearchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}

type SearchTemplateRequestBodyVariant interface {
	SearchTemplateRequestBodyCaster() *SearchTemplateRequestBody
}

func (s *SearchTemplateRequestBody) SearchTemplateRequestBodyCaster() *SearchTemplateRequestBody {
	_ = "STUB: not implemented"
	return nil
}
