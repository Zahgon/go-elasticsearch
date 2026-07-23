package removepolicy

type Response struct {
	FailedIndexes []string `json:"failed_indexes"`
	HasFailures   bool     `json:"has_failures"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
