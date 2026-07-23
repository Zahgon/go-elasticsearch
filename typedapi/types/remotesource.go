package types

type RemoteSource struct {
	ApiKey *string `json:"api_key,omitempty"`

	ConnectTimeout Duration `json:"connect_timeout,omitempty"`

	Headers map[string]string `json:"headers,omitempty"`

	Host string `json:"host"`

	Password *string `json:"password,omitempty"`

	SocketTimeout Duration `json:"socket_timeout,omitempty"`

	Username *string `json:"username,omitempty"`
}

func (s *RemoteSource) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRemoteSource() *RemoteSource { _ = "STUB: not implemented"; return nil }

type RemoteSourceVariant interface {
	RemoteSourceCaster() *RemoteSource
}

func (s *RemoteSource) RemoteSourceCaster() *RemoteSource { _ = "STUB: not implemented"; return nil }
