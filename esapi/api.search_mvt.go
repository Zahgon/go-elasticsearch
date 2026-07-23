package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSearchMvtFunc(t Transport) SearchMvt { _ = "STUB: not implemented"; return *new(SearchMvt) }

type SearchMvt func(index []string, field string, x *int, y *int, zoom *int, o ...func(*SearchMvtRequest)) (*Response, error)

type SearchMvtRequest struct {
	Index []string

	Body io.Reader

	Field string
	X     *int
	Y     *int
	Zoom  *int

	ExactBounds    *bool
	Extent         *int
	GridAgg        string
	GridPrecision  *int
	GridType       string
	Size           *int
	TrackTotalHits interface{}
	WithLabels     *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchMvtRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchMvt) WithContext(v context.Context) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithBody(v io.Reader) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithExactBounds(v bool) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithExtent(v int) func(*SearchMvtRequest) { _ = "STUB: not implemented"; return nil }

func (f SearchMvt) WithGridAgg(v string) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithGridPrecision(v int) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithGridType(v string) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithSize(v int) func(*SearchMvtRequest) { _ = "STUB: not implemented"; return nil }

func (f SearchMvt) WithTrackTotalHits(v interface{}) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithWithLabels(v bool) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithPretty() func(*SearchMvtRequest) { _ = "STUB: not implemented"; return nil }

func (f SearchMvt) WithHuman() func(*SearchMvtRequest) { _ = "STUB: not implemented"; return nil }

func (f SearchMvt) WithErrorTrace() func(*SearchMvtRequest) { _ = "STUB: not implemented"; return nil }

func (f SearchMvt) WithFilterPath(v ...string) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithHeader(h map[string]string) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchMvt) WithOpaqueID(s string) func(*SearchMvtRequest) {
	_ = "STUB: not implemented"
	return nil
}
