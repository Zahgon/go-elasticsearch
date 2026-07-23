package types

type PinnedDoc struct {
	Id_ string `json:"_id"`

	Index_ *string `json:"_index,omitempty"`
}

func (s *PinnedDoc) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPinnedDoc() *PinnedDoc { _ = "STUB: not implemented"; return nil }

type PinnedDocVariant interface {
	PinnedDocCaster() *PinnedDoc
}

func (s *PinnedDoc) PinnedDocCaster() *PinnedDoc { _ = "STUB: not implemented"; return nil }
