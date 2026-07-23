package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _httpEmailAttachment struct {
	v *types.HttpEmailAttachment
}

func NewHttpEmailAttachment() *_httpEmailAttachment { _ = "STUB: not implemented"; return nil }

func (s *_httpEmailAttachment) ContentType(contenttype string) *_httpEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpEmailAttachment) Inline(inline bool) *_httpEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpEmailAttachment) Request(request types.HttpInputRequestDefinitionVariant) *_httpEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpEmailAttachment) EmailAttachmentContainerCaster() *types.EmailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_httpEmailAttachment) HttpEmailAttachmentCaster() *types.HttpEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}
