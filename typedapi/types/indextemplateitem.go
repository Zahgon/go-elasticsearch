package types

type IndexTemplateItem struct {
	IndexTemplate IndexTemplateWithRollover `json:"index_template"`
	Name          string                    `json:"name"`
}

func (s *IndexTemplateItem) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexTemplateItem() *IndexTemplateItem { _ = "STUB: not implemented"; return nil }
