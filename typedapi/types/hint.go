package types

type Hint struct {
	Labels map[string][]string `json:"labels,omitempty"`

	Uids []string `json:"uids,omitempty"`
}

func (s *Hint) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHint() *Hint { _ = "STUB: not implemented"; return nil }

type HintVariant interface {
	HintCaster() *Hint
}

func (s *Hint) HintCaster() *Hint { _ = "STUB: not implemented"; return nil }
