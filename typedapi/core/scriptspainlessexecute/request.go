package scriptspainlessexecute

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/painlesscontext"
)

type Request struct {
	Context *painlesscontext.PainlessContext `json:"context,omitempty"`

	ContextSetup *types.PainlessContextSetup `json:"context_setup,omitempty"`

	Script *types.Script `json:"script,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
