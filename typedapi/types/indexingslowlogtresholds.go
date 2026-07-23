package types

type IndexingSlowlogTresholds struct {
	Index *SlowlogTresholdLevels `json:"index,omitempty"`
}

func NewIndexingSlowlogTresholds() *IndexingSlowlogTresholds { _ = "STUB: not implemented"; return nil }

type IndexingSlowlogTresholdsVariant interface {
	IndexingSlowlogTresholdsCaster() *IndexingSlowlogTresholds
}

func (s *IndexingSlowlogTresholds) IndexingSlowlogTresholdsCaster() *IndexingSlowlogTresholds {
	_ = "STUB: not implemented"
	return nil
}
