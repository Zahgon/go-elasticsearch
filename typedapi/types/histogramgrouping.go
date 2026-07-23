package types

type HistogramGrouping struct {
	Fields []string `json:"fields"`

	Interval int64 `json:"interval"`
}

func (s *HistogramGrouping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHistogramGrouping() *HistogramGrouping { _ = "STUB: not implemented"; return nil }

type HistogramGroupingVariant interface {
	HistogramGroupingCaster() *HistogramGrouping
}

func (s *HistogramGrouping) HistogramGroupingCaster() *HistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}
