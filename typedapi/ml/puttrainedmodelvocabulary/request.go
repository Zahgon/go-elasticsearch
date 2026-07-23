package puttrainedmodelvocabulary

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Merges []string `json:"merges,omitempty"`

	Scores []types.Float64 `json:"scores,omitempty"`

	Vocabulary []string `json:"vocabulary"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
