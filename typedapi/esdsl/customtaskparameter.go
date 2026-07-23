package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _customTaskParameter struct {
	v types.CustomTaskParameter
}

func NewCustomTaskParameter() *_customTaskParameter { _ = "STUB: not implemented"; return nil }

func (u *_customTaskParameter) String(string string) *_customTaskParameter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_customTaskParameter) Int(int int) *_customTaskParameter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_customTaskParameter) Float64(float64 types.Float64) *_customTaskParameter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_customTaskParameter) Float32(float32 float32) *_customTaskParameter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_customTaskParameter) Bool(bool bool) *_customTaskParameter {
	_ = "STUB: not implemented"
	return nil
}

func (u *_customTaskParameter) CustomTaskParameterCaster() *types.CustomTaskParameter {
	_ = "STUB: not implemented"
	return nil
}
