package types

type SummaryInfo struct {
	Read ReadSummaryInfo `json:"read"`

	Write WriteSummaryInfo `json:"write"`
}

func NewSummaryInfo() *SummaryInfo { _ = "STUB: not implemented"; return nil }
