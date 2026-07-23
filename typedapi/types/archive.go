package types

type Archive struct {
	Available    bool  `json:"available"`
	Enabled      bool  `json:"enabled"`
	IndicesCount int64 `json:"indices_count"`
}

func (s *Archive) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewArchive() *Archive { _ = "STUB: not implemented"; return nil }
