package putwatch

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

type PutWatch struct {
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

type NewPutWatch func(id string) *PutWatch

func NewPutWatchFunc(tp elastictransport.Interface) NewPutWatch {
	_ = "STUB: not implemented"
	return *new(NewPutWatch)
}

func New(tp elastictransport.Interface) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) Raw(raw io.Reader) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) Request(req *Request) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutWatch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutWatch) Header(key, value string) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) _id(id string) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) Active(active bool) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) IfPrimaryTerm(ifprimaryterm string) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) IfSeqNo(sequencenumber string) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) Version(versionnumber string) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) ErrorTrace(errortrace bool) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) FilterPath(filterpaths ...string) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) Human(human bool) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) Pretty(pretty bool) *PutWatch { _ = "STUB: not implemented"; return nil }

func (r *PutWatch) Actions(actions map[string]types.WatcherAction) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) AddAction(key string, value types.WatcherActionVariant) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) Condition(condition types.WatcherConditionVariant) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) Input(input types.WatcherInputVariant) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) Metadata(metadata types.MetadataVariant) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) ThrottlePeriod(duration types.DurationVariant) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) ThrottlePeriodInMillis(durationvalueunitmillis int64) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) Transform(transform types.TransformContainerVariant) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatch) Trigger(trigger types.TriggerContainerVariant) *PutWatch {
	_ = "STUB: not implemented"
	return nil
}
