package types

type MatrixStatsAggregate struct {
	DocCount int64               `json:"doc_count"`
	Fields   []MatrixStatsFields `json:"fields,omitempty"`
	Meta     Metadata            `json:"meta,omitempty"`
}

func (s *MatrixStatsAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMatrixStatsAggregate() *MatrixStatsAggregate { _ = "STUB: not implemented"; return nil }
