package types

type SlackMessage struct {
	Attachments        []SlackAttachment       `json:"attachments"`
	DynamicAttachments *SlackDynamicAttachment `json:"dynamic_attachments,omitempty"`
	From               string                  `json:"from"`
	Icon               *string                 `json:"icon,omitempty"`
	Text               string                  `json:"text"`
	To                 []string                `json:"to"`
}

func (s *SlackMessage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSlackMessage() *SlackMessage { _ = "STUB: not implemented"; return nil }

type SlackMessageVariant interface {
	SlackMessageCaster() *SlackMessage
}

func (s *SlackMessage) SlackMessageCaster() *SlackMessage { _ = "STUB: not implemented"; return nil }
