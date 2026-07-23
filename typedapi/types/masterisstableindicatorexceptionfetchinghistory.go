package types

type MasterIsStableIndicatorExceptionFetchingHistory struct {
	Message    string `json:"message"`
	StackTrace string `json:"stack_trace"`
}

func (s *MasterIsStableIndicatorExceptionFetchingHistory) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMasterIsStableIndicatorExceptionFetchingHistory() *MasterIsStableIndicatorExceptionFetchingHistory {
	_ = "STUB: not implemented"
	return nil
}
