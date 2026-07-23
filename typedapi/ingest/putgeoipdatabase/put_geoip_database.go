package putgeoipdatabase

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutGeoipDatabase struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutGeoipDatabase func(id string) *PutGeoipDatabase

func NewPutGeoipDatabaseFunc(tp elastictransport.Interface) NewPutGeoipDatabase {
	_ = "STUB: not implemented"
	return *new(NewPutGeoipDatabase)
}

func New(tp elastictransport.Interface) *PutGeoipDatabase { _ = "STUB: not implemented"; return nil }

func (r *PutGeoipDatabase) Raw(raw io.Reader) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) Request(req *Request) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutGeoipDatabase) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutGeoipDatabase) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutGeoipDatabase) Header(key, value string) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) _id(id string) *PutGeoipDatabase { _ = "STUB: not implemented"; return nil }

func (r *PutGeoipDatabase) MasterTimeout(duration string) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) Timeout(duration string) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) ErrorTrace(errortrace bool) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) FilterPath(filterpaths ...string) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) Human(human bool) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) Pretty(pretty bool) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) Maxmind(maxmind types.MaxmindVariant) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGeoipDatabase) Name(name string) *PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}
