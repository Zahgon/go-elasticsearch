package elasticsearch

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/elastic/go-elasticsearch/v9/typedapi"

	"github.com/elastic/go-elasticsearch/v9/esapi"
	"github.com/elastic/go-elasticsearch/v9/internal/version"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	defaultURL = "http://localhost:9200"

	Version        = version.Client
	unknownProduct = "the client noticed that the server is not Elasticsearch and we do not support this unknown product"

	HeaderClientMeta = "x-elastic-client-meta"

	compatibilityHeader = "application/vnd.elasticsearch+json;compatible-with=9"

	typedClientMetaFlag = "hl=1"
)

var (
	esCompatHeader = "ELASTIC_CLIENT_APIVERSIONING"
	userAgent      string
	reGoVersion    = regexp.MustCompile(`go(\d+\.\d+\..+)`)
	reMetaVersion  = regexp.MustCompile("([0-9.]+)(.*)")
)

var (
	ErrClosed = errors.New("client is closed")

	ErrAlreadyClosed = errors.New("client is already closed")
)

func init() {
	userAgent = initUserAgent()
}

type Config struct {
	Addresses []string
	Username  string
	Password  string

	CloudID                string
	APIKey                 string
	ServiceToken           string
	CertificateFingerprint string

	Header http.Header

	CACert []byte

	RetryOnStatus []int
	DisableRetry  bool
	MaxRetries    int
	RetryOnError  func(*http.Request, error) bool

	CompressRequestBody      bool
	CompressRequestBodyLevel int
	PoolCompressor           bool

	DiscoverNodesOnStart  bool
	DiscoverNodesInterval time.Duration

	EnableMetrics           bool
	EnableDebugLogger       bool
	EnableCompatibilityMode bool

	DisableMetaHeader bool

	AutoDrainBody bool

	RetryBackoff func(attempt int) time.Duration

	Transport http.RoundTripper
	Logger    elastictransport.Logger
	Selector  elastictransport.Selector

	ConnectionPoolFunc func([]*elastictransport.Connection, elastictransport.Selector) elastictransport.ConnectionPool

	Instrumentation elastictransport.Instrumentation

	Interceptors []elastictransport.InterceptorFunc
}

func NewOpenTelemetryInstrumentation(provider trace.TracerProvider, captureSearchBody bool) elastictransport.Instrumentation {
	_ = "STUB: not implemented"
	return *new(elastictransport.Instrumentation)
}

type BaseClient struct {
	Transport           elastictransport.Interface
	metaHeader          string
	compatibilityHeader bool

	autoDrainBody       bool
	disableMetaHeader   bool
	productCheckMu      sync.RWMutex
	productCheckSuccess bool

	closeDone uint32
}

type Client struct {
	BaseClient
	*esapi.API
}

type TypedClient struct {
	BaseClient
	*typedapi.MethodAPI
}

func NewBaseClient(cfg Config) (*BaseClient, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDefaultClient() (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

func NewClient(cfg Config) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTypedClient(cfg Config) (*TypedClient, error) { _ = "STUB: not implemented"; return nil, nil }

func New(opts ...Option) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

func NewBase(opts ...Option) (*BaseClient, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTyped(opts ...Option) (*TypedClient, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTypedFrom(c *Client) *TypedClient { _ = "STUB: not implemented"; return nil }

func newTransport(cfg Config) (*elastictransport.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *BaseClient) Perform(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *BaseClient) InstrumentationEnabled() elastictransport.Instrumentation {
	_ = "STUB: not implemented"
	return *new(elastictransport.Instrumentation)
}

func (c *BaseClient) doProductCheck(f func() error) error { _ = "STUB: not implemented"; return nil }

func genuineCheckHeader(header http.Header) error { _ = "STUB: not implemented"; return nil }

func (c *BaseClient) Metrics() (elastictransport.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(elastictransport.Metrics), nil
}

func (c *BaseClient) DiscoverNodes() error { _ = "STUB: not implemented"; return nil }

func (c *BaseClient) isClosed() bool { _ = "STUB: not implemented"; return false }

func (c *BaseClient) Close(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func addrsFromEnvironment() []string { _ = "STUB: not implemented"; return nil }

func parseDefaultURL() (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

func addrsToURLs(addrs []string) ([]*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

func addrFromCloudID(input string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func initUserAgent() string { _ = "STUB: not implemented"; return "" }

func initMetaHeader(transport interface{}) string { _ = "STUB: not implemented"; return "" }

func buildStrippedVersion(version string) string { _ = "STUB: not implemented"; return "" }

const autoDrainingMaxBytes = 512 * 1024

type autoDrainingReader struct {
	io.ReadCloser
}

func (a *autoDrainingReader) Close() error { _ = "STUB: not implemented"; return nil }
