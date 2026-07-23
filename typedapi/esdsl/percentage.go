package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _percentage struct {
	v types.Percentage
}

func NewPercentage() *_percentage { _ = "STUB: not implemented"; return nil }

func (u *_percentage) String(string string) *_percentage { _ = "STUB: not implemented"; return nil }

func (u *_percentage) Float32(float32 float32) *_percentage { _ = "STUB: not implemented"; return nil }

func (u *_percentage) PercentageCaster() *types.Percentage { _ = "STUB: not implemented"; return nil }
