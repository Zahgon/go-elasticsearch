package putcustom

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tasktypecustom"
)

type Response struct {
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`

	InferenceId string `json:"inference_id"`

	Service string `json:"service"`

	ServiceSettings json.RawMessage `json:"service_settings"`

	TaskSettings json.RawMessage `json:"task_settings,omitempty"`

	TaskType tasktypecustom.TaskTypeCustom `json:"task_type"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
