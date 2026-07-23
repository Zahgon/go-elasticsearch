package types

type ElasticsearchError struct {
	ErrorCause ErrorCause `json:"error"`
	Status     int        `json:"status"`
}

func (e ElasticsearchError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ElasticsearchError) Is(err error) bool { _ = "STUB: not implemented"; return false }

func (e ElasticsearchError) As(err interface{}) bool { _ = "STUB: not implemented"; return false }

func NewElasticsearchError() *ElasticsearchError { _ = "STUB: not implemented"; return nil }

func (e *ElasticsearchError) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
