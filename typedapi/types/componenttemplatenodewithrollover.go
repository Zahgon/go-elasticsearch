package types

type ComponentTemplateNodeWithRollover struct {
	CreatedDate DateTime `json:"created_date,omitempty"`

	CreatedDateMillis *int64   `json:"created_date_millis,omitempty"`
	Deprecated        *bool    `json:"deprecated,omitempty"`
	Meta_             Metadata `json:"_meta,omitempty"`

	ModifiedDate DateTime `json:"modified_date,omitempty"`

	ModifiedDateMillis *int64                      `json:"modified_date_millis,omitempty"`
	Template           ComponentTemplateSummaryRes `json:"template"`
	Version            *int64                      `json:"version,omitempty"`
}

func (s *ComponentTemplateNodeWithRollover) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewComponentTemplateNodeWithRollover() *ComponentTemplateNodeWithRollover {
	_ = "STUB: not implemented"
	return nil
}
