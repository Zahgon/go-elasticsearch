package previewtransform

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	GeneratedDestIndex types.IndexState  `json:"generated_dest_index"`
	Preview            []json.RawMessage `json:"preview"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
