package deletevotingconfigexclusions

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteVotingConfigExclusions struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteVotingConfigExclusions func() *DeleteVotingConfigExclusions

func NewDeleteVotingConfigExclusionsFunc(tp elastictransport.Interface) NewDeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return *new(NewDeleteVotingConfigExclusions)
}

func New(tp elastictransport.Interface) *DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteVotingConfigExclusions) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteVotingConfigExclusions) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteVotingConfigExclusions) Do(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r DeleteVotingConfigExclusions) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteVotingConfigExclusions) Header(key, value string) *DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteVotingConfigExclusions) MasterTimeout(duration string) *DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteVotingConfigExclusions) WaitForRemoval(waitforremoval bool) *DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteVotingConfigExclusions) ErrorTrace(errortrace bool) *DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteVotingConfigExclusions) FilterPath(filterpaths ...string) *DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteVotingConfigExclusions) Human(human bool) *DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteVotingConfigExclusions) Pretty(pretty bool) *DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}
