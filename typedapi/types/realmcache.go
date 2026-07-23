package types

type RealmCache struct {
	Size int64 `json:"size"`
}

func (s *RealmCache) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRealmCache() *RealmCache { _ = "STUB: not implemented"; return nil }
