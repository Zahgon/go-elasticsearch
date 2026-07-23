package datastreamsstats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DataStreamsStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDataStreamsStats func() *DataStreamsStats

func NewDataStreamsStatsFunc(tp elastictransport.Interface) NewDataStreamsStats {
	_ = "STUB: not implemented"
	return *new(NewDataStreamsStats)
}

func New(tp elastictransport.Interface) *DataStreamsStats { _ = "STUB: not implemented"; return nil }

func (r *DataStreamsStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DataStreamsStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DataStreamsStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DataStreamsStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DataStreamsStats) Header(key, value string) *DataStreamsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *DataStreamsStats) Name(name string) *DataStreamsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *DataStreamsStats) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DataStreamsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *DataStreamsStats) ErrorTrace(errortrace bool) *DataStreamsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *DataStreamsStats) FilterPath(filterpaths ...string) *DataStreamsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *DataStreamsStats) Human(human bool) *DataStreamsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *DataStreamsStats) Pretty(pretty bool) *DataStreamsStats {
	_ = "STUB: not implemented"
	return nil
}
