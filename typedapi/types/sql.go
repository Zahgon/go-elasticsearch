package types

type Sql struct {
	Available bool                  `json:"available"`
	Enabled   bool                  `json:"enabled"`
	Features  map[string]int        `json:"features"`
	Queries   map[string]XpackQuery `json:"queries"`
}

func (s *Sql) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSql() *Sql { _ = "STUB: not implemented"; return nil }
