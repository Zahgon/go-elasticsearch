package types

type RealmInfo struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s *RealmInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRealmInfo() *RealmInfo { _ = "STUB: not implemented"; return nil }
