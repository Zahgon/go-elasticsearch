package types

type ByteSize any

type ByteSizeVariant interface {
	ByteSizeCaster() *ByteSize
}
