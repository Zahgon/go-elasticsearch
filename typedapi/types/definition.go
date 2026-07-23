package types

type Definition struct {
	Preprocessors []Preprocessor `json:"preprocessors,omitempty"`

	TrainedModel TrainedModel `json:"trained_model"`
}

func NewDefinition() *Definition { _ = "STUB: not implemented"; return nil }

type DefinitionVariant interface {
	DefinitionCaster() *Definition
}

func (s *Definition) DefinitionCaster() *Definition { _ = "STUB: not implemented"; return nil }
