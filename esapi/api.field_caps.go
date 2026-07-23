package esapi

import (
	"context"
	"io"
	"net/http"
)

func newFieldCapsFunc(t Transport) FieldCaps { _ = "STUB: not implemented"; return *new(FieldCaps) }

type FieldCaps func(o ...func(*FieldCapsRequest)) (*Response, error)

type FieldCapsRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices     *bool
	ExpandWildcards    []string
	Fields             []string
	Filters            []string
	IgnoreUnavailable  *bool
	IncludeEmptyFields *bool
	IncludeUnmapped    *bool
	Types              []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FieldCapsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FieldCaps) WithContext(v context.Context) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithBody(v io.Reader) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithIndex(v ...string) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithAllowNoIndices(v bool) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithExpandWildcards(v ...string) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithFields(v ...string) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithFilters(v ...string) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithIgnoreUnavailable(v bool) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithIncludeEmptyFields(v bool) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithIncludeUnmapped(v bool) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithTypes(v ...string) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithPretty() func(*FieldCapsRequest) { _ = "STUB: not implemented"; return nil }

func (f FieldCaps) WithHuman() func(*FieldCapsRequest) { _ = "STUB: not implemented"; return nil }

func (f FieldCaps) WithErrorTrace() func(*FieldCapsRequest) { _ = "STUB: not implemented"; return nil }

func (f FieldCaps) WithFilterPath(v ...string) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithHeader(h map[string]string) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FieldCaps) WithOpaqueID(s string) func(*FieldCapsRequest) {
	_ = "STUB: not implemented"
	return nil
}
