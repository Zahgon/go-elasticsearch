package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slackMessage struct {
	v *types.SlackMessage
}

func NewSlackMessage(from string, text string) *_slackMessage {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackMessage) Attachments(attachments ...types.SlackAttachmentVariant) *_slackMessage {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackMessage) AttachmentsValues(attachmentsvalues []types.SlackAttachment) *_slackMessage {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackMessage) DynamicAttachments(dynamicattachments types.SlackDynamicAttachmentVariant) *_slackMessage {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackMessage) From(from string) *_slackMessage { _ = "STUB: not implemented"; return nil }

func (s *_slackMessage) Icon(icon string) *_slackMessage { _ = "STUB: not implemented"; return nil }

func (s *_slackMessage) Text(text string) *_slackMessage { _ = "STUB: not implemented"; return nil }

func (s *_slackMessage) To(tos ...string) *_slackMessage { _ = "STUB: not implemented"; return nil }

func (s *_slackMessage) SlackMessageCaster() *types.SlackMessage {
	_ = "STUB: not implemented"
	return nil
}
