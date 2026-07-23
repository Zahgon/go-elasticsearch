package esapi

import (
	"context"
	"net/http"
)

func newCatTransformsFunc(t Transport) CatTransforms {
	_ = "STUB: not implemented"
	return *new(CatTransforms)
}

type CatTransforms func(o ...func(*CatTransformsRequest)) (*Response, error)

type CatTransformsRequest struct {
	TransformID string

	AllowNoMatch *bool
	Bytes        string
	Format       string
	From         *int
	H            []string
	Help         *bool
	S            []string
	Size         *int
	Time         string
	V            *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatTransformsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatTransforms) WithContext(v context.Context) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithTransformID(v string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithAllowNoMatch(v bool) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithBytes(v string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithFormat(v string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithFrom(v int) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithH(v ...string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithHelp(v bool) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithS(v ...string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithSize(v int) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithTime(v string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithV(v bool) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithPretty() func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithHuman() func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithErrorTrace() func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithFilterPath(v ...string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithHeader(h map[string]string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatTransforms) WithOpaqueID(s string) func(*CatTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}
