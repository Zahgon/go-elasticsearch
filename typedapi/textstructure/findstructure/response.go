package findstructure

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Charset string `json:"charset"`

	ColumnNames         []string `json:"column_names,omitempty"`
	Delimiter           *string  `json:"delimiter,omitempty"`
	ExcludeLinesPattern *string  `json:"exclude_lines_pattern,omitempty"`
	Explanation         []string `json:"explanation,omitempty"`

	FieldStats map[string]types.FieldStat `json:"field_stats"`

	Format      string  `json:"format"`
	GrokPattern *string `json:"grok_pattern,omitempty"`

	HasByteOrderMarker bool                 `json:"has_byte_order_marker"`
	HasHeaderRow       *bool                `json:"has_header_row,omitempty"`
	IngestPipeline     types.PipelineConfig `json:"ingest_pipeline"`

	JavaTimestampFormats []string `json:"java_timestamp_formats,omitempty"`

	JodaTimestampFormats []string `json:"joda_timestamp_formats,omitempty"`

	Mappings              types.TypeMapping `json:"mappings"`
	MultilineStartPattern *string           `json:"multiline_start_pattern,omitempty"`

	NeedClientTimezone bool `json:"need_client_timezone"`

	NumLinesAnalyzed int `json:"num_lines_analyzed"`

	NumMessagesAnalyzed int     `json:"num_messages_analyzed"`
	Quote               *string `json:"quote,omitempty"`

	SampleStart      string `json:"sample_start"`
	ShouldTrimFields *bool  `json:"should_trim_fields,omitempty"`

	TimestampField *string `json:"timestamp_field,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
