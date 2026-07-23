package types

type PressureMemory struct {
	All ByteSize `json:"all,omitempty"`

	AllInBytes *int64 `json:"all_in_bytes,omitempty"`

	CombinedCoordinatingAndPrimary ByteSize `json:"combined_coordinating_and_primary,omitempty"`

	CombinedCoordinatingAndPrimaryInBytes *int64 `json:"combined_coordinating_and_primary_in_bytes,omitempty"`

	Coordinating ByteSize `json:"coordinating,omitempty"`

	CoordinatingInBytes *int64 `json:"coordinating_in_bytes,omitempty"`

	CoordinatingRejections   *int64 `json:"coordinating_rejections,omitempty"`
	LargeOperationRejections *int64 `json:"large_operation_rejections,omitempty"`

	Primary                   ByteSize `json:"primary,omitempty"`
	PrimaryDocumentRejections *int64   `json:"primary_document_rejections,omitempty"`

	PrimaryInBytes *int64 `json:"primary_in_bytes,omitempty"`

	PrimaryRejections *int64 `json:"primary_rejections,omitempty"`

	Replica ByteSize `json:"replica,omitempty"`

	ReplicaInBytes *int64 `json:"replica_in_bytes,omitempty"`

	ReplicaRejections *int64 `json:"replica_rejections,omitempty"`
}

func (s *PressureMemory) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPressureMemory() *PressureMemory { _ = "STUB: not implemented"; return nil }
