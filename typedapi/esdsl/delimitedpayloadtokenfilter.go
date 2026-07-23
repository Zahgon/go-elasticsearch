package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/delimitedpayloadencoding"
)

type _delimitedPayloadTokenFilter struct {
	v *types.DelimitedPayloadTokenFilter
}

func NewDelimitedPayloadTokenFilter() *_delimitedPayloadTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_delimitedPayloadTokenFilter) Delimiter(delimiter string) *_delimitedPayloadTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_delimitedPayloadTokenFilter) Encoding(encoding delimitedpayloadencoding.DelimitedPayloadEncoding) *_delimitedPayloadTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_delimitedPayloadTokenFilter) Version(versionstring string) *_delimitedPayloadTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_delimitedPayloadTokenFilter) DelimitedPayloadTokenFilterCaster() *types.DelimitedPayloadTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
