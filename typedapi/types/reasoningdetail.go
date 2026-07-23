package types

type ReasoningDetail any

type ReasoningDetailVariant interface {
	ReasoningDetailCaster() *ReasoningDetail
}
