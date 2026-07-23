package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dataattachmentformat"
)

type _dataEmailAttachment struct {
	v *types.DataEmailAttachment
}

func NewDataEmailAttachment() *_dataEmailAttachment { _ = "STUB: not implemented"; return nil }

func (s *_dataEmailAttachment) Format(format dataattachmentformat.DataAttachmentFormat) *_dataEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataEmailAttachment) EmailAttachmentContainerCaster() *types.EmailAttachmentContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataEmailAttachment) DataEmailAttachmentCaster() *types.DataEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}
