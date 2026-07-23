package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _requestItem struct {
	v types.RequestItem
}

func NewRequestItem() *_requestItem { _ = "STUB: not implemented"; return nil }

func (u *_requestItem) MultisearchHeader(multisearchheader types.MultisearchHeaderVariant) *_requestItem {
	_ = "STUB: not implemented"
	return nil
}

func (u *_multisearchHeader) RequestItemCaster() *types.RequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (u *_requestItem) TemplateConfig(templateconfig types.TemplateConfigVariant) *_requestItem {
	_ = "STUB: not implemented"
	return nil
}

func (u *_templateConfig) RequestItemCaster() *types.RequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (u *_requestItem) RequestItemCaster() *types.RequestItem {
	_ = "STUB: not implemented"
	return nil
}
