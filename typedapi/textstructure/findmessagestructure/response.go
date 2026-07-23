package findmessagestructure

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ecscompatibilitytype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/formattype"
)

type Response struct {
	Charset               string                                     `json:"charset"`
	EcsCompatibility      *ecscompatibilitytype.EcsCompatibilityType `json:"ecs_compatibility,omitempty"`
	FieldStats            map[string]types.FieldStat                 `json:"field_stats"`
	Format                formattype.FormatType                      `json:"format"`
	GrokPattern           *string                                    `json:"grok_pattern,omitempty"`
	IngestPipeline        types.PipelineConfig                       `json:"ingest_pipeline"`
	JavaTimestampFormats  []string                                   `json:"java_timestamp_formats,omitempty"`
	JodaTimestampFormats  []string                                   `json:"joda_timestamp_formats,omitempty"`
	Mappings              types.TypeMapping                          `json:"mappings"`
	MultilineStartPattern *string                                    `json:"multiline_start_pattern,omitempty"`
	NeedClientTimezone    bool                                       `json:"need_client_timezone"`
	NumLinesAnalyzed      int                                        `json:"num_lines_analyzed"`
	NumMessagesAnalyzed   int                                        `json:"num_messages_analyzed"`
	SampleStart           string                                     `json:"sample_start"`
	TimestampField        *string                                    `json:"timestamp_field,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
