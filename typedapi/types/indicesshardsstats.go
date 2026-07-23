package types

type IndicesShardsStats struct {
	AllFields FieldSummary            `json:"all_fields"`
	Fields    map[string]FieldSummary `json:"fields"`
}

func NewIndicesShardsStats() *IndicesShardsStats { _ = "STUB: not implemented"; return nil }
