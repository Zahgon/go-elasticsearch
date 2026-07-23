package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _msearchRequestItem struct {
	v types.MsearchRequestItem
}

func NewMsearchRequestItem() *_msearchRequestItem { _ = "STUB: not implemented"; return nil }

func (u *_msearchRequestItem) MultisearchHeader(multisearchheader types.MultisearchHeaderVariant) *_msearchRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (u *_multisearchHeader) MsearchRequestItemCaster() *types.MsearchRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (u *_msearchRequestItem) SearchRequestBody(searchrequestbody types.SearchRequestBodyVariant) *_msearchRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (u *_searchRequestBody) MsearchRequestItemCaster() *types.MsearchRequestItem {
	_ = "STUB: not implemented"
	return nil
}

func (u *_msearchRequestItem) MsearchRequestItemCaster() *types.MsearchRequestItem {
	_ = "STUB: not implemented"
	return nil
}
