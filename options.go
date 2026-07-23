package elasticsearch

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

type Option struct {
	apply func(*clientOptions) error
}

type clientOptions struct {
	addresses            []string
	cloudID              string
	discoverNodesOnStart bool
	compatibilityMode    bool
	disableMetaHeader    bool
	autoDrainBody        bool
	transportOptions     []elastictransport.Option
}

func WithAddresses(addrs ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCloudID(cloudID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDiscoverNodesOnStart() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCompatibilityMode() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDisableMetaHeader() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAutoDrainBody() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAPIKey(apiKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBasicAuth(username, password string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithServiceToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCACert(cert []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCertificateFingerprint(fingerprint string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRetry(maxRetries int, onStatus ...int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCompression(level ...int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithInstrumentation(i elastictransport.Instrumentation) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLogger(l elastictransport.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTransportOptions(opts ...elastictransport.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func withTransportOption(to elastictransport.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type resolvedOptions struct {
	transport            *elastictransport.Client
	metaHeader           string
	disableMetaHeader    bool
	compatibilityHeader  bool
	discoverNodesOnStart bool
	autoDrainBody        bool
}

func resolveOptions(opts []Option, metaHeaderSuffix string) (*resolvedOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
