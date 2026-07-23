package elasticsearchservicetype

type ElasticsearchServiceType struct {
	Name string
}

var (
	Elasticsearch = ElasticsearchServiceType{"elasticsearch"}
)

func (e ElasticsearchServiceType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ElasticsearchServiceType) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e ElasticsearchServiceType) String() string { _ = "STUB: not implemented"; return "" }
