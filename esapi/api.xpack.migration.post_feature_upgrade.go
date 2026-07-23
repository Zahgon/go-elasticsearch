package esapi

import (
	"context"
	"net/http"
)

func newMigrationPostFeatureUpgradeFunc(t Transport) MigrationPostFeatureUpgrade {
	_ = "STUB: not implemented"
	return *new(MigrationPostFeatureUpgrade)
}

type MigrationPostFeatureUpgrade func(o ...func(*MigrationPostFeatureUpgradeRequest)) (*Response, error)

type MigrationPostFeatureUpgradeRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MigrationPostFeatureUpgradeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MigrationPostFeatureUpgrade) WithContext(v context.Context) func(*MigrationPostFeatureUpgradeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationPostFeatureUpgrade) WithPretty() func(*MigrationPostFeatureUpgradeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationPostFeatureUpgrade) WithHuman() func(*MigrationPostFeatureUpgradeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationPostFeatureUpgrade) WithErrorTrace() func(*MigrationPostFeatureUpgradeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationPostFeatureUpgrade) WithFilterPath(v ...string) func(*MigrationPostFeatureUpgradeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationPostFeatureUpgrade) WithHeader(h map[string]string) func(*MigrationPostFeatureUpgradeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationPostFeatureUpgrade) WithOpaqueID(s string) func(*MigrationPostFeatureUpgradeRequest) {
	_ = "STUB: not implemented"
	return nil
}
