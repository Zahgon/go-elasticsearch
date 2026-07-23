package types

type IndexAndDataStreamAction struct {
	DataStream string `json:"data_stream"`

	Index string `json:"index"`
}

func (s *IndexAndDataStreamAction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexAndDataStreamAction() *IndexAndDataStreamAction { _ = "STUB: not implemented"; return nil }

type IndexAndDataStreamActionVariant interface {
	IndexAndDataStreamActionCaster() *IndexAndDataStreamAction
}

func (s *IndexAndDataStreamAction) IndexAndDataStreamActionCaster() *IndexAndDataStreamAction {
	_ = "STUB: not implemented"
	return nil
}
