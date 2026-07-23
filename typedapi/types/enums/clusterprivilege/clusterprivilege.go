package clusterprivilege

type ClusterPrivilege struct {
	Name string
}

var (
	All = ClusterPrivilege{"all"}

	Canceltask = ClusterPrivilege{"cancel_task"}

	Createsnapshot = ClusterPrivilege{"create_snapshot"}

	Crossclusterreplication = ClusterPrivilege{"cross_cluster_replication"}

	Crossclustersearch = ClusterPrivilege{"cross_cluster_search"}

	Delegatepki = ClusterPrivilege{"delegate_pki"}

	Grantapikey = ClusterPrivilege{"grant_api_key"}

	Manage = ClusterPrivilege{"manage"}

	Manageapikey = ClusterPrivilege{"manage_api_key"}

	Manageautoscaling = ClusterPrivilege{"manage_autoscaling"}

	Managebehavioralanalytics = ClusterPrivilege{"manage_behavioral_analytics"}

	Manageccr = ClusterPrivilege{"manage_ccr"}

	Managedataframetransforms = ClusterPrivilege{"manage_data_frame_transforms"}

	Managedatastreamglobalretention = ClusterPrivilege{"manage_data_stream_global_retention"}

	Manageenrich = ClusterPrivilege{"manage_enrich"}

	Manageesql = ClusterPrivilege{"manage_esql"}

	Manageilm = ClusterPrivilege{"manage_ilm"}

	Manageindextemplates = ClusterPrivilege{"manage_index_templates"}

	Manageinference = ClusterPrivilege{"manage_inference"}

	Manageingestpipelines = ClusterPrivilege{"manage_ingest_pipelines"}

	Managelogstashpipelines = ClusterPrivilege{"manage_logstash_pipelines"}

	Manageml = ClusterPrivilege{"manage_ml"}

	Manageoidc = ClusterPrivilege{"manage_oidc"}

	Manageownapikey = ClusterPrivilege{"manage_own_api_key"}

	Managepipeline = ClusterPrivilege{"manage_pipeline"}

	Managereindex = ClusterPrivilege{"manage_reindex"}

	Managerollup = ClusterPrivilege{"manage_rollup"}

	Managesaml = ClusterPrivilege{"manage_saml"}

	Managesearchapplication = ClusterPrivilege{"manage_search_application"}

	Managesearchqueryrules = ClusterPrivilege{"manage_search_query_rules"}

	Managesearchsynonyms = ClusterPrivilege{"manage_search_synonyms"}

	Managesecurity = ClusterPrivilege{"manage_security"}

	Manageserviceaccount = ClusterPrivilege{"manage_service_account"}

	Manageslm = ClusterPrivilege{"manage_slm"}

	Managetoken = ClusterPrivilege{"manage_token"}

	Managetransform = ClusterPrivilege{"manage_transform"}

	Manageuserprofile = ClusterPrivilege{"manage_user_profile"}

	Managewatcher = ClusterPrivilege{"manage_watcher"}

	Monitor = ClusterPrivilege{"monitor"}

	Monitordataframetransforms = ClusterPrivilege{"monitor_data_frame_transforms"}

	Monitordatastreamglobalretention = ClusterPrivilege{"monitor_data_stream_global_retention"}

	Monitorenrich = ClusterPrivilege{"monitor_enrich"}

	Monitoresql = ClusterPrivilege{"monitor_esql"}

	Monitorinference = ClusterPrivilege{"monitor_inference"}

	Monitorml = ClusterPrivilege{"monitor_ml"}

	Monitorreindex = ClusterPrivilege{"monitor_reindex"}

	Monitorrollup = ClusterPrivilege{"monitor_rollup"}

	Monitorsnapshot = ClusterPrivilege{"monitor_snapshot"}

	Monitorstats = ClusterPrivilege{"monitor_stats"}

	Monitortextstructure = ClusterPrivilege{"monitor_text_structure"}

	Monitortransform = ClusterPrivilege{"monitor_transform"}

	Monitorwatcher = ClusterPrivilege{"monitor_watcher"}

	None = ClusterPrivilege{"none"}

	Postbehavioralanalyticsevent = ClusterPrivilege{"post_behavioral_analytics_event"}

	Readccr = ClusterPrivilege{"read_ccr"}

	Readfleetsecrets = ClusterPrivilege{"read_fleet_secrets"}

	Readilm = ClusterPrivilege{"read_ilm"}

	Readpipeline = ClusterPrivilege{"read_pipeline"}

	Readsecurity = ClusterPrivilege{"read_security"}

	Readslm = ClusterPrivilege{"read_slm"}

	Transportclient = ClusterPrivilege{"transport_client"}

	Writeconnectorsecrets = ClusterPrivilege{"write_connector_secrets"}

	Writefleetsecrets = ClusterPrivilege{"write_fleet_secrets"}

	Readprojectrouting = ClusterPrivilege{"read_project_routing"}

	Manageprojectrouting = ClusterPrivilege{"manage_project_routing"}
)

func (c ClusterPrivilege) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClusterPrivilege) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c ClusterPrivilege) String() string { _ = "STUB: not implemented"; return "" }
