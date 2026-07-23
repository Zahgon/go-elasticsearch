package types

type Eql struct {
	Available bool                  `json:"available"`
	Enabled   bool                  `json:"enabled"`
	Features  EqlFeatures           `json:"features"`
	Queries   map[string]XpackQuery `json:"queries"`
}

func (s *Eql) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEql() *Eql { _ = "STUB: not implemented"; return nil }
