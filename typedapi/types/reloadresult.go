package types

type ReloadResult struct {
	ReloadDetails []ReloadDetails `json:"reload_details"`
	Shards_       ShardStatistics `json:"_shards"`
}

func NewReloadResult() *ReloadResult { _ = "STUB: not implemented"; return nil }
