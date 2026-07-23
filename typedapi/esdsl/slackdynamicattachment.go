package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slackDynamicAttachment struct {
	v *types.SlackDynamicAttachment
}

func NewSlackDynamicAttachment(attachmenttemplate types.SlackAttachmentVariant, listpath string) *_slackDynamicAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackDynamicAttachment) AttachmentTemplate(attachmenttemplate types.SlackAttachmentVariant) *_slackDynamicAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackDynamicAttachment) ListPath(listpath string) *_slackDynamicAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackDynamicAttachment) SlackDynamicAttachmentCaster() *types.SlackDynamicAttachment {
	_ = "STUB: not implemented"
	return nil
}
