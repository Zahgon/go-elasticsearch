package scriptspainlessexecute

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/painlesscontext"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ScriptsPainlessExecute struct {
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

type NewScriptsPainlessExecute func() *ScriptsPainlessExecute

func NewScriptsPainlessExecuteFunc(tp elastictransport.Interface) NewScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return *new(NewScriptsPainlessExecute)
}

func New(tp elastictransport.Interface) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) Raw(raw io.Reader) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) Request(req *Request) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ScriptsPainlessExecute) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ScriptsPainlessExecute) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ScriptsPainlessExecute) Header(key, value string) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) ErrorTrace(errortrace bool) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) FilterPath(filterpaths ...string) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) Human(human bool) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) Pretty(pretty bool) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) Context(context painlesscontext.PainlessContext) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) ContextSetup(contextsetup types.PainlessContextSetupVariant) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScriptsPainlessExecute) Script(script types.ScriptVariant) *ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}
