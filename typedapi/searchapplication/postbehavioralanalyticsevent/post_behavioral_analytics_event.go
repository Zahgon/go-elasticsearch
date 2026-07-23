package postbehavioralanalyticsevent

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	collectionnameMask = iota + 1

	eventtypeMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostBehavioralAnalyticsEvent struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      any
	deferred []func(request any) error
	buf      *gobytes.Buffer

	paramSet int

	collectionname string
	eventtype      string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPostBehavioralAnalyticsEvent func(collectionname, eventtype string) *PostBehavioralAnalyticsEvent

func NewPostBehavioralAnalyticsEventFunc(tp elastictransport.Interface) NewPostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return *new(NewPostBehavioralAnalyticsEvent)
}

func New(tp elastictransport.Interface) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) Raw(raw io.Reader) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) Request(req any) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) Payload(payload any) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostBehavioralAnalyticsEvent) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostBehavioralAnalyticsEvent) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PostBehavioralAnalyticsEvent) Header(key, value string) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) _collectionname(collectionname string) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) _eventtype(eventtype string) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) Debug(debug bool) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) ErrorTrace(errortrace bool) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) FilterPath(filterpaths ...string) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) Human(human bool) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostBehavioralAnalyticsEvent) Pretty(pretty bool) *PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}
