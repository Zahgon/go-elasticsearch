package types

type Lifecycle struct {
	ModifiedDate DateTime  `json:"modified_date"`
	Policy       IlmPolicy `json:"policy"`
	Version      int64     `json:"version"`
}

func (s *Lifecycle) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLifecycle() *Lifecycle { _ = "STUB: not implemented"; return nil }
