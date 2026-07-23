package mldataframeanalytics

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catdfacolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlDataFrameAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewMlDataFrameAnalytics func() *MlDataFrameAnalytics

func NewMlDataFrameAnalyticsFunc(tp elastictransport.Interface) NewMlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewMlDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MlDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MlDataFrameAnalytics) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r MlDataFrameAnalytics) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *MlDataFrameAnalytics) Header(key, value string) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) Id(id string) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) AllowNoMatch(allownomatch bool) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) H(catdfacolumns ...catdfacolumn.CatDfaColumn) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) S(catdfacolumns ...catdfacolumn.CatDfaColumn) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) Bytes(bytes bytes.Bytes) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) Format(format string) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) Help(help bool) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) Time(time timeunit.TimeUnit) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) V(v bool) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) ErrorTrace(errortrace bool) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) FilterPath(filterpaths ...string) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) Human(human bool) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDataFrameAnalytics) Pretty(pretty bool) *MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
