package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/emailpriority"
)

type _emailAction struct {
	v *types.EmailAction
}

func NewEmailAction(subject string) *_emailAction { _ = "STUB: not implemented"; return nil }

func (s *_emailAction) Attachments(attachments map[string]types.EmailAttachmentContainer) *_emailAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAction) AddAttachment(key string, value types.EmailAttachmentContainerVariant) *_emailAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAction) Bcc(bccs ...string) *_emailAction { _ = "STUB: not implemented"; return nil }

func (s *_emailAction) Body(body types.EmailBodyVariant) *_emailAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAction) Cc(ccs ...string) *_emailAction { _ = "STUB: not implemented"; return nil }

func (s *_emailAction) From(from string) *_emailAction { _ = "STUB: not implemented"; return nil }

func (s *_emailAction) Id(id string) *_emailAction { _ = "STUB: not implemented"; return nil }

func (s *_emailAction) Priority(priority emailpriority.EmailPriority) *_emailAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAction) ReplyTo(replytos ...string) *_emailAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAction) SentDate(datetime types.DateTimeVariant) *_emailAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_emailAction) Subject(subject string) *_emailAction { _ = "STUB: not implemented"; return nil }

func (s *_emailAction) To(tos ...string) *_emailAction { _ = "STUB: not implemented"; return nil }

func (s *_emailAction) EmailActionCaster() *types.EmailAction {
	_ = "STUB: not implemented"
	return nil
}
