package circuitbreaker

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catcircuitbreakercolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	circuitbreakerpatternsMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CircuitBreaker struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	circuitbreakerpatterns string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCircuitBreaker func() *CircuitBreaker

func NewCircuitBreakerFunc(tp elastictransport.Interface) NewCircuitBreaker {
	_ = "STUB: not implemented"
	return *new(NewCircuitBreaker)
}

func New(tp elastictransport.Interface) *CircuitBreaker { _ = "STUB: not implemented"; return nil }

func (r *CircuitBreaker) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CircuitBreaker) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CircuitBreaker) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r CircuitBreaker) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *CircuitBreaker) Header(key, value string) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) CircuitBreakerPatterns(circuitbreakerpatterns ...string) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) H(catcircuitbreakercolumns ...catcircuitbreakercolumn.CatCircuitBreakerColumn) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) S(names ...string) *CircuitBreaker { _ = "STUB: not implemented"; return nil }

func (r *CircuitBreaker) Local(local bool) *CircuitBreaker { _ = "STUB: not implemented"; return nil }

func (r *CircuitBreaker) MasterTimeout(duration string) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) Bytes(bytes bytes.Bytes) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) Format(format string) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) Help(help bool) *CircuitBreaker { _ = "STUB: not implemented"; return nil }

func (r *CircuitBreaker) Time(time timeunit.TimeUnit) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) V(v bool) *CircuitBreaker { _ = "STUB: not implemented"; return nil }

func (r *CircuitBreaker) ErrorTrace(errortrace bool) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) FilterPath(filterpaths ...string) *CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *CircuitBreaker) Human(human bool) *CircuitBreaker { _ = "STUB: not implemented"; return nil }

func (r *CircuitBreaker) Pretty(pretty bool) *CircuitBreaker { _ = "STUB: not implemented"; return nil }
