package types

type ScriptScoreQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	MinScore *float32 `json:"min_score,omitempty"`

	Query      Query   `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`

	Script Script `json:"script"`
}

func (s *ScriptScoreQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScriptScoreQuery() *ScriptScoreQuery { _ = "STUB: not implemented"; return nil }

type ScriptScoreQueryVariant interface {
	ScriptScoreQueryCaster() *ScriptScoreQuery
}

func (s *ScriptScoreQuery) ScriptScoreQueryCaster() *ScriptScoreQuery {
	_ = "STUB: not implemented"
	return nil
}
