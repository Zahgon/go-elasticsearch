package types

type Groupings struct {
	DateHistogram *DateHistogramGrouping `json:"date_histogram,omitempty"`

	Histogram *HistogramGrouping `json:"histogram,omitempty"`

	Terms *TermsGrouping `json:"terms,omitempty"`
}

func NewGroupings() *Groupings { _ = "STUB: not implemented"; return nil }

type GroupingsVariant interface {
	GroupingsCaster() *Groupings
}

func (s *Groupings) GroupingsCaster() *Groupings { _ = "STUB: not implemented"; return nil }
