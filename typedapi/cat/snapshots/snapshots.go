package snapshots

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catsnapshotscolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	repositoryMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Snapshots struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	repository string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSnapshots func() *Snapshots

func NewSnapshotsFunc(tp elastictransport.Interface) NewSnapshots {
	_ = "STUB: not implemented"
	return *new(NewSnapshots)
}

func New(tp elastictransport.Interface) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Snapshots) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Snapshots) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Snapshots) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Snapshots) Header(key, value string) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) Repository(repository string) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) IgnoreUnavailable(ignoreunavailable bool) *Snapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *Snapshots) H(catsnapshotscolumns ...catsnapshotscolumn.CatSnapshotsColumn) *Snapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *Snapshots) S(names ...string) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) MasterTimeout(duration string) *Snapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *Snapshots) Bytes(bytes bytes.Bytes) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) Format(format string) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) Help(help bool) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) Time(time timeunit.TimeUnit) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) V(v bool) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) ErrorTrace(errortrace bool) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) FilterPath(filterpaths ...string) *Snapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *Snapshots) Human(human bool) *Snapshots { _ = "STUB: not implemented"; return nil }

func (r *Snapshots) Pretty(pretty bool) *Snapshots { _ = "STUB: not implemented"; return nil }
