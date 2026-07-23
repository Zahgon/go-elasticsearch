package updatedataframeanalytics

type Request struct {
	AllowLazyStart *bool `json:"allow_lazy_start,omitempty"`

	Description *string `json:"description,omitempty"`

	MaxNumThreads *int `json:"max_num_threads,omitempty"`

	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
