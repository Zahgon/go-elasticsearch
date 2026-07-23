package puttrainedmodel

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainedmodeltype"
)

type Response struct {
	CompressedDefinition *string `json:"compressed_definition,omitempty"`

	CreateTime types.DateTime `json:"create_time,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	DefaultFieldMap map[string]string `json:"default_field_map,omitempty"`

	Description *string `json:"description,omitempty"`

	EstimatedHeapMemoryUsageBytes *int `json:"estimated_heap_memory_usage_bytes,omitempty"`

	EstimatedOperations *int `json:"estimated_operations,omitempty"`

	FullyDefined *bool `json:"fully_defined,omitempty"`

	InferenceConfig *types.InferenceConfigCreateContainer `json:"inference_config,omitempty"`

	Input types.TrainedModelConfigInput `json:"input"`

	LicenseLevel *string                     `json:"license_level,omitempty"`
	Location     *types.TrainedModelLocation `json:"location,omitempty"`

	Metadata *types.TrainedModelConfigMetadata `json:"metadata,omitempty"`

	ModelId        string                    `json:"model_id"`
	ModelPackage   *types.ModelPackageConfig `json:"model_package,omitempty"`
	ModelSizeBytes types.ByteSize            `json:"model_size_bytes,omitempty"`

	ModelType            *trainedmodeltype.TrainedModelType `json:"model_type,omitempty"`
	PlatformArchitecture *string                            `json:"platform_architecture,omitempty"`
	PrefixStrings        *types.TrainedModelPrefixStrings   `json:"prefix_strings,omitempty"`

	Tags []string `json:"tags"`

	Version *string `json:"version,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
