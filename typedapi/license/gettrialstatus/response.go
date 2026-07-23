package gettrialstatus

type Response struct {
	EligibleToStartTrial bool `json:"eligible_to_start_trial"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
