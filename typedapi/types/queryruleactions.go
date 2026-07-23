package types

type QueryRuleActions struct {
	Docs []PinnedDoc `json:"docs,omitempty"`

	Ids []string `json:"ids,omitempty"`
}

func NewQueryRuleActions() *QueryRuleActions { _ = "STUB: not implemented"; return nil }

type QueryRuleActionsVariant interface {
	QueryRuleActionsCaster() *QueryRuleActions
}

func (s *QueryRuleActions) QueryRuleActionsCaster() *QueryRuleActions {
	_ = "STUB: not implemented"
	return nil
}
