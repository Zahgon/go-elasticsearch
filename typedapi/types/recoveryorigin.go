package types

type RecoveryOrigin struct {
	BootstrapNewHistoryUuid *bool   `json:"bootstrap_new_history_uuid,omitempty"`
	Host                    *string `json:"host,omitempty"`
	Hostname                *string `json:"hostname,omitempty"`
	Id                      *string `json:"id,omitempty"`
	Index                   *string `json:"index,omitempty"`
	Ip                      *string `json:"ip,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Repository              *string `json:"repository,omitempty"`
	RestoreUUID             *string `json:"restoreUUID,omitempty"`
	Snapshot                *string `json:"snapshot,omitempty"`
	TransportAddress        *string `json:"transport_address,omitempty"`
	Version                 *string `json:"version,omitempty"`
}

func (s *RecoveryOrigin) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRecoveryOrigin() *RecoveryOrigin { _ = "STUB: not implemented"; return nil }
