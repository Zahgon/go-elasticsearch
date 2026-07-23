package types

type MsearchRequestItem any

type MsearchRequestItemVariant interface {
	MsearchRequestItemCaster() *MsearchRequestItem
}
