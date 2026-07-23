package types

type SortOptions struct {
	Doc_         *ScoreSort           `json:"_doc,omitempty"`
	GeoDistance_ *GeoDistanceSort     `json:"_geo_distance,omitempty"`
	Score_       *ScoreSort           `json:"_score,omitempty"`
	Script_      *ScriptSort          `json:"_script,omitempty"`
	SortOptions  map[string]FieldSort `json:"-"`
}

func (s *SortOptions) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SortOptions) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSortOptions() *SortOptions { _ = "STUB: not implemented"; return nil }

type SortOptionsVariant interface {
	SortOptionsCaster() *SortOptions
}

func (s *SortOptions) SortOptionsCaster() *SortOptions { _ = "STUB: not implemented"; return nil }

func (s *SortOptions) SortCombinationsCaster() *SortCombinations {
	_ = "STUB: not implemented"
	return nil
}
