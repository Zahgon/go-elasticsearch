package getquery

type Response struct {
	CoordinatingNode string   `json:"coordinating_node"`
	DataNodes        []string `json:"data_nodes"`
	Id               int64    `json:"id"`
	Node             string   `json:"node"`
	Query            string   `json:"query"`
	RunningTimeNanos int64    `json:"running_time_nanos"`
	StartTimeMillis  int64    `json:"start_time_millis"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
