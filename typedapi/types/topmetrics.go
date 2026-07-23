package types

type TopMetrics struct {
	Metrics map[string]FieldValue `json:"metrics"`
	Sort    []FieldValue          `json:"sort"`
}

func NewTopMetrics() *TopMetrics { _ = "STUB: not implemented"; return nil }
