package postvotingconfigexclusions

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostVotingConfigExclusions struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPostVotingConfigExclusions func() *PostVotingConfigExclusions

func NewPostVotingConfigExclusionsFunc(tp elastictransport.Interface) NewPostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return *new(NewPostVotingConfigExclusions)
}

func New(tp elastictransport.Interface) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostVotingConfigExclusions) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostVotingConfigExclusions) Do(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r PostVotingConfigExclusions) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PostVotingConfigExclusions) Header(key, value string) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) NodeNames(names ...string) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) NodeIds(ids ...string) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) MasterTimeout(duration string) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) Timeout(duration string) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) ErrorTrace(errortrace bool) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) FilterPath(filterpaths ...string) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) Human(human bool) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostVotingConfigExclusions) Pretty(pretty bool) *PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}
