package esapi

import (
	"context"
	"net/http"
)

func newMigrationGetFeatureUpgradeStatusFunc(t Transport) MigrationGetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return *new(MigrationGetFeatureUpgradeStatus)
}

type MigrationGetFeatureUpgradeStatus func(o ...func(*MigrationGetFeatureUpgradeStatusRequest)) (*Response, error)

type MigrationGetFeatureUpgradeStatusRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MigrationGetFeatureUpgradeStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MigrationGetFeatureUpgradeStatus) WithContext(v context.Context) func(*MigrationGetFeatureUpgradeStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationGetFeatureUpgradeStatus) WithPretty() func(*MigrationGetFeatureUpgradeStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationGetFeatureUpgradeStatus) WithHuman() func(*MigrationGetFeatureUpgradeStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationGetFeatureUpgradeStatus) WithErrorTrace() func(*MigrationGetFeatureUpgradeStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationGetFeatureUpgradeStatus) WithFilterPath(v ...string) func(*MigrationGetFeatureUpgradeStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationGetFeatureUpgradeStatus) WithHeader(h map[string]string) func(*MigrationGetFeatureUpgradeStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MigrationGetFeatureUpgradeStatus) WithOpaqueID(s string) func(*MigrationGetFeatureUpgradeStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
