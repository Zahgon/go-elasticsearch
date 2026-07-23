package types

type ThreadPoolRecord struct {
	Active *string `json:"active,omitempty"`

	Completed *string `json:"completed,omitempty"`

	Core *string `json:"core,omitempty"`

	EphemeralNodeId *string `json:"ephemeral_node_id,omitempty"`

	Host *string `json:"host,omitempty"`

	Ip *string `json:"ip,omitempty"`

	KeepAlive *string `json:"keep_alive,omitempty"`

	Largest *string `json:"largest,omitempty"`

	Max *string `json:"max,omitempty"`

	Name *string `json:"name,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	NodeName *string `json:"node_name,omitempty"`

	Pid *string `json:"pid,omitempty"`

	PoolSize *string `json:"pool_size,omitempty"`

	Port *string `json:"port,omitempty"`

	Queue *string `json:"queue,omitempty"`

	QueueSize *string `json:"queue_size,omitempty"`

	Rejected *string `json:"rejected,omitempty"`

	Size *string `json:"size,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (s *ThreadPoolRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewThreadPoolRecord() *ThreadPoolRecord { _ = "STUB: not implemented"; return nil }
