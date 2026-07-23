package types

type Suggester struct {
	Suggesters map[string]FieldSuggester `json:"-"`

	Text *string `json:"text,omitempty"`
}

func (s *Suggester) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s Suggester) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSuggester() *Suggester { _ = "STUB: not implemented"; return nil }

type SuggesterVariant interface {
	SuggesterCaster() *Suggester
}

func (s *Suggester) SuggesterCaster() *Suggester { _ = "STUB: not implemented"; return nil }
