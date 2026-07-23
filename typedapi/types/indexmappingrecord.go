package types

type IndexMappingRecord struct {
	Item     *TypeMapping `json:"item,omitempty"`
	Mappings TypeMapping  `json:"mappings"`
}

func NewIndexMappingRecord() *IndexMappingRecord { _ = "STUB: not implemented"; return nil }
