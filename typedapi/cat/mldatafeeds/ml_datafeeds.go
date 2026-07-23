package mldatafeeds

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catdatafeedcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	datafeedidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlDatafeeds struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewMlDatafeeds func() *MlDatafeeds

func NewMlDatafeedsFunc(tp elastictransport.Interface) NewMlDatafeeds {
	_ = "STUB: not implemented"
	return *new(NewMlDatafeeds)
}

func New(tp elastictransport.Interface) *MlDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *MlDatafeeds) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MlDatafeeds) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MlDatafeeds) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r MlDatafeeds) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *MlDatafeeds) Header(key, value string) *MlDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *MlDatafeeds) DatafeedId(datafeedid string) *MlDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDatafeeds) AllowNoMatch(allownomatch bool) *MlDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDatafeeds) H(catdatafeedcolumns ...catdatafeedcolumn.CatDatafeedColumn) *MlDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDatafeeds) S(catdatafeedcolumns ...catdatafeedcolumn.CatDatafeedColumn) *MlDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDatafeeds) Bytes(bytes bytes.Bytes) *MlDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *MlDatafeeds) Format(format string) *MlDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *MlDatafeeds) Help(help bool) *MlDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *MlDatafeeds) Time(time timeunit.TimeUnit) *MlDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDatafeeds) V(v bool) *MlDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *MlDatafeeds) ErrorTrace(errortrace bool) *MlDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDatafeeds) FilterPath(filterpaths ...string) *MlDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlDatafeeds) Human(human bool) *MlDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *MlDatafeeds) Pretty(pretty bool) *MlDatafeeds { _ = "STUB: not implemented"; return nil }
