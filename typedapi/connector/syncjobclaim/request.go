package syncjobclaim

import (
	"encoding/json"
)

type Request struct {
	SyncCursor json.RawMessage `json:"sync_cursor,omitempty"`

	WorkerHostname string `json:"worker_hostname"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
