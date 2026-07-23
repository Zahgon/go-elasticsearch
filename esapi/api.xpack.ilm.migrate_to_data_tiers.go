package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newILMMigrateToDataTiersFunc(t Transport) ILMMigrateToDataTiers {
	_ = "STUB: not implemented"
	return *new(ILMMigrateToDataTiers)
}

type ILMMigrateToDataTiers func(o ...func(*ILMMigrateToDataTiersRequest)) (*Response, error)

type ILMMigrateToDataTiersRequest struct {
	Body io.Reader

	DryRun        *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ILMMigrateToDataTiersRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMMigrateToDataTiers) WithContext(v context.Context) func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithBody(v io.Reader) func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithDryRun(v bool) func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithMasterTimeout(v time.Duration) func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithPretty() func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithHuman() func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithErrorTrace() func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithFilterPath(v ...string) func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithHeader(h map[string]string) func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMMigrateToDataTiers) WithOpaqueID(s string) func(*ILMMigrateToDataTiersRequest) {
	_ = "STUB: not implemented"
	return nil
}
