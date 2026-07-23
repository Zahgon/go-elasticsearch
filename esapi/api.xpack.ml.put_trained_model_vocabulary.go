package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPutTrainedModelVocabularyFunc(t Transport) MLPutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return *new(MLPutTrainedModelVocabulary)
}

type MLPutTrainedModelVocabulary func(body io.Reader, model_id string, o ...func(*MLPutTrainedModelVocabularyRequest)) (*Response, error)

type MLPutTrainedModelVocabularyRequest struct {
	Body io.Reader

	ModelID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPutTrainedModelVocabularyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutTrainedModelVocabulary) WithContext(v context.Context) func(*MLPutTrainedModelVocabularyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelVocabulary) WithPretty() func(*MLPutTrainedModelVocabularyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelVocabulary) WithHuman() func(*MLPutTrainedModelVocabularyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelVocabulary) WithErrorTrace() func(*MLPutTrainedModelVocabularyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelVocabulary) WithFilterPath(v ...string) func(*MLPutTrainedModelVocabularyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelVocabulary) WithHeader(h map[string]string) func(*MLPutTrainedModelVocabularyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelVocabulary) WithOpaqueID(s string) func(*MLPutTrainedModelVocabularyRequest) {
	_ = "STUB: not implemented"
	return nil
}
