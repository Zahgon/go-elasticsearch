package previewdataframeanalytics

type Response struct {
	FeatureValues []map[string]string `json:"feature_values"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
