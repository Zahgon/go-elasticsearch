package types

type SlackAttachment struct {
	AuthorIcon *string                `json:"author_icon,omitempty"`
	AuthorLink *string                `json:"author_link,omitempty"`
	AuthorName string                 `json:"author_name"`
	Color      *string                `json:"color,omitempty"`
	Fallback   *string                `json:"fallback,omitempty"`
	Fields     []SlackAttachmentField `json:"fields,omitempty"`
	Footer     *string                `json:"footer,omitempty"`
	FooterIcon *string                `json:"footer_icon,omitempty"`
	ImageUrl   *string                `json:"image_url,omitempty"`
	Pretext    *string                `json:"pretext,omitempty"`
	Text       *string                `json:"text,omitempty"`
	ThumbUrl   *string                `json:"thumb_url,omitempty"`
	Title      string                 `json:"title"`
	TitleLink  *string                `json:"title_link,omitempty"`
	Ts         *int64                 `json:"ts,omitempty"`
}

func (s *SlackAttachment) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSlackAttachment() *SlackAttachment { _ = "STUB: not implemented"; return nil }

type SlackAttachmentVariant interface {
	SlackAttachmentCaster() *SlackAttachment
}

func (s *SlackAttachment) SlackAttachmentCaster() *SlackAttachment {
	_ = "STUB: not implemented"
	return nil
}
