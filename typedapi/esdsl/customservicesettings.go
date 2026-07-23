package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/customserviceinputtype"
)

type _customServiceSettings struct {
	v *types.CustomServiceSettings
}

func NewCustomServiceSettings(request types.CustomRequestParamsVariant, response types.CustomResponseParamsVariant) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) BatchSize(batchsize int) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) Headers(headers map[string]string) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) AddHeader(key string, value string) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) InputType(inputtype map[customserviceinputtype.CustomServiceInputType]string) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) AddInputType(key customserviceinputtype.CustomServiceInputType, value string) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) QueryParameters(queryparameters ...[]string) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) Request(request types.CustomRequestParamsVariant) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) Response(response types.CustomResponseParamsVariant) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) SecretParameters(secretparameters map[string]string) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) AddSecretParameter(key string, value string) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) Url(url string) *_customServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customServiceSettings) CustomServiceSettingsCaster() *types.CustomServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
