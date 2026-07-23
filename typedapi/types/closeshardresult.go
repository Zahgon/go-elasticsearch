package types

type CloseShardResult struct {
	Failures []ShardFailure `json:"failures"`
}

func NewCloseShardResult() *CloseShardResult { _ = "STUB: not implemented"; return nil }
