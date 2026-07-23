package geoipstats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GeoIpStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGeoIpStats func() *GeoIpStats

func NewGeoIpStatsFunc(tp elastictransport.Interface) NewGeoIpStats {
	_ = "STUB: not implemented"
	return *new(NewGeoIpStats)
}

func New(tp elastictransport.Interface) *GeoIpStats { _ = "STUB: not implemented"; return nil }

func (r *GeoIpStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GeoIpStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GeoIpStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GeoIpStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GeoIpStats) Header(key, value string) *GeoIpStats { _ = "STUB: not implemented"; return nil }

func (r *GeoIpStats) ErrorTrace(errortrace bool) *GeoIpStats { _ = "STUB: not implemented"; return nil }

func (r *GeoIpStats) FilterPath(filterpaths ...string) *GeoIpStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GeoIpStats) Human(human bool) *GeoIpStats { _ = "STUB: not implemented"; return nil }

func (r *GeoIpStats) Pretty(pretty bool) *GeoIpStats { _ = "STUB: not implemented"; return nil }
