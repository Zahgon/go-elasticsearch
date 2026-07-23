package repositoryanalyze

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	repositoryMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RepositoryAnalyze struct {
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

type NewRepositoryAnalyze func(repository string) *RepositoryAnalyze

func NewRepositoryAnalyzeFunc(tp elastictransport.Interface) NewRepositoryAnalyze {
	_ = "STUB: not implemented"
	return *new(NewRepositoryAnalyze)
}

func New(tp elastictransport.Interface) *RepositoryAnalyze { _ = "STUB: not implemented"; return nil }

func (r *RepositoryAnalyze) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RepositoryAnalyze) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RepositoryAnalyze) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RepositoryAnalyze) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *RepositoryAnalyze) Header(key, value string) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) _repository(repository string) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) BlobCount(blobcount int) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) CheckOverwriteProtection(checkoverwriteprotection bool) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) Concurrency(concurrency int) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) Detailed(detailed bool) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) EarlyReadNodeCount(earlyreadnodecount int) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) MaxBlobSize(bytesize string) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) MaxTotalDataSize(bytesize string) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) RareActionProbability(rareactionprobability string) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) RarelyAbortWrites(rarelyabortwrites bool) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) ReadNodeCount(readnodecount int) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) RegisterOperationCount(registeroperationcount int) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) Seed(seed int) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) Timeout(duration string) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) ErrorTrace(errortrace bool) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) FilterPath(filterpaths ...string) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) Human(human bool) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryAnalyze) Pretty(pretty bool) *RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}
