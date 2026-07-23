package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _completionToolFunction struct {
	v *types.CompletionToolFunction
}

func NewCompletionToolFunction(name string) *_completionToolFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionToolFunction) Description(description string) *_completionToolFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionToolFunction) Name(name string) *_completionToolFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionToolFunction) Parameters(parameters json.RawMessage) *_completionToolFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionToolFunction) Strict(strict bool) *_completionToolFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_completionToolFunction) CompletionToolFunctionCaster() *types.CompletionToolFunction {
	_ = "STUB: not implemented"
	return nil
}
