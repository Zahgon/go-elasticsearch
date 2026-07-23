package types

type Defaults struct {
	AnomalyDetectors AnomalyDetectors `json:"anomaly_detectors"`
	Datafeeds        Datafeeds        `json:"datafeeds"`
}

func NewDefaults() *Defaults { _ = "STUB: not implemented"; return nil }
