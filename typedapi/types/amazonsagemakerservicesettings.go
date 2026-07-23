package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakerapi"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakerelementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakersimilarity"
)

type AmazonSageMakerServiceSettings struct {
	AccessKey string `json:"access_key"`

	Api amazonsagemakerapi.AmazonSageMakerApi `json:"api"`

	BatchSize *int `json:"batch_size,omitempty"`

	Dimensions *int `json:"dimensions,omitempty"`

	ElementType *amazonsagemakerelementtype.AmazonSageMakerElementType `json:"element_type,omitempty"`

	EndpointName string `json:"endpoint_name"`

	InferenceComponentName *string `json:"inference_component_name,omitempty"`

	Region string `json:"region"`

	SecretKey string `json:"secret_key"`

	Similarity *amazonsagemakersimilarity.AmazonSageMakerSimilarity `json:"similarity,omitempty"`

	TargetContainerHostname *string `json:"target_container_hostname,omitempty"`

	TargetModel *string `json:"target_model,omitempty"`
}

func (s *AmazonSageMakerServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAmazonSageMakerServiceSettings() *AmazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

type AmazonSageMakerServiceSettingsVariant interface {
	AmazonSageMakerServiceSettingsCaster() *AmazonSageMakerServiceSettings
}

func (s *AmazonSageMakerServiceSettings) AmazonSageMakerServiceSettingsCaster() *AmazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
