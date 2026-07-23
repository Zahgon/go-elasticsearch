package suggestuserprofiles

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

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SuggestUserProfiles struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSuggestUserProfiles func() *SuggestUserProfiles

func NewSuggestUserProfilesFunc(tp elastictransport.Interface) NewSuggestUserProfiles {
	_ = "STUB: not implemented"
	return *new(NewSuggestUserProfiles)
}

func New(tp elastictransport.Interface) *SuggestUserProfiles { _ = "STUB: not implemented"; return nil }

func (r *SuggestUserProfiles) Raw(raw io.Reader) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) Request(req *Request) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SuggestUserProfiles) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SuggestUserProfiles) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SuggestUserProfiles) Header(key, value string) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) ErrorTrace(errortrace bool) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) FilterPath(filterpaths ...string) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) Human(human bool) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) Pretty(pretty bool) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) Data(data ...string) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) Hint(hint types.HintVariant) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) Name(name string) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (r *SuggestUserProfiles) Size(size int64) *SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}
