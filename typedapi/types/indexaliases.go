package types

type IndexAliases struct {
	Aliases map[string]AliasDefinition `json:"aliases"`
}

func NewIndexAliases() *IndexAliases { _ = "STUB: not implemented"; return nil }
