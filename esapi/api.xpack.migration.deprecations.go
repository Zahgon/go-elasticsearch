package esapi

import (
	"context"
	"net/http"
)

func newMigrationDeprecationsFunc(t Transport) MigrationDeprecations {
	_ = "STUB: not implemented"
	return *new(MigrationDeprecations)
}

type MigrationDeprecations func(o ...func(*MigrationDeprecationsRequest)) (*Response, error)

type MigrationDeprecationsRequest struct {
	Index string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MigrationDeprecationsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MigrationDeprecations) WithContext(v context.Context) func(*MigrationDeprecationsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationDeprecations) WithIndex(v string) func(*MigrationDeprecationsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationDeprecations) WithPretty() func(*MigrationDeprecationsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationDeprecations) WithHuman() func(*MigrationDeprecationsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationDeprecations) WithErrorTrace() func(*MigrationDeprecationsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationDeprecations) WithFilterPath(v ...string) func(*MigrationDeprecationsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationDeprecations) WithHeader(h map[string]string) func(*MigrationDeprecationsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationDeprecations) WithOpaqueID(s string) func(*MigrationDeprecationsRequest) {
	_ = "STUB: not implemented"
	return nil
}
