package types

type ShardsStatsSummaryItem struct {
	FileCount   int64 `json:"file_count"`
	SizeInBytes int64 `json:"size_in_bytes"`
}

func (s *ShardsStatsSummaryItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardsStatsSummaryItem() *ShardsStatsSummaryItem { _ = "STUB: not implemented"; return nil }
