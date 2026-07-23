package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _duration struct {
	v types.Duration
}

func NewDuration() *_duration { _ = "STUB: not implemented"; return nil }

func (u *_duration) String(string string) *_duration { _ = "STUB: not implemented"; return nil }

func (u *_duration) DurationCaster() *types.Duration { _ = "STUB: not implemented"; return nil }
