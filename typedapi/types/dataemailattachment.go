package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dataattachmentformat"
)

type DataEmailAttachment struct {
	Format *dataattachmentformat.DataAttachmentFormat `json:"format,omitempty"`
}

func NewDataEmailAttachment() *DataEmailAttachment { _ = "STUB: not implemented"; return nil }

type DataEmailAttachmentVariant interface {
	DataEmailAttachmentCaster() *DataEmailAttachment
}

func (s *DataEmailAttachment) DataEmailAttachmentCaster() *DataEmailAttachment {
	_ = "STUB: not implemented"
	return nil
}
