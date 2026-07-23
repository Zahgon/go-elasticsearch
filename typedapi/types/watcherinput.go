package types

import (
	"encoding/json"
)

type WatcherInput struct {
	Chain  *ChainInput                `json:"chain,omitempty"`
	Http   *HttpInput                 `json:"http,omitempty"`
	Search *SearchInput               `json:"search,omitempty"`
	Simple map[string]json.RawMessage `json:"simple,omitempty"`
}

func NewWatcherInput() *WatcherInput { _ = "STUB: not implemented"; return nil }

type WatcherInputVariant interface {
	WatcherInputCaster() *WatcherInput
}

func (s *WatcherInput) WatcherInputCaster() *WatcherInput { _ = "STUB: not implemented"; return nil }
