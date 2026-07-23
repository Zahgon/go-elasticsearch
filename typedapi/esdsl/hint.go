package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hint struct {
	v *types.Hint
}

func NewHint() *_hint { _ = "STUB: not implemented"; return nil }

func (s *_hint) Labels(labels map[string][]string) *_hint { _ = "STUB: not implemented"; return nil }

func (s *_hint) Uids(uids ...string) *_hint { _ = "STUB: not implemented"; return nil }

func (s *_hint) HintCaster() *types.Hint { _ = "STUB: not implemented"; return nil }
