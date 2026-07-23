package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatAllocationFunc(t Transport) CatAllocation {
	_ = "STUB: not implemented"
	return *new(CatAllocation)
}

type CatAllocation func(o ...func(*CatAllocationRequest)) (*Response, error)

type CatAllocationRequest struct {
	NodeID []string

	Bytes         string
	Format        string
	H             []string
	Help          *bool
	Local         *bool
	MasterTimeout time.Duration
	S             []string
	Time          string
	V             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatAllocationRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatAllocation) WithContext(v context.Context) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithNodeID(v ...string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithBytes(v string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithFormat(v string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithH(v ...string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithHelp(v bool) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithLocal(v bool) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithMasterTimeout(v time.Duration) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithS(v ...string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithTime(v string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithV(v bool) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithPretty() func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithHuman() func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithErrorTrace() func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithFilterPath(v ...string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithHeader(h map[string]string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatAllocation) WithOpaqueID(s string) func(*CatAllocationRequest) {
	_ = "STUB: not implemented"
	return nil
}
