package repositoryverifyintegrity

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

type RepositoryVerifyIntegrity struct {
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

type NewRepositoryVerifyIntegrity func(repository string) *RepositoryVerifyIntegrity

func NewRepositoryVerifyIntegrityFunc(tp elastictransport.Interface) NewRepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return *new(NewRepositoryVerifyIntegrity)
}

func New(tp elastictransport.Interface) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RepositoryVerifyIntegrity) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RepositoryVerifyIntegrity) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r RepositoryVerifyIntegrity) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *RepositoryVerifyIntegrity) Header(key, value string) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) _repository(repository string) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) BlobThreadPoolConcurrency(blobthreadpoolconcurrency int) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) IndexSnapshotVerificationConcurrency(indexsnapshotverificationconcurrency int) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) IndexVerificationConcurrency(indexverificationconcurrency int) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) MaxBytesPerSec(maxbytespersec string) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) MaxFailedShardSnapshots(maxfailedshardsnapshots int) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) MetaThreadPoolConcurrency(metathreadpoolconcurrency int) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) SnapshotVerificationConcurrency(snapshotverificationconcurrency int) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) VerifyBlobContents(verifyblobcontents bool) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) ErrorTrace(errortrace bool) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) FilterPath(filterpaths ...string) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) Human(human bool) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVerifyIntegrity) Pretty(pretty bool) *RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}
