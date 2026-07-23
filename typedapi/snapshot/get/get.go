package get

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/snapshotsort"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/snapshotstate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Get struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	repository string
	snapshot   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGet func(repository, snapshot string) *Get

func NewGetFunc(tp elastictransport.Interface) NewGet {
	_ = "STUB: not implemented"
	return *new(NewGet)
}

func New(tp elastictransport.Interface) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Get) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Get) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Get) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Get) Header(key, value string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) _repository(repository string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) _snapshot(snapshot string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) After(after string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) FromSortValue(fromsortvalue string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) IgnoreUnavailable(ignoreunavailable bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) IndexDetails(indexdetails bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) IndexNames(indexnames bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) IncludeRepository(includerepository bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) MasterTimeout(duration string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) Order(order sortorder.SortOrder) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) Offset(offset int) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) Size(size int) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) SlmPolicyFilter(name string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) Sort(sort snapshotsort.SnapshotSort) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) State(states ...snapshotstate.SnapshotState) *Get {
	_ = "STUB: not implemented"
	return nil
}

func (r *Get) Verbose(verbose bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) ErrorTrace(errortrace bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) FilterPath(filterpaths ...string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) Human(human bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) Pretty(pretty bool) *Get { _ = "STUB: not implemented"; return nil }
