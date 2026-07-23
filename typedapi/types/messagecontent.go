package types

type MessageContent any

type MessageContentVariant interface {
	MessageContentCaster() *MessageContent
}
