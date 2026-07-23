package types

type SimulateDocumentResult struct {
	Doc              *DocumentSimulation       `json:"doc,omitempty"`
	Error            *ErrorCause               `json:"error,omitempty"`
	ProcessorResults []PipelineProcessorResult `json:"processor_results,omitempty"`
}

func NewSimulateDocumentResult() *SimulateDocumentResult { _ = "STUB: not implemented"; return nil }
