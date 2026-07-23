package samlserviceprovidermetadata

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	realmnameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SamlServiceProviderMetadata struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	realmname string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSamlServiceProviderMetadata func(realmname string) *SamlServiceProviderMetadata

func NewSamlServiceProviderMetadataFunc(tp elastictransport.Interface) NewSamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return *new(NewSamlServiceProviderMetadata)
}

func New(tp elastictransport.Interface) *SamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlServiceProviderMetadata) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlServiceProviderMetadata) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlServiceProviderMetadata) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SamlServiceProviderMetadata) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SamlServiceProviderMetadata) Header(key, value string) *SamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlServiceProviderMetadata) _realmname(realmname string) *SamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlServiceProviderMetadata) ErrorTrace(errortrace bool) *SamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlServiceProviderMetadata) FilterPath(filterpaths ...string) *SamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlServiceProviderMetadata) Human(human bool) *SamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (r *SamlServiceProviderMetadata) Pretty(pretty bool) *SamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return nil
}
