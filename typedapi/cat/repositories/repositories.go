package repositories

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Repositories struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRepositories func() *Repositories

func NewRepositoriesFunc(tp elastictransport.Interface) NewRepositories {
	_ = "STUB: not implemented"
	return *new(NewRepositories)
}

func New(tp elastictransport.Interface) *Repositories { _ = "STUB: not implemented"; return nil }

func (r *Repositories) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Repositories) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Repositories) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Repositories) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Repositories) Header(key, value string) *Repositories {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repositories) H(names ...string) *Repositories { _ = "STUB: not implemented"; return nil }

func (r *Repositories) S(names ...string) *Repositories { _ = "STUB: not implemented"; return nil }

func (r *Repositories) Local(local bool) *Repositories { _ = "STUB: not implemented"; return nil }

func (r *Repositories) MasterTimeout(duration string) *Repositories {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repositories) Bytes(bytes bytes.Bytes) *Repositories {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repositories) Format(format string) *Repositories { _ = "STUB: not implemented"; return nil }

func (r *Repositories) Help(help bool) *Repositories { _ = "STUB: not implemented"; return nil }

func (r *Repositories) Time(time timeunit.TimeUnit) *Repositories {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repositories) V(v bool) *Repositories { _ = "STUB: not implemented"; return nil }

func (r *Repositories) ErrorTrace(errortrace bool) *Repositories {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repositories) FilterPath(filterpaths ...string) *Repositories {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repositories) Human(human bool) *Repositories { _ = "STUB: not implemented"; return nil }

func (r *Repositories) Pretty(pretty bool) *Repositories { _ = "STUB: not implemented"; return nil }
