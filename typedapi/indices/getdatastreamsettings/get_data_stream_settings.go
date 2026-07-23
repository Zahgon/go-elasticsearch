package getdatastreamsettings

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDataStreamSettings struct {
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

type NewGetDataStreamSettings func(name string) *GetDataStreamSettings

func NewGetDataStreamSettingsFunc(tp elastictransport.Interface) NewGetDataStreamSettings {
	_ = "STUB: not implemented"
	return *new(NewGetDataStreamSettings)
}

func New(tp elastictransport.Interface) *GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamSettings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamSettings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamSettings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamSettings) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDataStreamSettings) Header(key, value string) *GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamSettings) _name(name string) *GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamSettings) MasterTimeout(duration string) *GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamSettings) ErrorTrace(errortrace bool) *GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamSettings) FilterPath(filterpaths ...string) *GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamSettings) Human(human bool) *GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamSettings) Pretty(pretty bool) *GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}
