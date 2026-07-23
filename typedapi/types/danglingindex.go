package types

type DanglingIndex struct {
	CreationDateMillis int64    `json:"creation_date_millis"`
	IndexName          string   `json:"index_name"`
	IndexUuid          string   `json:"index_uuid"`
	NodeIds            []string `json:"node_ids"`
}

func (s *DanglingIndex) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDanglingIndex() *DanglingIndex { _ = "STUB: not implemented"; return nil }
