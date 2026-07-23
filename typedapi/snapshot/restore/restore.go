package restore

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
	repositoryMask = iota + 1

	snapshotMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Restore struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	repository string
	snapshot   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRestore func(repository, snapshot string) *Restore

func NewRestoreFunc(tp elastictransport.Interface) NewRestore {
	_ = "STUB: not implemented"
	return *new(NewRestore)
}

func New(tp elastictransport.Interface) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) Raw(raw io.Reader) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) Request(req *Request) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Restore) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Restore) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Restore) Header(key, value string) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) _repository(repository string) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) _snapshot(snapshot string) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) MasterTimeout(duration string) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) WaitForCompletion(waitforcompletion bool) *Restore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Restore) ErrorTrace(errortrace bool) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) FilterPath(filterpaths ...string) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) Human(human bool) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) Pretty(pretty bool) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) FeatureStates(featurestates ...string) *Restore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Restore) IgnoreIndexSettings(ignoreindexsettings ...string) *Restore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Restore) IgnoreUnavailable(ignoreunavailable bool) *Restore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Restore) IncludeAliases(includealiases bool) *Restore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Restore) IncludeGlobalState(includeglobalstate bool) *Restore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Restore) IndexSettings(indexsettings types.IndexSettingsVariant) *Restore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Restore) Indices(indices ...string) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) Partial(partial bool) *Restore { _ = "STUB: not implemented"; return nil }

func (r *Restore) RenamePattern(renamepattern string) *Restore {
	_ = "STUB: not implemented"
	return nil
}

func (r *Restore) RenameReplacement(renamereplacement string) *Restore {
	_ = "STUB: not implemented"
	return nil
}
