package puttrainedmodelvocabulary

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
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutTrainedModelVocabulary struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutTrainedModelVocabulary func(modelid string) *PutTrainedModelVocabulary

func NewPutTrainedModelVocabularyFunc(tp elastictransport.Interface) NewPutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return *new(NewPutTrainedModelVocabulary)
}

func New(tp elastictransport.Interface) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) Raw(raw io.Reader) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) Request(req *Request) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModelVocabulary) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModelVocabulary) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutTrainedModelVocabulary) Header(key, value string) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) _modelid(modelid string) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) ErrorTrace(errortrace bool) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) FilterPath(filterpaths ...string) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) Human(human bool) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) Pretty(pretty bool) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) Merges(merges ...string) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) Scores(scores ...types.Float64) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelVocabulary) Vocabulary(vocabularies ...string) *PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}
