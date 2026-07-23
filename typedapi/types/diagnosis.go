package types

type Diagnosis struct {
	Action            string                     `json:"action"`
	AffectedResources DiagnosisAffectedResources `json:"affected_resources"`
	Cause             string                     `json:"cause"`
	HelpUrl           string                     `json:"help_url"`
	Id                string                     `json:"id"`
}

func (s *Diagnosis) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDiagnosis() *Diagnosis { _ = "STUB: not implemented"; return nil }
