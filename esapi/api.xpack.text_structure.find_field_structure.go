package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTextStructureFindFieldStructureFunc(t Transport) TextStructureFindFieldStructure {
	_ = "STUB: not implemented"
	return *new(TextStructureFindFieldStructure)
}

type TextStructureFindFieldStructure func(field string, index string, o ...func(*TextStructureFindFieldStructureRequest)) (*Response, error)

type TextStructureFindFieldStructureRequest struct {
	ColumnNames       []string
	Delimiter         string
	DocumentsToSample *int
	EcsCompatibility  string
	Explain           *bool
	Field             string
	Format            string
	GrokPattern       string
	Index             string
	Quote             string
	ShouldTrimFields  *bool
	Timeout           time.Duration
	TimestampField    string
	TimestampFormat   string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TextStructureFindFieldStructureRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TextStructureFindFieldStructure) WithContext(v context.Context) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithColumnNames(v ...string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithDelimiter(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithDocumentsToSample(v int) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithEcsCompatibility(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithExplain(v bool) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithField(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithFormat(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithGrokPattern(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithIndex(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithQuote(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithShouldTrimFields(v bool) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithTimeout(v time.Duration) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithTimestampField(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithTimestampFormat(v string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithPretty() func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithHuman() func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithErrorTrace() func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithFilterPath(v ...string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithHeader(h map[string]string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindFieldStructure) WithOpaqueID(s string) func(*TextStructureFindFieldStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}
