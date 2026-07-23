package migratetodatatiers

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MigrateToDataTiers struct {
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

type NewMigrateToDataTiers func() *MigrateToDataTiers

func NewMigrateToDataTiersFunc(tp elastictransport.Interface) NewMigrateToDataTiers {
	_ = "STUB: not implemented"
	return *new(NewMigrateToDataTiers)
}

func New(tp elastictransport.Interface) *MigrateToDataTiers { _ = "STUB: not implemented"; return nil }

func (r *MigrateToDataTiers) Raw(raw io.Reader) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) Request(req *Request) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MigrateToDataTiers) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MigrateToDataTiers) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *MigrateToDataTiers) Header(key, value string) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) DryRun(dryrun bool) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) MasterTimeout(duration string) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) ErrorTrace(errortrace bool) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) FilterPath(filterpaths ...string) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) Human(human bool) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) Pretty(pretty bool) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) LegacyTemplateToDelete(legacytemplatetodelete string) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (r *MigrateToDataTiers) NodeAttribute(nodeattribute string) *MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}
