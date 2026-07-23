package types

type FetchProfileDebug struct {
	FastPath     *int     `json:"fast_path,omitempty"`
	StoredFields []string `json:"stored_fields,omitempty"`
}

func (s *FetchProfileDebug) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFetchProfileDebug() *FetchProfileDebug { _ = "STUB: not implemented"; return nil }
