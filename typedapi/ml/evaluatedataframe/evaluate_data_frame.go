package evaluatedataframe

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

type EvaluateDataFrame struct {
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

type NewEvaluateDataFrame func() *EvaluateDataFrame

func NewEvaluateDataFrameFunc(tp elastictransport.Interface) NewEvaluateDataFrame {
	_ = "STUB: not implemented"
	return *new(NewEvaluateDataFrame)
}

func New(tp elastictransport.Interface) *EvaluateDataFrame { _ = "STUB: not implemented"; return nil }

func (r *EvaluateDataFrame) Raw(raw io.Reader) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) Request(req *Request) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EvaluateDataFrame) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EvaluateDataFrame) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *EvaluateDataFrame) Header(key, value string) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) ErrorTrace(errortrace bool) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) FilterPath(filterpaths ...string) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) Human(human bool) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) Pretty(pretty bool) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) Evaluation(evaluation types.DataframeEvaluationContainerVariant) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) Index(indexname string) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (r *EvaluateDataFrame) Query(query types.QueryVariant) *EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}
