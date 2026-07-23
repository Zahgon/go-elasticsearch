package types

type RRFRetriever struct {
	Fields []string `json:"fields,omitempty"`

	Filter []Query `json:"filter,omitempty"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_ *string `json:"_name,omitempty"`
	Query *string `json:"query,omitempty"`

	RankConstant *int `json:"rank_constant,omitempty"`

	RankWindowSize *int `json:"rank_window_size,omitempty"`

	Retrievers []RRFRetrieverEntry `json:"retrievers"`
}

func (s *RRFRetriever) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRRFRetriever() *RRFRetriever { _ = "STUB: not implemented"; return nil }

type RRFRetrieverVariant interface {
	RRFRetrieverCaster() *RRFRetriever
}

func (s *RRFRetriever) RRFRetrieverCaster() *RRFRetriever { _ = "STUB: not implemented"; return nil }
