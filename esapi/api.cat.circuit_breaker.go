package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatCircuitBreakerFunc(t Transport) CatCircuitBreaker {
	_ = "STUB: not implemented"
	return *new(CatCircuitBreaker)
}

type CatCircuitBreaker func(o ...func(*CatCircuitBreakerRequest)) (*Response, error)

type CatCircuitBreakerRequest struct {
	CircuitBreakerPatterns []string

	Bytes         string
	Format        string
	H             []string
	Help          *bool
	Local         *bool
	MasterTimeout time.Duration
	S             []string
	Time          string
	V             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatCircuitBreakerRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatCircuitBreaker) WithContext(v context.Context) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithCircuitBreakerPatterns(v ...string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithBytes(v string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithFormat(v string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithH(v ...string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithHelp(v bool) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithLocal(v bool) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithMasterTimeout(v time.Duration) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithS(v ...string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithTime(v string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithV(v bool) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithPretty() func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithHuman() func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithErrorTrace() func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithFilterPath(v ...string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithHeader(h map[string]string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatCircuitBreaker) WithOpaqueID(s string) func(*CatCircuitBreakerRequest) {
	_ = "STUB: not implemented"
	return nil
}
