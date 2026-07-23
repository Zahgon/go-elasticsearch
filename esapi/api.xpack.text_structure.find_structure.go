package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newTextStructureFindStructureFunc(t Transport) TextStructureFindStructure {
	_ = "STUB: not implemented"
	return *new(TextStructureFindStructure)
}

type TextStructureFindStructure func(body io.Reader, o ...func(*TextStructureFindStructureRequest)) (*Response, error)

type TextStructureFindStructureRequest struct {
	Body io.Reader

	Charset            string
	ColumnNames        []string
	Delimiter          string
	EcsCompatibility   string
	Explain            *bool
	Format             string
	GrokPattern        string
	HasHeaderRow       *bool
	LineMergeSizeLimit *int
	LinesToSample      *int
	Quote              string
	ShouldTrimFields   *bool
	Timeout            time.Duration
	TimestampField     string
	TimestampFormat    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TextStructureFindStructureRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TextStructureFindStructure) WithContext(v context.Context) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithCharset(v string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithColumnNames(v ...string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithDelimiter(v string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithEcsCompatibility(v string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithExplain(v bool) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithFormat(v string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithGrokPattern(v string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithHasHeaderRow(v bool) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithLineMergeSizeLimit(v int) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithLinesToSample(v int) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithQuote(v string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithShouldTrimFields(v bool) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithTimeout(v time.Duration) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithTimestampField(v string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithTimestampFormat(v string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithPretty() func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithHuman() func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithErrorTrace() func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithFilterPath(v ...string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithHeader(h map[string]string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureFindStructure) WithOpaqueID(s string) func(*TextStructureFindStructureRequest) {
	_ = "STUB: not implemented"
	return nil
}
