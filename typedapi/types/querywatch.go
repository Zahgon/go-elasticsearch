package types

type QueryWatch struct {
	Id_          string       `json:"_id"`
	PrimaryTerm_ *int         `json:"_primary_term,omitempty"`
	SeqNo_       *int64       `json:"_seq_no,omitempty"`
	Status       *WatchStatus `json:"status,omitempty"`
	Watch        *Watch       `json:"watch,omitempty"`
}

func (s *QueryWatch) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewQueryWatch() *QueryWatch { _ = "STUB: not implemented"; return nil }
