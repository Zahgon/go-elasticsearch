package migratetodatatiers

type Response struct {
	DryRun bool `json:"dry_run"`

	MigratedComponentTemplates []string `json:"migrated_component_templates"`

	MigratedComposableTemplates []string `json:"migrated_composable_templates"`

	MigratedIlmPolicies []string `json:"migrated_ilm_policies"`

	MigratedIndices []string `json:"migrated_indices"`

	MigratedLegacyTemplates []string `json:"migrated_legacy_templates"`

	RemovedLegacyTemplate string `json:"removed_legacy_template"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
