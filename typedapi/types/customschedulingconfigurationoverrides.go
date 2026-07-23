package types

type CustomSchedulingConfigurationOverrides struct {
	DomainAllowlist          []string `json:"domain_allowlist,omitempty"`
	MaxCrawlDepth            *int     `json:"max_crawl_depth,omitempty"`
	SeedUrls                 []string `json:"seed_urls,omitempty"`
	SitemapDiscoveryDisabled *bool    `json:"sitemap_discovery_disabled,omitempty"`
	SitemapUrls              []string `json:"sitemap_urls,omitempty"`
}

func (s *CustomSchedulingConfigurationOverrides) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCustomSchedulingConfigurationOverrides() *CustomSchedulingConfigurationOverrides {
	_ = "STUB: not implemented"
	return nil
}
