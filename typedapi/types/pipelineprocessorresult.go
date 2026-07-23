package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/pipelinesimulationstatusoptions"
)

type PipelineProcessorResult struct {
	Description   *string                                                          `json:"description,omitempty"`
	Doc           *DocumentSimulation                                              `json:"doc,omitempty"`
	Error         *ErrorCause                                                      `json:"error,omitempty"`
	IgnoredError  *ErrorCause                                                      `json:"ignored_error,omitempty"`
	ProcessorType *string                                                          `json:"processor_type,omitempty"`
	Status        *pipelinesimulationstatusoptions.PipelineSimulationStatusOptions `json:"status,omitempty"`
	Tag           *string                                                          `json:"tag,omitempty"`
}

func (s *PipelineProcessorResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPipelineProcessorResult() *PipelineProcessorResult { _ = "STUB: not implemented"; return nil }
