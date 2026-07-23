package types

type NotFoundAliases struct {
	Error           string                  `json:"error"`
	NotFoundAliases map[string]IndexAliases `json:"-"`
	Status          int                     `json:"status"`
}

func (s *NotFoundAliases) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NotFoundAliases) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNotFoundAliases() *NotFoundAliases { _ = "STUB: not implemented"; return nil }
