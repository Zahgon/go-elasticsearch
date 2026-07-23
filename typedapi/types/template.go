package types

type Template struct {
	Aliases  map[string]Alias `json:"aliases"`
	Mappings TypeMapping      `json:"mappings"`
	Settings IndexSettings    `json:"settings"`
}

func NewTemplate() *Template { _ = "STUB: not implemented"; return nil }
