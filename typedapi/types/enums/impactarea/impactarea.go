package impactarea

type ImpactArea struct {
	Name string
}

var (
	Search = ImpactArea{"search"}

	Ingest = ImpactArea{"ingest"}

	Backup = ImpactArea{"backup"}

	Deploymentmanagement = ImpactArea{"deployment_management"}
)

func (i ImpactArea) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ImpactArea) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (i ImpactArea) String() string { _ = "STUB: not implemented"; return "" }
