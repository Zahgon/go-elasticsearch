package types

type MatrixStatsFields struct {
	Correlation map[string]Float64 `json:"correlation"`
	Count       int64              `json:"count"`
	Covariance  map[string]Float64 `json:"covariance"`
	Kurtosis    Float64            `json:"kurtosis"`
	Mean        Float64            `json:"mean"`
	Name        string             `json:"name"`
	Skewness    Float64            `json:"skewness"`
	Variance    Float64            `json:"variance"`
}

func (s *MatrixStatsFields) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMatrixStatsFields() *MatrixStatsFields { _ = "STUB: not implemented"; return nil }
