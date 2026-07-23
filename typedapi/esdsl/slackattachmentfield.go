package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slackAttachmentField struct {
	v *types.SlackAttachmentField
}

func NewSlackAttachmentField(int bool, title string, value string) *_slackAttachmentField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackAttachmentField) Int(int bool) *_slackAttachmentField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackAttachmentField) Title(title string) *_slackAttachmentField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackAttachmentField) Value(value string) *_slackAttachmentField {
	_ = "STUB: not implemented"
	return nil
}

func (s *_slackAttachmentField) SlackAttachmentFieldCaster() *types.SlackAttachmentField {
	_ = "STUB: not implemented"
	return nil
}
