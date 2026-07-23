package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _watcherInput struct {
	v *types.WatcherInput
}

func NewWatcherInput() *_watcherInput { _ = "STUB: not implemented"; return nil }

func (s *_watcherInput) Chain(chain types.ChainInputVariant) *_watcherInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherInput) Http(http types.HttpInputVariant) *_watcherInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherInput) Search(search types.SearchInputVariant) *_watcherInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherInput) Simple(simple map[string]json.RawMessage) *_watcherInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherInput) AddSimple(key string, value json.RawMessage) *_watcherInput {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherInput) WatcherInputCaster() *types.WatcherInput {
	_ = "STUB: not implemented"
	return nil
}
