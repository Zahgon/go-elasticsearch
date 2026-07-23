package types

type UnratedDocument struct {
	Id_    string `json:"_id"`
	Index_ string `json:"_index"`
}

func (s *UnratedDocument) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUnratedDocument() *UnratedDocument { _ = "STUB: not implemented"; return nil }
