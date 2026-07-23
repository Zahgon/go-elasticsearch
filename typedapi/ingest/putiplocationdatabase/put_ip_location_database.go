package putiplocationdatabase

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

type PutIpLocationDatabase struct {
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

type NewPutIpLocationDatabase func(id string) *PutIpLocationDatabase

func NewPutIpLocationDatabaseFunc(tp elastictransport.Interface) NewPutIpLocationDatabase {
	_ = "STUB: not implemented"
	return *new(NewPutIpLocationDatabase)
}

func New(tp elastictransport.Interface) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) Raw(raw io.Reader) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) Request(req *Request) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutIpLocationDatabase) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutIpLocationDatabase) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutIpLocationDatabase) Header(key, value string) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) _id(id string) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) MasterTimeout(duration string) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) Timeout(duration string) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) ErrorTrace(errortrace bool) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) FilterPath(filterpaths ...string) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) Human(human bool) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) Pretty(pretty bool) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) Ipinfo(ipinfo types.IpinfoVariant) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) Maxmind(maxmind types.MaxmindVariant) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIpLocationDatabase) Name(name string) *PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}
