package types

type AliasesRecord struct {
	Alias *string `json:"alias,omitempty"`

	Filter *string `json:"filter,omitempty"`

	Index *string `json:"index,omitempty"`

	IsWriteIndex *string `json:"is_write_index,omitempty"`

	RoutingIndex *string `json:"routing.index,omitempty"`

	RoutingSearch *string `json:"routing.search,omitempty"`
}

func (s *AliasesRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAliasesRecord() *AliasesRecord { _ = "STUB: not implemented"; return nil }
