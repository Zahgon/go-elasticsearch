package esapi

import (
	"context"
	"net/http"
)

func newCatMLTrainedModelsFunc(t Transport) CatMLTrainedModels {
	_ = "STUB: not implemented"
	return *new(CatMLTrainedModels)
}

type CatMLTrainedModels func(o ...func(*CatMLTrainedModelsRequest)) (*Response, error)

type CatMLTrainedModelsRequest struct {
	ModelID string

	AllowNoMatch *bool
	Bytes        string
	Format       string
	From         *int
	H            []string
	Help         *bool
	S            []string
	Size         *int
	Time         string
	V            *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatMLTrainedModelsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatMLTrainedModels) WithContext(v context.Context) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithModelID(v string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithAllowNoMatch(v bool) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithBytes(v string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithFormat(v string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithFrom(v int) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithH(v ...string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithHelp(v bool) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithS(v ...string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithSize(v int) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithTime(v string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithV(v bool) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithPretty() func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithHuman() func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithErrorTrace() func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithFilterPath(v ...string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithHeader(h map[string]string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLTrainedModels) WithOpaqueID(s string) func(*CatMLTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}
