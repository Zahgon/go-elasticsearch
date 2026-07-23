package types

type DfsProfile struct {
	Knn        []DfsKnnProfile       `json:"knn,omitempty"`
	Statistics *DfsStatisticsProfile `json:"statistics,omitempty"`
}

func NewDfsProfile() *DfsProfile { _ = "STUB: not implemented"; return nil }
