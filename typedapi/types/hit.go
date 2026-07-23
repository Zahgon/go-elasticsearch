package types

import (
	"encoding/json"
)

type Hit struct {
	Explanation_       *Explanation                 `json:"_explanation,omitempty"`
	Fields             map[string]json.RawMessage   `json:"fields,omitempty"`
	Highlight          map[string][]string          `json:"highlight,omitempty"`
	Id_                *string                      `json:"_id,omitempty"`
	IgnoredFieldValues map[string][]json.RawMessage `json:"ignored_field_values,omitempty"`
	Ignored_           []string                     `json:"_ignored,omitempty"`
	Index_             string                       `json:"_index"`
	InnerHits          map[string]InnerHitsResult   `json:"inner_hits,omitempty"`
	MatchedQueries     any                          `json:"matched_queries,omitempty"`
	Nested_            *NestedIdentity              `json:"_nested,omitempty"`
	Node_              *string                      `json:"_node,omitempty"`
	PrimaryTerm_       *int64                       `json:"_primary_term,omitempty"`
	Rank_              *int                         `json:"_rank,omitempty"`
	Routing_           *string                      `json:"_routing,omitempty"`
	Score_             *Float64                     `json:"_score,omitempty"`
	SeqNo_             *int64                       `json:"_seq_no,omitempty"`
	Shard_             *string                      `json:"_shard,omitempty"`
	Sort               []FieldValue                 `json:"sort,omitempty"`
	Source_            json.RawMessage              `json:"_source,omitempty"`
	Version_           *int64                       `json:"_version,omitempty"`
}

func (s *Hit) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHit() *Hit { _ = "STUB: not implemented"; return nil }
