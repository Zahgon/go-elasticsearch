package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _messageContent struct {
	v types.MessageContent
}

func NewMessageContent() *_messageContent { _ = "STUB: not implemented"; return nil }

func (u *_messageContent) String(string string) *_messageContent {
	_ = "STUB: not implemented"
	return nil
}

func (u *_messageContent) ContentObjects(contentobjects ...types.ContentObjectVariant) *_messageContent {
	_ = "STUB: not implemented"
	return nil
}

func (u *_messageContent) MessageContentCaster() *types.MessageContent {
	_ = "STUB: not implemented"
	return nil
}
