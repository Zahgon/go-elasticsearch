package validatedetector

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/excludefrequent"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ValidateDetector struct {
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

type NewValidateDetector func() *ValidateDetector

func NewValidateDetectorFunc(tp elastictransport.Interface) NewValidateDetector {
	_ = "STUB: not implemented"
	return *new(NewValidateDetector)
}

func New(tp elastictransport.Interface) *ValidateDetector { _ = "STUB: not implemented"; return nil }

func (r *ValidateDetector) Raw(raw io.Reader) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) Request(req *Request) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ValidateDetector) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ValidateDetector) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ValidateDetector) Header(key, value string) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) ErrorTrace(errortrace bool) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) FilterPath(filterpaths ...string) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) Human(human bool) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) Pretty(pretty bool) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) ByFieldName(field string) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) CustomRules(customrules ...types.DetectionRuleVariant) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) CustomRulesValues(customrulesvalues []types.DetectionRule) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) DetectorDescription(detectordescription string) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) DetectorIndex(detectorindex int) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) ExcludeFrequent(excludefrequent excludefrequent.ExcludeFrequent) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) FieldName(field string) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) Function(function string) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) OverFieldName(field string) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) PartitionFieldName(field string) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateDetector) UseNull(usenull bool) *ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}
