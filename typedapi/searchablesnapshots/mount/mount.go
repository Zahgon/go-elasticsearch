package mount

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/storageoption"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Mount struct {
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

type NewMount func(repository, snapshot string) *Mount

func NewMountFunc(tp elastictransport.Interface) NewMount {
	_ = "STUB: not implemented"
	return *new(NewMount)
}

func New(tp elastictransport.Interface) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) Raw(raw io.Reader) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) Request(req *Request) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Mount) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Mount) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Mount) Header(key, value string) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) _repository(repository string) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) _snapshot(snapshot string) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) MasterTimeout(duration string) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) WaitForCompletion(waitforcompletion bool) *Mount {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mount) Storage(storage storageoption.StorageOption) *Mount {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mount) ErrorTrace(errortrace bool) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) FilterPath(filterpaths ...string) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) Human(human bool) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) Pretty(pretty bool) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) IgnoreIndexSettings(ignoreindexsettings ...string) *Mount {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mount) Index(indexname string) *Mount { _ = "STUB: not implemented"; return nil }

func (r *Mount) IndexSettings(indexsettings map[string]json.RawMessage) *Mount {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mount) AddIndexSetting(key string, value json.RawMessage) *Mount {
	_ = "STUB: not implemented"
	return nil
}

func (r *Mount) RenamedIndex(indexname string) *Mount { _ = "STUB: not implemented"; return nil }
