package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newTextStructureFindMessageStructureFunc(t Transport) TextStructureFindMessageStructure {
	_ = "STUB: not implemented"
	return *new(TextStructureFindMessageStructure)
}

type TextStructureFindMessageStructure func(body io.Reader, o ...func(*TextStructureFindMessageStructureRequest)) (*Response, error)

type TextStructureFindMessageStructureRequest struct {
	Body io.Reader

	ColumnNames      []string
	Delimiter        string
	EcsCompatibility string
	Explain          *bool
	Format           string
	GrokPattern      string
	Quote            string
	ShouldTrimFields *bool
	Timeout          time.Duration
	TimestampField   string
	TimestampFormat  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TextStructureFindMessageStructureRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TextStructureFindMessageStructure) WithContext(v context.Context) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithColumnNames(v ...string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithDelimiter(v string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithEcsCompatibility(v string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithExplain(v bool) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithFormat(v string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithGrokPattern(v string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithQuote(v string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithShouldTrimFields(v bool) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithTimeout(v time.Duration) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithTimestampField(v string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithTimestampFormat(v string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithPretty() func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithHuman() func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithErrorTrace() func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithFilterPath(v ...string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithHeader(h map[string]string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindMessageStructure) WithOpaqueID(s string) func(*TextStructureFindMessageStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}
