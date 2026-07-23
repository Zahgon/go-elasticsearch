package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoverystage"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoverytype"
)

type ShardRecovery struct {
	Id      int64               `json:"id"`
	Index   RecoveryIndexStatus `json:"index"`
	Primary bool                `json:"primary"`
	Source  RecoveryOrigin      `json:"source"`

	Stage             recoverystage.RecoveryStage `json:"stage"`
	Start             *RecoveryStartStatus        `json:"start,omitempty"`
	StartTime         DateTime                    `json:"start_time,omitempty"`
	StartTimeInMillis int64                       `json:"start_time_in_millis"`
	StopTime          DateTime                    `json:"stop_time,omitempty"`
	StopTimeInMillis  *int64                      `json:"stop_time_in_millis,omitempty"`
	Target            RecoveryOrigin              `json:"target"`
	TotalTime         Duration                    `json:"total_time,omitempty"`
	TotalTimeInMillis int64                       `json:"total_time_in_millis"`
	Translog          TranslogStatus              `json:"translog"`

	Type        recoverytype.RecoveryType `json:"type"`
	VerifyIndex VerifyIndex               `json:"verify_index"`
}

func (s *ShardRecovery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardRecovery() *ShardRecovery { _ = "STUB: not implemented"; return nil }
