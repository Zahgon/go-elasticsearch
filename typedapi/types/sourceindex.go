package types

type SourceIndex struct {
	Index string `json:"index"`
}

func (s *SourceIndex) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSourceIndex() *SourceIndex { _ = "STUB: not implemented"; return nil }

type SourceIndexVariant interface {
	SourceIndexCaster() *SourceIndex
}

func (s *SourceIndex) SourceIndexCaster() *SourceIndex { _ = "STUB: not implemented"; return nil }
