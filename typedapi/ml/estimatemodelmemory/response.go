package estimatemodelmemory

type Response struct {
	ModelMemoryEstimate string `json:"model_memory_estimate"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
