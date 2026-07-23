package types

type ResolveIndexAliasItem struct {
	Indices []string `json:"indices"`
	Name    string   `json:"name"`
}

func (s *ResolveIndexAliasItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewResolveIndexAliasItem() *ResolveIndexAliasItem { _ = "STUB: not implemented"; return nil }
