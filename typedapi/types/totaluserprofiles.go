package types

type TotalUserProfiles struct {
	Relation string `json:"relation"`
	Value    int64  `json:"value"`
}

func (s *TotalUserProfiles) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTotalUserProfiles() *TotalUserProfiles { _ = "STUB: not implemented"; return nil }
