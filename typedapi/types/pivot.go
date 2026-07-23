package types

type Pivot struct {
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`

	GroupBy map[string]PivotGroupByContainer `json:"group_by,omitempty"`
}

func NewPivot() *Pivot { _ = "STUB: not implemented"; return nil }

type PivotVariant interface {
	PivotCaster() *Pivot
}

func (s *Pivot) PivotCaster() *Pivot { _ = "STUB: not implemented"; return nil }
