package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakerapi"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakerelementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakersimilarity"
)

type _amazonSageMakerServiceSettings struct {
	v *types.AmazonSageMakerServiceSettings
}

func NewAmazonSageMakerServiceSettings(accesskey string, api amazonsagemakerapi.AmazonSageMakerApi, endpointname string, region string, secretkey string) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) AccessKey(accesskey string) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) Api(api amazonsagemakerapi.AmazonSageMakerApi) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) BatchSize(batchsize int) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) Dimensions(dimensions int) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) ElementType(elementtype amazonsagemakerelementtype.AmazonSageMakerElementType) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) EndpointName(endpointname string) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) InferenceComponentName(inferencecomponentname string) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) Region(region string) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) SecretKey(secretkey string) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) Similarity(similarity amazonsagemakersimilarity.AmazonSageMakerSimilarity) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) TargetContainerHostname(targetcontainerhostname string) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) TargetModel(targetmodel string) *_amazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_amazonSageMakerServiceSettings) AmazonSageMakerServiceSettingsCaster() *types.AmazonSageMakerServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
