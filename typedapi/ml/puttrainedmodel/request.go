package puttrainedmodel

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainedmodeltype"
)

type Request struct {
	CompressedDefinition *string `json:"compressed_definition,omitempty"`

	Definition *types.Definition `json:"definition,omitempty"`

	Description *string `json:"description,omitempty"`

	InferenceConfig *types.InferenceConfigCreateContainer `json:"inference_config,omitempty"`

	Input *types.Input `json:"input,omitempty"`

	Metadata json.RawMessage `json:"metadata,omitempty"`

	ModelSizeBytes *int64 `json:"model_size_bytes,omitempty"`

	ModelType *trainedmodeltype.TrainedModelType `json:"model_type,omitempty"`

	PlatformArchitecture *string `json:"platform_architecture,omitempty"`

	PrefixStrings *types.TrainedModelPrefixStrings `json:"prefix_strings,omitempty"`

	Tags []string `json:"tags,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
