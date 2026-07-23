package types

type Transport struct {
	InboundHandlingTimeHistogram []TransportHistogram `json:"inbound_handling_time_histogram,omitempty"`

	OutboundHandlingTimeHistogram []TransportHistogram `json:"outbound_handling_time_histogram,omitempty"`

	RxCount *int64 `json:"rx_count,omitempty"`

	RxSize *string `json:"rx_size,omitempty"`

	RxSizeInBytes *int64 `json:"rx_size_in_bytes,omitempty"`

	ServerOpen *int `json:"server_open,omitempty"`

	TotalOutboundConnections *int64 `json:"total_outbound_connections,omitempty"`

	TxCount *int64 `json:"tx_count,omitempty"`

	TxSize *string `json:"tx_size,omitempty"`

	TxSizeInBytes *int64 `json:"tx_size_in_bytes,omitempty"`
}

func (s *Transport) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTransport() *Transport { _ = "STUB: not implemented"; return nil }
