package types

type RequestItem any

type RequestItemVariant interface {
	RequestItemCaster() *RequestItem
}
