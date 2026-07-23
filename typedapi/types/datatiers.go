package types

type DataTiers struct {
	Available   bool                     `json:"available"`
	DataCold    DataTierPhaseStatistics  `json:"data_cold"`
	DataContent DataTierPhaseStatistics  `json:"data_content"`
	DataFrozen  *DataTierPhaseStatistics `json:"data_frozen,omitempty"`
	DataHot     DataTierPhaseStatistics  `json:"data_hot"`
	DataWarm    DataTierPhaseStatistics  `json:"data_warm"`
	Enabled     bool                     `json:"enabled"`
}

func (s *DataTiers) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDataTiers() *DataTiers { _ = "STUB: not implemented"; return nil }
