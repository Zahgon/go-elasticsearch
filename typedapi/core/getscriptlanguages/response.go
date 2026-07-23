package getscriptlanguages

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	LanguageContexts []types.LanguageContext `json:"language_contexts"`
	TypesAllowed     []string                `json:"types_allowed"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
