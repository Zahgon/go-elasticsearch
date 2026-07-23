package putuser

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	usernameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutUser struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	username string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutUser func(username string) *PutUser

func NewPutUserFunc(tp elastictransport.Interface) NewPutUser {
	_ = "STUB: not implemented"
	return *new(NewPutUser)
}

func New(tp elastictransport.Interface) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Raw(raw io.Reader) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Request(req *Request) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutUser) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutUser) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutUser) Header(key, value string) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) _username(username string) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Refresh(refresh refresh.Refresh) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) ErrorTrace(errortrace bool) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) FilterPath(filterpaths ...string) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Human(human bool) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Pretty(pretty bool) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Email(email string) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Enabled(enabled bool) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) FullName(fullname string) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Metadata(metadata types.MetadataVariant) *PutUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutUser) Password(password string) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) PasswordHash(passwordhash string) *PutUser { _ = "STUB: not implemented"; return nil }

func (r *PutUser) Roles(roles ...string) *PutUser { _ = "STUB: not implemented"; return nil }
