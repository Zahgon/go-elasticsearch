package typedapi

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	async_search_delete "github.com/elastic/go-elasticsearch/v9/typedapi/asyncsearch/delete"
	async_search_get "github.com/elastic/go-elasticsearch/v9/typedapi/asyncsearch/get"
	async_search_status "github.com/elastic/go-elasticsearch/v9/typedapi/asyncsearch/status"
	async_search_submit "github.com/elastic/go-elasticsearch/v9/typedapi/asyncsearch/submit"
	autoscaling_delete_autoscaling_policy "github.com/elastic/go-elasticsearch/v9/typedapi/autoscaling/deleteautoscalingpolicy"
	autoscaling_get_autoscaling_capacity "github.com/elastic/go-elasticsearch/v9/typedapi/autoscaling/getautoscalingcapacity"
	autoscaling_get_autoscaling_policy "github.com/elastic/go-elasticsearch/v9/typedapi/autoscaling/getautoscalingpolicy"
	autoscaling_put_autoscaling_policy "github.com/elastic/go-elasticsearch/v9/typedapi/autoscaling/putautoscalingpolicy"
	cat_aliases "github.com/elastic/go-elasticsearch/v9/typedapi/cat/aliases"
	cat_allocation "github.com/elastic/go-elasticsearch/v9/typedapi/cat/allocation"
	cat_circuit_breaker "github.com/elastic/go-elasticsearch/v9/typedapi/cat/circuitbreaker"
	cat_component_templates "github.com/elastic/go-elasticsearch/v9/typedapi/cat/componenttemplates"
	cat_count "github.com/elastic/go-elasticsearch/v9/typedapi/cat/count"
	cat_fielddata "github.com/elastic/go-elasticsearch/v9/typedapi/cat/fielddata"
	cat_health "github.com/elastic/go-elasticsearch/v9/typedapi/cat/health"
	cat_help "github.com/elastic/go-elasticsearch/v9/typedapi/cat/help"
	cat_indices "github.com/elastic/go-elasticsearch/v9/typedapi/cat/indices"
	cat_master "github.com/elastic/go-elasticsearch/v9/typedapi/cat/master"
	cat_ml_datafeeds "github.com/elastic/go-elasticsearch/v9/typedapi/cat/mldatafeeds"
	cat_ml_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/cat/mldataframeanalytics"
	cat_ml_jobs "github.com/elastic/go-elasticsearch/v9/typedapi/cat/mljobs"
	cat_ml_trained_models "github.com/elastic/go-elasticsearch/v9/typedapi/cat/mltrainedmodels"
	cat_nodeattrs "github.com/elastic/go-elasticsearch/v9/typedapi/cat/nodeattrs"
	cat_nodes "github.com/elastic/go-elasticsearch/v9/typedapi/cat/nodes"
	cat_pending_tasks "github.com/elastic/go-elasticsearch/v9/typedapi/cat/pendingtasks"
	cat_plugins "github.com/elastic/go-elasticsearch/v9/typedapi/cat/plugins"
	cat_recovery "github.com/elastic/go-elasticsearch/v9/typedapi/cat/recovery"
	cat_repositories "github.com/elastic/go-elasticsearch/v9/typedapi/cat/repositories"
	cat_segments "github.com/elastic/go-elasticsearch/v9/typedapi/cat/segments"
	cat_shards "github.com/elastic/go-elasticsearch/v9/typedapi/cat/shards"
	cat_snapshots "github.com/elastic/go-elasticsearch/v9/typedapi/cat/snapshots"
	cat_tasks "github.com/elastic/go-elasticsearch/v9/typedapi/cat/tasks"
	cat_templates "github.com/elastic/go-elasticsearch/v9/typedapi/cat/templates"
	cat_thread_pool "github.com/elastic/go-elasticsearch/v9/typedapi/cat/threadpool"
	cat_transforms "github.com/elastic/go-elasticsearch/v9/typedapi/cat/transforms"
	ccr_delete_auto_follow_pattern "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/deleteautofollowpattern"
	ccr_follow "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/follow"
	ccr_follow_info "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/followinfo"
	ccr_follow_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/followstats"
	ccr_forget_follower "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/forgetfollower"
	ccr_get_auto_follow_pattern "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/getautofollowpattern"
	ccr_pause_auto_follow_pattern "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/pauseautofollowpattern"
	ccr_pause_follow "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/pausefollow"
	ccr_put_auto_follow_pattern "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/putautofollowpattern"
	ccr_resume_auto_follow_pattern "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/resumeautofollowpattern"
	ccr_resume_follow "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/resumefollow"
	ccr_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/stats"
	ccr_unfollow "github.com/elastic/go-elasticsearch/v9/typedapi/ccr/unfollow"
	cluster_allocation_explain "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/allocationexplain"
	cluster_delete_component_template "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/deletecomponenttemplate"
	cluster_delete_voting_config_exclusions "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/deletevotingconfigexclusions"
	cluster_exists_component_template "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/existscomponenttemplate"
	cluster_get_component_template "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/getcomponenttemplate"
	cluster_get_settings "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/getsettings"
	cluster_health "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/health"
	cluster_info "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/info"
	cluster_pending_tasks "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/pendingtasks"
	cluster_post_voting_config_exclusions "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/postvotingconfigexclusions"
	cluster_put_component_template "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/putcomponenttemplate"
	cluster_put_settings "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/putsettings"
	cluster_remote_info "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/remoteinfo"
	cluster_reroute "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/reroute"
	cluster_state "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/state"
	cluster_stats "github.com/elastic/go-elasticsearch/v9/typedapi/cluster/stats"
	connector_check_in "github.com/elastic/go-elasticsearch/v9/typedapi/connector/checkin"
	connector_delete "github.com/elastic/go-elasticsearch/v9/typedapi/connector/delete"
	connector_get "github.com/elastic/go-elasticsearch/v9/typedapi/connector/get"
	connector_last_sync "github.com/elastic/go-elasticsearch/v9/typedapi/connector/lastsync"
	connector_list "github.com/elastic/go-elasticsearch/v9/typedapi/connector/list"
	connector_post "github.com/elastic/go-elasticsearch/v9/typedapi/connector/post"
	connector_put "github.com/elastic/go-elasticsearch/v9/typedapi/connector/put"
	connector_secret_delete "github.com/elastic/go-elasticsearch/v9/typedapi/connector/secretdelete"
	connector_secret_get "github.com/elastic/go-elasticsearch/v9/typedapi/connector/secretget"
	connector_secret_post "github.com/elastic/go-elasticsearch/v9/typedapi/connector/secretpost"
	connector_secret_put "github.com/elastic/go-elasticsearch/v9/typedapi/connector/secretput"
	connector_sync_job_cancel "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjobcancel"
	connector_sync_job_check_in "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjobcheckin"
	connector_sync_job_claim "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjobclaim"
	connector_sync_job_delete "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjobdelete"
	connector_sync_job_error "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjoberror"
	connector_sync_job_get "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjobget"
	connector_sync_job_list "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjoblist"
	connector_sync_job_post "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjobpost"
	connector_sync_job_update_stats "github.com/elastic/go-elasticsearch/v9/typedapi/connector/syncjobupdatestats"
	connector_update_active_filtering "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updateactivefiltering"
	connector_update_api_key_id "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updateapikeyid"
	connector_update_configuration "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updateconfiguration"
	connector_update_error "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updateerror"
	connector_update_features "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updatefeatures"
	connector_update_filtering "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updatefiltering"
	connector_update_filtering_validation "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updatefilteringvalidation"
	connector_update_index_name "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updateindexname"
	connector_update_name "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updatename"
	connector_update_native "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updatenative"
	connector_update_pipeline "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updatepipeline"
	connector_update_scheduling "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updatescheduling"
	connector_update_service_type "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updateservicetype"
	connector_update_status "github.com/elastic/go-elasticsearch/v9/typedapi/connector/updatestatus"
	core_bulk "github.com/elastic/go-elasticsearch/v9/typedapi/core/bulk"
	core_cancel_reindex "github.com/elastic/go-elasticsearch/v9/typedapi/core/cancelreindex"
	core_capabilities "github.com/elastic/go-elasticsearch/v9/typedapi/core/capabilities"
	core_clear_scroll "github.com/elastic/go-elasticsearch/v9/typedapi/core/clearscroll"
	core_close_point_in_time "github.com/elastic/go-elasticsearch/v9/typedapi/core/closepointintime"
	core_count "github.com/elastic/go-elasticsearch/v9/typedapi/core/count"
	core_create "github.com/elastic/go-elasticsearch/v9/typedapi/core/create"
	core_delete "github.com/elastic/go-elasticsearch/v9/typedapi/core/delete"
	core_delete_by_query "github.com/elastic/go-elasticsearch/v9/typedapi/core/deletebyquery"
	core_delete_by_query_rethrottle "github.com/elastic/go-elasticsearch/v9/typedapi/core/deletebyqueryrethrottle"
	core_delete_script "github.com/elastic/go-elasticsearch/v9/typedapi/core/deletescript"
	core_exists "github.com/elastic/go-elasticsearch/v9/typedapi/core/exists"
	core_exists_source "github.com/elastic/go-elasticsearch/v9/typedapi/core/existssource"
	core_explain "github.com/elastic/go-elasticsearch/v9/typedapi/core/explain"
	core_field_caps "github.com/elastic/go-elasticsearch/v9/typedapi/core/fieldcaps"
	core_get "github.com/elastic/go-elasticsearch/v9/typedapi/core/get"
	core_get_reindex "github.com/elastic/go-elasticsearch/v9/typedapi/core/getreindex"
	core_get_script "github.com/elastic/go-elasticsearch/v9/typedapi/core/getscript"
	core_get_script_context "github.com/elastic/go-elasticsearch/v9/typedapi/core/getscriptcontext"
	core_get_script_languages "github.com/elastic/go-elasticsearch/v9/typedapi/core/getscriptlanguages"
	core_get_source "github.com/elastic/go-elasticsearch/v9/typedapi/core/getsource"
	core_health_report "github.com/elastic/go-elasticsearch/v9/typedapi/core/healthreport"
	core_index "github.com/elastic/go-elasticsearch/v9/typedapi/core/index"
	core_info "github.com/elastic/go-elasticsearch/v9/typedapi/core/info"
	core_knn_search "github.com/elastic/go-elasticsearch/v9/typedapi/core/knnsearch"
	core_list_reindex "github.com/elastic/go-elasticsearch/v9/typedapi/core/listreindex"
	core_mget "github.com/elastic/go-elasticsearch/v9/typedapi/core/mget"
	core_msearch "github.com/elastic/go-elasticsearch/v9/typedapi/core/msearch"
	core_msearch_template "github.com/elastic/go-elasticsearch/v9/typedapi/core/msearchtemplate"
	core_mtermvectors "github.com/elastic/go-elasticsearch/v9/typedapi/core/mtermvectors"
	core_open_point_in_time "github.com/elastic/go-elasticsearch/v9/typedapi/core/openpointintime"
	core_ping "github.com/elastic/go-elasticsearch/v9/typedapi/core/ping"
	core_put_script "github.com/elastic/go-elasticsearch/v9/typedapi/core/putscript"
	core_rank_eval "github.com/elastic/go-elasticsearch/v9/typedapi/core/rankeval"
	core_reindex "github.com/elastic/go-elasticsearch/v9/typedapi/core/reindex"
	core_reindex_rethrottle "github.com/elastic/go-elasticsearch/v9/typedapi/core/reindexrethrottle"
	core_render_search_template "github.com/elastic/go-elasticsearch/v9/typedapi/core/rendersearchtemplate"
	core_scripts_painless_execute "github.com/elastic/go-elasticsearch/v9/typedapi/core/scriptspainlessexecute"
	core_scroll "github.com/elastic/go-elasticsearch/v9/typedapi/core/scroll"
	core_search "github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	core_search_mvt "github.com/elastic/go-elasticsearch/v9/typedapi/core/searchmvt"
	core_search_shards "github.com/elastic/go-elasticsearch/v9/typedapi/core/searchshards"
	core_search_template "github.com/elastic/go-elasticsearch/v9/typedapi/core/searchtemplate"
	core_terms_enum "github.com/elastic/go-elasticsearch/v9/typedapi/core/termsenum"
	core_termvectors "github.com/elastic/go-elasticsearch/v9/typedapi/core/termvectors"
	core_update "github.com/elastic/go-elasticsearch/v9/typedapi/core/update"
	core_update_by_query "github.com/elastic/go-elasticsearch/v9/typedapi/core/updatebyquery"
	core_update_by_query_rethrottle "github.com/elastic/go-elasticsearch/v9/typedapi/core/updatebyqueryrethrottle"
	dangling_indices_delete_dangling_index "github.com/elastic/go-elasticsearch/v9/typedapi/danglingindices/deletedanglingindex"
	dangling_indices_import_dangling_index "github.com/elastic/go-elasticsearch/v9/typedapi/danglingindices/importdanglingindex"
	dangling_indices_list_dangling_indices "github.com/elastic/go-elasticsearch/v9/typedapi/danglingindices/listdanglingindices"
	enrich_delete_policy "github.com/elastic/go-elasticsearch/v9/typedapi/enrich/deletepolicy"
	enrich_execute_policy "github.com/elastic/go-elasticsearch/v9/typedapi/enrich/executepolicy"
	enrich_get_policy "github.com/elastic/go-elasticsearch/v9/typedapi/enrich/getpolicy"
	enrich_put_policy "github.com/elastic/go-elasticsearch/v9/typedapi/enrich/putpolicy"
	enrich_stats "github.com/elastic/go-elasticsearch/v9/typedapi/enrich/stats"
	eql_delete "github.com/elastic/go-elasticsearch/v9/typedapi/eql/delete"
	eql_get "github.com/elastic/go-elasticsearch/v9/typedapi/eql/get"
	eql_get_status "github.com/elastic/go-elasticsearch/v9/typedapi/eql/getstatus"
	eql_search "github.com/elastic/go-elasticsearch/v9/typedapi/eql/search"
	esql_async_query "github.com/elastic/go-elasticsearch/v9/typedapi/esql/asyncquery"
	esql_async_query_delete "github.com/elastic/go-elasticsearch/v9/typedapi/esql/asyncquerydelete"
	esql_async_query_get "github.com/elastic/go-elasticsearch/v9/typedapi/esql/asyncqueryget"
	esql_async_query_stop "github.com/elastic/go-elasticsearch/v9/typedapi/esql/asyncquerystop"
	esql_delete_view "github.com/elastic/go-elasticsearch/v9/typedapi/esql/deleteview"
	esql_get_query "github.com/elastic/go-elasticsearch/v9/typedapi/esql/getquery"
	esql_get_view "github.com/elastic/go-elasticsearch/v9/typedapi/esql/getview"
	esql_list_queries "github.com/elastic/go-elasticsearch/v9/typedapi/esql/listqueries"
	esql_put_view "github.com/elastic/go-elasticsearch/v9/typedapi/esql/putview"
	esql_query "github.com/elastic/go-elasticsearch/v9/typedapi/esql/query"
	features_get_features "github.com/elastic/go-elasticsearch/v9/typedapi/features/getfeatures"
	features_reset_features "github.com/elastic/go-elasticsearch/v9/typedapi/features/resetfeatures"
	fleet_delete_secret "github.com/elastic/go-elasticsearch/v9/typedapi/fleet/deletesecret"
	fleet_get_secret "github.com/elastic/go-elasticsearch/v9/typedapi/fleet/getsecret"
	fleet_global_checkpoints "github.com/elastic/go-elasticsearch/v9/typedapi/fleet/globalcheckpoints"
	fleet_msearch "github.com/elastic/go-elasticsearch/v9/typedapi/fleet/msearch"
	fleet_post_secret "github.com/elastic/go-elasticsearch/v9/typedapi/fleet/postsecret"
	fleet_search "github.com/elastic/go-elasticsearch/v9/typedapi/fleet/search"
	graph_explore "github.com/elastic/go-elasticsearch/v9/typedapi/graph/explore"
	ilm_delete_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/deletelifecycle"
	ilm_explain_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/explainlifecycle"
	ilm_get_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/getlifecycle"
	ilm_get_status "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/getstatus"
	ilm_migrate_to_data_tiers "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/migratetodatatiers"
	ilm_move_to_step "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/movetostep"
	ilm_put_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/putlifecycle"
	ilm_remove_policy "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/removepolicy"
	ilm_retry "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/retry"
	ilm_start "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/start"
	ilm_stop "github.com/elastic/go-elasticsearch/v9/typedapi/ilm/stop"
	indices_add_block "github.com/elastic/go-elasticsearch/v9/typedapi/indices/addblock"
	indices_analyze "github.com/elastic/go-elasticsearch/v9/typedapi/indices/analyze"
	indices_cancel_migrate_reindex "github.com/elastic/go-elasticsearch/v9/typedapi/indices/cancelmigratereindex"
	indices_clear_cache "github.com/elastic/go-elasticsearch/v9/typedapi/indices/clearcache"
	indices_clone "github.com/elastic/go-elasticsearch/v9/typedapi/indices/clone"
	indices_close "github.com/elastic/go-elasticsearch/v9/typedapi/indices/close"
	indices_create "github.com/elastic/go-elasticsearch/v9/typedapi/indices/create"
	indices_create_data_stream "github.com/elastic/go-elasticsearch/v9/typedapi/indices/createdatastream"
	indices_create_from "github.com/elastic/go-elasticsearch/v9/typedapi/indices/createfrom"
	indices_data_streams_stats "github.com/elastic/go-elasticsearch/v9/typedapi/indices/datastreamsstats"
	indices_delete "github.com/elastic/go-elasticsearch/v9/typedapi/indices/delete"
	indices_delete_alias "github.com/elastic/go-elasticsearch/v9/typedapi/indices/deletealias"
	indices_delete_data_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/indices/deletedatalifecycle"
	indices_delete_data_stream "github.com/elastic/go-elasticsearch/v9/typedapi/indices/deletedatastream"
	indices_delete_data_stream_options "github.com/elastic/go-elasticsearch/v9/typedapi/indices/deletedatastreamoptions"
	indices_delete_index_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/deleteindextemplate"
	indices_delete_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/deletetemplate"
	indices_disk_usage "github.com/elastic/go-elasticsearch/v9/typedapi/indices/diskusage"
	indices_downsample "github.com/elastic/go-elasticsearch/v9/typedapi/indices/downsample"
	indices_exists "github.com/elastic/go-elasticsearch/v9/typedapi/indices/exists"
	indices_exists_alias "github.com/elastic/go-elasticsearch/v9/typedapi/indices/existsalias"
	indices_exists_index_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/existsindextemplate"
	indices_exists_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/existstemplate"
	indices_explain_data_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/indices/explaindatalifecycle"
	indices_field_usage_stats "github.com/elastic/go-elasticsearch/v9/typedapi/indices/fieldusagestats"
	indices_flush "github.com/elastic/go-elasticsearch/v9/typedapi/indices/flush"
	indices_forcemerge "github.com/elastic/go-elasticsearch/v9/typedapi/indices/forcemerge"
	indices_get "github.com/elastic/go-elasticsearch/v9/typedapi/indices/get"
	indices_get_alias "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getalias"
	indices_get_data_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getdatalifecycle"
	indices_get_data_lifecycle_stats "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getdatalifecyclestats"
	indices_get_data_stream "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getdatastream"
	indices_get_data_stream_mappings "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getdatastreammappings"
	indices_get_data_stream_options "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getdatastreamoptions"
	indices_get_data_stream_settings "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getdatastreamsettings"
	indices_get_field_mapping "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getfieldmapping"
	indices_get_index_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getindextemplate"
	indices_get_mapping "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getmapping"
	indices_get_migrate_reindex_status "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getmigratereindexstatus"
	indices_get_settings "github.com/elastic/go-elasticsearch/v9/typedapi/indices/getsettings"
	indices_get_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/gettemplate"
	indices_migrate_reindex "github.com/elastic/go-elasticsearch/v9/typedapi/indices/migratereindex"
	indices_migrate_to_data_stream "github.com/elastic/go-elasticsearch/v9/typedapi/indices/migratetodatastream"
	indices_modify_data_stream "github.com/elastic/go-elasticsearch/v9/typedapi/indices/modifydatastream"
	indices_open "github.com/elastic/go-elasticsearch/v9/typedapi/indices/open"
	indices_promote_data_stream "github.com/elastic/go-elasticsearch/v9/typedapi/indices/promotedatastream"
	indices_put_alias "github.com/elastic/go-elasticsearch/v9/typedapi/indices/putalias"
	indices_put_data_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/indices/putdatalifecycle"
	indices_put_data_stream_mappings "github.com/elastic/go-elasticsearch/v9/typedapi/indices/putdatastreammappings"
	indices_put_data_stream_options "github.com/elastic/go-elasticsearch/v9/typedapi/indices/putdatastreamoptions"
	indices_put_data_stream_settings "github.com/elastic/go-elasticsearch/v9/typedapi/indices/putdatastreamsettings"
	indices_put_index_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/putindextemplate"
	indices_put_mapping "github.com/elastic/go-elasticsearch/v9/typedapi/indices/putmapping"
	indices_put_settings "github.com/elastic/go-elasticsearch/v9/typedapi/indices/putsettings"
	indices_put_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/puttemplate"
	indices_recovery "github.com/elastic/go-elasticsearch/v9/typedapi/indices/recovery"
	indices_refresh "github.com/elastic/go-elasticsearch/v9/typedapi/indices/refresh"
	indices_reload_search_analyzers "github.com/elastic/go-elasticsearch/v9/typedapi/indices/reloadsearchanalyzers"
	indices_remove_block "github.com/elastic/go-elasticsearch/v9/typedapi/indices/removeblock"
	indices_resolve_cluster "github.com/elastic/go-elasticsearch/v9/typedapi/indices/resolvecluster"
	indices_resolve_index "github.com/elastic/go-elasticsearch/v9/typedapi/indices/resolveindex"
	indices_rollover "github.com/elastic/go-elasticsearch/v9/typedapi/indices/rollover"
	indices_segments "github.com/elastic/go-elasticsearch/v9/typedapi/indices/segments"
	indices_shard_stores "github.com/elastic/go-elasticsearch/v9/typedapi/indices/shardstores"
	indices_shrink "github.com/elastic/go-elasticsearch/v9/typedapi/indices/shrink"
	indices_simulate_index_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/simulateindextemplate"
	indices_simulate_template "github.com/elastic/go-elasticsearch/v9/typedapi/indices/simulatetemplate"
	indices_split "github.com/elastic/go-elasticsearch/v9/typedapi/indices/split"
	indices_stats "github.com/elastic/go-elasticsearch/v9/typedapi/indices/stats"
	indices_update_aliases "github.com/elastic/go-elasticsearch/v9/typedapi/indices/updatealiases"
	indices_validate_query "github.com/elastic/go-elasticsearch/v9/typedapi/indices/validatequery"
	inference_chat_completion_unified "github.com/elastic/go-elasticsearch/v9/typedapi/inference/chatcompletionunified"
	inference_completion "github.com/elastic/go-elasticsearch/v9/typedapi/inference/completion"
	inference_delete "github.com/elastic/go-elasticsearch/v9/typedapi/inference/delete"
	inference_embedding "github.com/elastic/go-elasticsearch/v9/typedapi/inference/embedding"
	inference_get "github.com/elastic/go-elasticsearch/v9/typedapi/inference/get"
	inference_inference "github.com/elastic/go-elasticsearch/v9/typedapi/inference/inference"
	inference_put "github.com/elastic/go-elasticsearch/v9/typedapi/inference/put"
	inference_put_ai21 "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putai21"
	inference_put_alibabacloud "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putalibabacloud"
	inference_put_amazonbedrock "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putamazonbedrock"
	inference_put_amazonsagemaker "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putamazonsagemaker"
	inference_put_anthropic "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putanthropic"
	inference_put_azureaistudio "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putazureaistudio"
	inference_put_azureopenai "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putazureopenai"
	inference_put_cohere "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putcohere"
	inference_put_contextualai "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putcontextualai"
	inference_put_custom "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putcustom"
	inference_put_deepseek "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putdeepseek"
	inference_put_elasticsearch "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putelasticsearch"
	inference_put_elser "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putelser"
	inference_put_fireworksai "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putfireworksai"
	inference_put_googleaistudio "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putgoogleaistudio"
	inference_put_googlevertexai "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putgooglevertexai"
	inference_put_groq "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putgroq"
	inference_put_hugging_face "github.com/elastic/go-elasticsearch/v9/typedapi/inference/puthuggingface"
	inference_put_jinaai "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putjinaai"
	inference_put_llama "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putllama"
	inference_put_mistral "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putmistral"
	inference_put_nvidia "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putnvidia"
	inference_put_openai "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putopenai"
	inference_put_openshift_ai "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putopenshiftai"
	inference_put_voyageai "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putvoyageai"
	inference_put_watsonx "github.com/elastic/go-elasticsearch/v9/typedapi/inference/putwatsonx"
	inference_rerank "github.com/elastic/go-elasticsearch/v9/typedapi/inference/rerank"
	inference_sparse_embedding "github.com/elastic/go-elasticsearch/v9/typedapi/inference/sparseembedding"
	inference_stream_completion "github.com/elastic/go-elasticsearch/v9/typedapi/inference/streamcompletion"
	inference_text_embedding "github.com/elastic/go-elasticsearch/v9/typedapi/inference/textembedding"
	inference_update "github.com/elastic/go-elasticsearch/v9/typedapi/inference/update"
	ingest_delete_geoip_database "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/deletegeoipdatabase"
	ingest_delete_ip_location_database "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/deleteiplocationdatabase"
	ingest_delete_pipeline "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/deletepipeline"
	ingest_geo_ip_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/geoipstats"
	ingest_get_geoip_database "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/getgeoipdatabase"
	ingest_get_ip_location_database "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/getiplocationdatabase"
	ingest_get_pipeline "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/getpipeline"
	ingest_processor_grok "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/processorgrok"
	ingest_put_geoip_database "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/putgeoipdatabase"
	ingest_put_ip_location_database "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/putiplocationdatabase"
	ingest_put_pipeline "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/putpipeline"
	ingest_simulate "github.com/elastic/go-elasticsearch/v9/typedapi/ingest/simulate"
	license_delete "github.com/elastic/go-elasticsearch/v9/typedapi/license/delete"
	license_get "github.com/elastic/go-elasticsearch/v9/typedapi/license/get"
	license_get_basic_status "github.com/elastic/go-elasticsearch/v9/typedapi/license/getbasicstatus"
	license_get_trial_status "github.com/elastic/go-elasticsearch/v9/typedapi/license/gettrialstatus"
	license_post "github.com/elastic/go-elasticsearch/v9/typedapi/license/post"
	license_post_start_basic "github.com/elastic/go-elasticsearch/v9/typedapi/license/poststartbasic"
	license_post_start_trial "github.com/elastic/go-elasticsearch/v9/typedapi/license/poststarttrial"
	logstash_delete_pipeline "github.com/elastic/go-elasticsearch/v9/typedapi/logstash/deletepipeline"
	logstash_get_pipeline "github.com/elastic/go-elasticsearch/v9/typedapi/logstash/getpipeline"
	logstash_put_pipeline "github.com/elastic/go-elasticsearch/v9/typedapi/logstash/putpipeline"
	migration_deprecations "github.com/elastic/go-elasticsearch/v9/typedapi/migration/deprecations"
	migration_get_feature_upgrade_status "github.com/elastic/go-elasticsearch/v9/typedapi/migration/getfeatureupgradestatus"
	migration_post_feature_upgrade "github.com/elastic/go-elasticsearch/v9/typedapi/migration/postfeatureupgrade"
	ml_clear_trained_model_deployment_cache "github.com/elastic/go-elasticsearch/v9/typedapi/ml/cleartrainedmodeldeploymentcache"
	ml_close_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/closejob"
	ml_delete_calendar "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletecalendar"
	ml_delete_calendar_event "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletecalendarevent"
	ml_delete_calendar_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletecalendarjob"
	ml_delete_datafeed "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletedatafeed"
	ml_delete_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletedataframeanalytics"
	ml_delete_expired_data "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deleteexpireddata"
	ml_delete_filter "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletefilter"
	ml_delete_forecast "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deleteforecast"
	ml_delete_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletejob"
	ml_delete_model_snapshot "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletemodelsnapshot"
	ml_delete_trained_model "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletetrainedmodel"
	ml_delete_trained_model_alias "github.com/elastic/go-elasticsearch/v9/typedapi/ml/deletetrainedmodelalias"
	ml_estimate_model_memory "github.com/elastic/go-elasticsearch/v9/typedapi/ml/estimatemodelmemory"
	ml_evaluate_data_frame "github.com/elastic/go-elasticsearch/v9/typedapi/ml/evaluatedataframe"
	ml_explain_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/ml/explaindataframeanalytics"
	ml_flush_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/flushjob"
	ml_forecast "github.com/elastic/go-elasticsearch/v9/typedapi/ml/forecast"
	ml_get_buckets "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getbuckets"
	ml_get_calendar_events "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getcalendarevents"
	ml_get_calendars "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getcalendars"
	ml_get_categories "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getcategories"
	ml_get_datafeeds "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getdatafeeds"
	ml_get_datafeed_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getdatafeedstats"
	ml_get_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getdataframeanalytics"
	ml_get_data_frame_analytics_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getdataframeanalyticsstats"
	ml_get_filters "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getfilters"
	ml_get_influencers "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getinfluencers"
	ml_get_jobs "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getjobs"
	ml_get_job_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getjobstats"
	ml_get_memory_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getmemorystats"
	ml_get_model_snapshots "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getmodelsnapshots"
	ml_get_model_snapshot_upgrade_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getmodelsnapshotupgradestats"
	ml_get_overall_buckets "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getoverallbuckets"
	ml_get_records "github.com/elastic/go-elasticsearch/v9/typedapi/ml/getrecords"
	ml_get_trained_models "github.com/elastic/go-elasticsearch/v9/typedapi/ml/gettrainedmodels"
	ml_get_trained_models_stats "github.com/elastic/go-elasticsearch/v9/typedapi/ml/gettrainedmodelsstats"
	ml_infer_trained_model "github.com/elastic/go-elasticsearch/v9/typedapi/ml/infertrainedmodel"
	ml_info "github.com/elastic/go-elasticsearch/v9/typedapi/ml/info"
	ml_open_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/openjob"
	ml_post_calendar_events "github.com/elastic/go-elasticsearch/v9/typedapi/ml/postcalendarevents"
	ml_post_data "github.com/elastic/go-elasticsearch/v9/typedapi/ml/postdata"
	ml_preview_datafeed "github.com/elastic/go-elasticsearch/v9/typedapi/ml/previewdatafeed"
	ml_preview_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/ml/previewdataframeanalytics"
	ml_put_calendar "github.com/elastic/go-elasticsearch/v9/typedapi/ml/putcalendar"
	ml_put_calendar_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/putcalendarjob"
	ml_put_datafeed "github.com/elastic/go-elasticsearch/v9/typedapi/ml/putdatafeed"
	ml_put_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/ml/putdataframeanalytics"
	ml_put_filter "github.com/elastic/go-elasticsearch/v9/typedapi/ml/putfilter"
	ml_put_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/putjob"
	ml_put_trained_model "github.com/elastic/go-elasticsearch/v9/typedapi/ml/puttrainedmodel"
	ml_put_trained_model_alias "github.com/elastic/go-elasticsearch/v9/typedapi/ml/puttrainedmodelalias"
	ml_put_trained_model_definition_part "github.com/elastic/go-elasticsearch/v9/typedapi/ml/puttrainedmodeldefinitionpart"
	ml_put_trained_model_vocabulary "github.com/elastic/go-elasticsearch/v9/typedapi/ml/puttrainedmodelvocabulary"
	ml_reset_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/resetjob"
	ml_revert_model_snapshot "github.com/elastic/go-elasticsearch/v9/typedapi/ml/revertmodelsnapshot"
	ml_set_upgrade_mode "github.com/elastic/go-elasticsearch/v9/typedapi/ml/setupgrademode"
	ml_start_datafeed "github.com/elastic/go-elasticsearch/v9/typedapi/ml/startdatafeed"
	ml_start_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/ml/startdataframeanalytics"
	ml_start_trained_model_deployment "github.com/elastic/go-elasticsearch/v9/typedapi/ml/starttrainedmodeldeployment"
	ml_stop_datafeed "github.com/elastic/go-elasticsearch/v9/typedapi/ml/stopdatafeed"
	ml_stop_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/ml/stopdataframeanalytics"
	ml_stop_trained_model_deployment "github.com/elastic/go-elasticsearch/v9/typedapi/ml/stoptrainedmodeldeployment"
	ml_update_datafeed "github.com/elastic/go-elasticsearch/v9/typedapi/ml/updatedatafeed"
	ml_update_data_frame_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/ml/updatedataframeanalytics"
	ml_update_filter "github.com/elastic/go-elasticsearch/v9/typedapi/ml/updatefilter"
	ml_update_job "github.com/elastic/go-elasticsearch/v9/typedapi/ml/updatejob"
	ml_update_model_snapshot "github.com/elastic/go-elasticsearch/v9/typedapi/ml/updatemodelsnapshot"
	ml_update_trained_model_deployment "github.com/elastic/go-elasticsearch/v9/typedapi/ml/updatetrainedmodeldeployment"
	ml_upgrade_job_snapshot "github.com/elastic/go-elasticsearch/v9/typedapi/ml/upgradejobsnapshot"
	ml_validate "github.com/elastic/go-elasticsearch/v9/typedapi/ml/validate"
	ml_validate_detector "github.com/elastic/go-elasticsearch/v9/typedapi/ml/validatedetector"
	monitoring_bulk "github.com/elastic/go-elasticsearch/v9/typedapi/monitoring/bulk"
	nodes_clear_repositories_metering_archive "github.com/elastic/go-elasticsearch/v9/typedapi/nodes/clearrepositoriesmeteringarchive"
	nodes_get_repositories_metering_info "github.com/elastic/go-elasticsearch/v9/typedapi/nodes/getrepositoriesmeteringinfo"
	nodes_hot_threads "github.com/elastic/go-elasticsearch/v9/typedapi/nodes/hotthreads"
	nodes_info "github.com/elastic/go-elasticsearch/v9/typedapi/nodes/info"
	nodes_reload_secure_settings "github.com/elastic/go-elasticsearch/v9/typedapi/nodes/reloadsecuresettings"
	nodes_stats "github.com/elastic/go-elasticsearch/v9/typedapi/nodes/stats"
	nodes_usage "github.com/elastic/go-elasticsearch/v9/typedapi/nodes/usage"
	profiling_flamegraph "github.com/elastic/go-elasticsearch/v9/typedapi/profiling/flamegraph"
	profiling_stacktraces "github.com/elastic/go-elasticsearch/v9/typedapi/profiling/stacktraces"
	profiling_status "github.com/elastic/go-elasticsearch/v9/typedapi/profiling/status"
	profiling_topn_functions "github.com/elastic/go-elasticsearch/v9/typedapi/profiling/topnfunctions"
	project_create_many_routing "github.com/elastic/go-elasticsearch/v9/typedapi/project/createmanyrouting"
	project_create_routing "github.com/elastic/go-elasticsearch/v9/typedapi/project/createrouting"
	project_delete_routing "github.com/elastic/go-elasticsearch/v9/typedapi/project/deleterouting"
	project_get_many_routing "github.com/elastic/go-elasticsearch/v9/typedapi/project/getmanyrouting"
	project_get_routing "github.com/elastic/go-elasticsearch/v9/typedapi/project/getrouting"
	project_tags "github.com/elastic/go-elasticsearch/v9/typedapi/project/tags"
	query_rules_delete_rule "github.com/elastic/go-elasticsearch/v9/typedapi/queryrules/deleterule"
	query_rules_delete_ruleset "github.com/elastic/go-elasticsearch/v9/typedapi/queryrules/deleteruleset"
	query_rules_get_rule "github.com/elastic/go-elasticsearch/v9/typedapi/queryrules/getrule"
	query_rules_get_ruleset "github.com/elastic/go-elasticsearch/v9/typedapi/queryrules/getruleset"
	query_rules_list_rulesets "github.com/elastic/go-elasticsearch/v9/typedapi/queryrules/listrulesets"
	query_rules_put_rule "github.com/elastic/go-elasticsearch/v9/typedapi/queryrules/putrule"
	query_rules_put_ruleset "github.com/elastic/go-elasticsearch/v9/typedapi/queryrules/putruleset"
	query_rules_test "github.com/elastic/go-elasticsearch/v9/typedapi/queryrules/test"
	rollup_delete_job "github.com/elastic/go-elasticsearch/v9/typedapi/rollup/deletejob"
	rollup_get_jobs "github.com/elastic/go-elasticsearch/v9/typedapi/rollup/getjobs"
	rollup_get_rollup_caps "github.com/elastic/go-elasticsearch/v9/typedapi/rollup/getrollupcaps"
	rollup_get_rollup_index_caps "github.com/elastic/go-elasticsearch/v9/typedapi/rollup/getrollupindexcaps"
	rollup_put_job "github.com/elastic/go-elasticsearch/v9/typedapi/rollup/putjob"
	rollup_rollup_search "github.com/elastic/go-elasticsearch/v9/typedapi/rollup/rollupsearch"
	rollup_start_job "github.com/elastic/go-elasticsearch/v9/typedapi/rollup/startjob"
	rollup_stop_job "github.com/elastic/go-elasticsearch/v9/typedapi/rollup/stopjob"
	searchable_snapshots_cache_stats "github.com/elastic/go-elasticsearch/v9/typedapi/searchablesnapshots/cachestats"
	searchable_snapshots_clear_cache "github.com/elastic/go-elasticsearch/v9/typedapi/searchablesnapshots/clearcache"
	searchable_snapshots_mount "github.com/elastic/go-elasticsearch/v9/typedapi/searchablesnapshots/mount"
	searchable_snapshots_stats "github.com/elastic/go-elasticsearch/v9/typedapi/searchablesnapshots/stats"
	search_application_delete "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/delete"
	search_application_delete_behavioral_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/deletebehavioralanalytics"
	search_application_get "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/get"
	search_application_get_behavioral_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/getbehavioralanalytics"
	search_application_list "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/list"
	search_application_post_behavioral_analytics_event "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/postbehavioralanalyticsevent"
	search_application_put "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/put"
	search_application_put_behavioral_analytics "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/putbehavioralanalytics"
	search_application_render_query "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/renderquery"
	search_application_search "github.com/elastic/go-elasticsearch/v9/typedapi/searchapplication/search"
	security_activate_user_profile "github.com/elastic/go-elasticsearch/v9/typedapi/security/activateuserprofile"
	security_authenticate "github.com/elastic/go-elasticsearch/v9/typedapi/security/authenticate"
	security_bulk_delete_role "github.com/elastic/go-elasticsearch/v9/typedapi/security/bulkdeleterole"
	security_bulk_put_role "github.com/elastic/go-elasticsearch/v9/typedapi/security/bulkputrole"
	security_bulk_update_api_keys "github.com/elastic/go-elasticsearch/v9/typedapi/security/bulkupdateapikeys"
	security_change_password "github.com/elastic/go-elasticsearch/v9/typedapi/security/changepassword"
	security_clear_api_key_cache "github.com/elastic/go-elasticsearch/v9/typedapi/security/clearapikeycache"
	security_clear_cached_privileges "github.com/elastic/go-elasticsearch/v9/typedapi/security/clearcachedprivileges"
	security_clear_cached_realms "github.com/elastic/go-elasticsearch/v9/typedapi/security/clearcachedrealms"
	security_clear_cached_roles "github.com/elastic/go-elasticsearch/v9/typedapi/security/clearcachedroles"
	security_clear_cached_service_tokens "github.com/elastic/go-elasticsearch/v9/typedapi/security/clearcachedservicetokens"
	security_clone_api_key "github.com/elastic/go-elasticsearch/v9/typedapi/security/cloneapikey"
	security_create_api_key "github.com/elastic/go-elasticsearch/v9/typedapi/security/createapikey"
	security_create_cross_cluster_api_key "github.com/elastic/go-elasticsearch/v9/typedapi/security/createcrossclusterapikey"
	security_create_service_token "github.com/elastic/go-elasticsearch/v9/typedapi/security/createservicetoken"
	security_delegate_pki "github.com/elastic/go-elasticsearch/v9/typedapi/security/delegatepki"
	security_delete_privileges "github.com/elastic/go-elasticsearch/v9/typedapi/security/deleteprivileges"
	security_delete_role "github.com/elastic/go-elasticsearch/v9/typedapi/security/deleterole"
	security_delete_role_mapping "github.com/elastic/go-elasticsearch/v9/typedapi/security/deleterolemapping"
	security_delete_service_token "github.com/elastic/go-elasticsearch/v9/typedapi/security/deleteservicetoken"
	security_delete_user "github.com/elastic/go-elasticsearch/v9/typedapi/security/deleteuser"
	security_disable_user "github.com/elastic/go-elasticsearch/v9/typedapi/security/disableuser"
	security_disable_user_profile "github.com/elastic/go-elasticsearch/v9/typedapi/security/disableuserprofile"
	security_enable_user "github.com/elastic/go-elasticsearch/v9/typedapi/security/enableuser"
	security_enable_user_profile "github.com/elastic/go-elasticsearch/v9/typedapi/security/enableuserprofile"
	security_enroll_kibana "github.com/elastic/go-elasticsearch/v9/typedapi/security/enrollkibana"
	security_enroll_node "github.com/elastic/go-elasticsearch/v9/typedapi/security/enrollnode"
	security_get_api_key "github.com/elastic/go-elasticsearch/v9/typedapi/security/getapikey"
	security_get_builtin_privileges "github.com/elastic/go-elasticsearch/v9/typedapi/security/getbuiltinprivileges"
	security_get_privileges "github.com/elastic/go-elasticsearch/v9/typedapi/security/getprivileges"
	security_get_role "github.com/elastic/go-elasticsearch/v9/typedapi/security/getrole"
	security_get_role_mapping "github.com/elastic/go-elasticsearch/v9/typedapi/security/getrolemapping"
	security_get_service_accounts "github.com/elastic/go-elasticsearch/v9/typedapi/security/getserviceaccounts"
	security_get_service_credentials "github.com/elastic/go-elasticsearch/v9/typedapi/security/getservicecredentials"
	security_get_settings "github.com/elastic/go-elasticsearch/v9/typedapi/security/getsettings"
	security_get_stats "github.com/elastic/go-elasticsearch/v9/typedapi/security/getstats"
	security_get_token "github.com/elastic/go-elasticsearch/v9/typedapi/security/gettoken"
	security_get_user "github.com/elastic/go-elasticsearch/v9/typedapi/security/getuser"
	security_get_user_privileges "github.com/elastic/go-elasticsearch/v9/typedapi/security/getuserprivileges"
	security_get_user_profile "github.com/elastic/go-elasticsearch/v9/typedapi/security/getuserprofile"
	security_grant_api_key "github.com/elastic/go-elasticsearch/v9/typedapi/security/grantapikey"
	security_has_privileges "github.com/elastic/go-elasticsearch/v9/typedapi/security/hasprivileges"
	security_has_privileges_user_profile "github.com/elastic/go-elasticsearch/v9/typedapi/security/hasprivilegesuserprofile"
	security_invalidate_api_key "github.com/elastic/go-elasticsearch/v9/typedapi/security/invalidateapikey"
	security_invalidate_token "github.com/elastic/go-elasticsearch/v9/typedapi/security/invalidatetoken"
	security_oidc_authenticate "github.com/elastic/go-elasticsearch/v9/typedapi/security/oidcauthenticate"
	security_oidc_logout "github.com/elastic/go-elasticsearch/v9/typedapi/security/oidclogout"
	security_oidc_prepare_authentication "github.com/elastic/go-elasticsearch/v9/typedapi/security/oidcprepareauthentication"
	security_put_privileges "github.com/elastic/go-elasticsearch/v9/typedapi/security/putprivileges"
	security_put_role "github.com/elastic/go-elasticsearch/v9/typedapi/security/putrole"
	security_put_role_mapping "github.com/elastic/go-elasticsearch/v9/typedapi/security/putrolemapping"
	security_put_user "github.com/elastic/go-elasticsearch/v9/typedapi/security/putuser"
	security_query_api_keys "github.com/elastic/go-elasticsearch/v9/typedapi/security/queryapikeys"
	security_query_role "github.com/elastic/go-elasticsearch/v9/typedapi/security/queryrole"
	security_query_user "github.com/elastic/go-elasticsearch/v9/typedapi/security/queryuser"
	security_saml_authenticate "github.com/elastic/go-elasticsearch/v9/typedapi/security/samlauthenticate"
	security_saml_complete_logout "github.com/elastic/go-elasticsearch/v9/typedapi/security/samlcompletelogout"
	security_saml_invalidate "github.com/elastic/go-elasticsearch/v9/typedapi/security/samlinvalidate"
	security_saml_logout "github.com/elastic/go-elasticsearch/v9/typedapi/security/samllogout"
	security_saml_prepare_authentication "github.com/elastic/go-elasticsearch/v9/typedapi/security/samlprepareauthentication"
	security_saml_service_provider_metadata "github.com/elastic/go-elasticsearch/v9/typedapi/security/samlserviceprovidermetadata"
	security_suggest_user_profiles "github.com/elastic/go-elasticsearch/v9/typedapi/security/suggestuserprofiles"
	security_update_api_key "github.com/elastic/go-elasticsearch/v9/typedapi/security/updateapikey"
	security_update_cross_cluster_api_key "github.com/elastic/go-elasticsearch/v9/typedapi/security/updatecrossclusterapikey"
	security_update_settings "github.com/elastic/go-elasticsearch/v9/typedapi/security/updatesettings"
	security_update_user_profile_data "github.com/elastic/go-elasticsearch/v9/typedapi/security/updateuserprofiledata"
	shutdown_delete_node "github.com/elastic/go-elasticsearch/v9/typedapi/shutdown/deletenode"
	shutdown_get_node "github.com/elastic/go-elasticsearch/v9/typedapi/shutdown/getnode"
	shutdown_put_node "github.com/elastic/go-elasticsearch/v9/typedapi/shutdown/putnode"
	simulate_ingest "github.com/elastic/go-elasticsearch/v9/typedapi/simulate/ingest"
	slm_delete_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/slm/deletelifecycle"
	slm_execute_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/slm/executelifecycle"
	slm_execute_retention "github.com/elastic/go-elasticsearch/v9/typedapi/slm/executeretention"
	slm_get_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/slm/getlifecycle"
	slm_get_stats "github.com/elastic/go-elasticsearch/v9/typedapi/slm/getstats"
	slm_get_status "github.com/elastic/go-elasticsearch/v9/typedapi/slm/getstatus"
	slm_put_lifecycle "github.com/elastic/go-elasticsearch/v9/typedapi/slm/putlifecycle"
	slm_start "github.com/elastic/go-elasticsearch/v9/typedapi/slm/start"
	slm_stop "github.com/elastic/go-elasticsearch/v9/typedapi/slm/stop"
	snapshot_cleanup_repository "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/cleanuprepository"
	snapshot_clone "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/clone"
	snapshot_create "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/create"
	snapshot_create_repository "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/createrepository"
	snapshot_delete "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/delete"
	snapshot_delete_repository "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/deleterepository"
	snapshot_get "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/get"
	snapshot_get_repository "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/getrepository"
	snapshot_repository_analyze "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/repositoryanalyze"
	snapshot_repository_verify_integrity "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/repositoryverifyintegrity"
	snapshot_restore "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/restore"
	snapshot_status "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/status"
	snapshot_verify_repository "github.com/elastic/go-elasticsearch/v9/typedapi/snapshot/verifyrepository"
	sql_clear_cursor "github.com/elastic/go-elasticsearch/v9/typedapi/sql/clearcursor"
	sql_delete_async "github.com/elastic/go-elasticsearch/v9/typedapi/sql/deleteasync"
	sql_get_async "github.com/elastic/go-elasticsearch/v9/typedapi/sql/getasync"
	sql_get_async_status "github.com/elastic/go-elasticsearch/v9/typedapi/sql/getasyncstatus"
	sql_query "github.com/elastic/go-elasticsearch/v9/typedapi/sql/query"
	sql_translate "github.com/elastic/go-elasticsearch/v9/typedapi/sql/translate"
	ssl_certificates "github.com/elastic/go-elasticsearch/v9/typedapi/ssl/certificates"
	streams_logs_disable "github.com/elastic/go-elasticsearch/v9/typedapi/streams/logsdisable"
	streams_logs_enable "github.com/elastic/go-elasticsearch/v9/typedapi/streams/logsenable"
	streams_status "github.com/elastic/go-elasticsearch/v9/typedapi/streams/status"
	synonyms_delete_synonym "github.com/elastic/go-elasticsearch/v9/typedapi/synonyms/deletesynonym"
	synonyms_delete_synonym_rule "github.com/elastic/go-elasticsearch/v9/typedapi/synonyms/deletesynonymrule"
	synonyms_get_synonym "github.com/elastic/go-elasticsearch/v9/typedapi/synonyms/getsynonym"
	synonyms_get_synonym_rule "github.com/elastic/go-elasticsearch/v9/typedapi/synonyms/getsynonymrule"
	synonyms_get_synonyms_sets "github.com/elastic/go-elasticsearch/v9/typedapi/synonyms/getsynonymssets"
	synonyms_put_synonym "github.com/elastic/go-elasticsearch/v9/typedapi/synonyms/putsynonym"
	synonyms_put_synonym_rule "github.com/elastic/go-elasticsearch/v9/typedapi/synonyms/putsynonymrule"
	tasks_cancel "github.com/elastic/go-elasticsearch/v9/typedapi/tasks/cancel"
	tasks_get "github.com/elastic/go-elasticsearch/v9/typedapi/tasks/get"
	tasks_list "github.com/elastic/go-elasticsearch/v9/typedapi/tasks/list"
	text_structure_find_field_structure "github.com/elastic/go-elasticsearch/v9/typedapi/textstructure/findfieldstructure"
	text_structure_find_message_structure "github.com/elastic/go-elasticsearch/v9/typedapi/textstructure/findmessagestructure"
	text_structure_find_structure "github.com/elastic/go-elasticsearch/v9/typedapi/textstructure/findstructure"
	text_structure_test_grok_pattern "github.com/elastic/go-elasticsearch/v9/typedapi/textstructure/testgrokpattern"
	transform_delete_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/deletetransform"
	transform_get_node_stats "github.com/elastic/go-elasticsearch/v9/typedapi/transform/getnodestats"
	transform_get_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/gettransform"
	transform_get_transform_stats "github.com/elastic/go-elasticsearch/v9/typedapi/transform/gettransformstats"
	transform_preview_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/previewtransform"
	transform_put_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/puttransform"
	transform_reset_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/resettransform"
	transform_schedule_now_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/schedulenowtransform"
	transform_set_upgrade_mode "github.com/elastic/go-elasticsearch/v9/typedapi/transform/setupgrademode"
	transform_start_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/starttransform"
	transform_stop_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/stoptransform"
	transform_update_transform "github.com/elastic/go-elasticsearch/v9/typedapi/transform/updatetransform"
	transform_upgrade_transforms "github.com/elastic/go-elasticsearch/v9/typedapi/transform/upgradetransforms"
	watcher_ack_watch "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/ackwatch"
	watcher_activate_watch "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/activatewatch"
	watcher_deactivate_watch "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/deactivatewatch"
	watcher_delete_watch "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/deletewatch"
	watcher_execute_watch "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/executewatch"
	watcher_get_settings "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/getsettings"
	watcher_get_watch "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/getwatch"
	watcher_put_watch "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/putwatch"
	watcher_query_watches "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/querywatches"
	watcher_start "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/start"
	watcher_stats "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/stats"
	watcher_stop "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/stop"
	watcher_update_settings "github.com/elastic/go-elasticsearch/v9/typedapi/watcher/updatesettings"
	xpack_info "github.com/elastic/go-elasticsearch/v9/typedapi/xpack/info"
	xpack_usage "github.com/elastic/go-elasticsearch/v9/typedapi/xpack/usage"
)

type AsyncSearch struct {
	Delete async_search_delete.NewDelete

	Get async_search_get.NewGet

	Status async_search_status.NewStatus

	Submit async_search_submit.NewSubmit
}

type Autoscaling struct {
	DeleteAutoscalingPolicy autoscaling_delete_autoscaling_policy.NewDeleteAutoscalingPolicy

	GetAutoscalingCapacity autoscaling_get_autoscaling_capacity.NewGetAutoscalingCapacity

	GetAutoscalingPolicy autoscaling_get_autoscaling_policy.NewGetAutoscalingPolicy

	PutAutoscalingPolicy autoscaling_put_autoscaling_policy.NewPutAutoscalingPolicy
}

type Cat struct {
	Aliases cat_aliases.NewAliases

	Allocation cat_allocation.NewAllocation

	CircuitBreaker cat_circuit_breaker.NewCircuitBreaker

	ComponentTemplates cat_component_templates.NewComponentTemplates

	Count cat_count.NewCount

	Fielddata cat_fielddata.NewFielddata

	Health cat_health.NewHealth

	Help cat_help.NewHelp

	Indices cat_indices.NewIndices

	Master cat_master.NewMaster

	MlDataFrameAnalytics cat_ml_data_frame_analytics.NewMlDataFrameAnalytics

	MlDatafeeds cat_ml_datafeeds.NewMlDatafeeds

	MlJobs cat_ml_jobs.NewMlJobs

	MlTrainedModels cat_ml_trained_models.NewMlTrainedModels

	Nodeattrs cat_nodeattrs.NewNodeattrs

	Nodes cat_nodes.NewNodes

	PendingTasks cat_pending_tasks.NewPendingTasks

	Plugins cat_plugins.NewPlugins

	Recovery cat_recovery.NewRecovery

	Repositories cat_repositories.NewRepositories

	Segments cat_segments.NewSegments

	Shards cat_shards.NewShards

	Snapshots cat_snapshots.NewSnapshots

	Tasks cat_tasks.NewTasks

	Templates cat_templates.NewTemplates

	ThreadPool cat_thread_pool.NewThreadPool

	Transforms cat_transforms.NewTransforms
}

type Ccr struct {
	DeleteAutoFollowPattern ccr_delete_auto_follow_pattern.NewDeleteAutoFollowPattern

	Follow ccr_follow.NewFollow

	FollowInfo ccr_follow_info.NewFollowInfo

	FollowStats ccr_follow_stats.NewFollowStats

	ForgetFollower ccr_forget_follower.NewForgetFollower

	GetAutoFollowPattern ccr_get_auto_follow_pattern.NewGetAutoFollowPattern

	PauseAutoFollowPattern ccr_pause_auto_follow_pattern.NewPauseAutoFollowPattern

	PauseFollow ccr_pause_follow.NewPauseFollow

	PutAutoFollowPattern ccr_put_auto_follow_pattern.NewPutAutoFollowPattern

	ResumeAutoFollowPattern ccr_resume_auto_follow_pattern.NewResumeAutoFollowPattern

	ResumeFollow ccr_resume_follow.NewResumeFollow

	Stats ccr_stats.NewStats

	Unfollow ccr_unfollow.NewUnfollow
}

type Cluster struct {
	AllocationExplain cluster_allocation_explain.NewAllocationExplain

	DeleteComponentTemplate cluster_delete_component_template.NewDeleteComponentTemplate

	DeleteVotingConfigExclusions cluster_delete_voting_config_exclusions.NewDeleteVotingConfigExclusions

	ExistsComponentTemplate cluster_exists_component_template.NewExistsComponentTemplate

	GetComponentTemplate cluster_get_component_template.NewGetComponentTemplate

	GetSettings cluster_get_settings.NewGetSettings

	Health cluster_health.NewHealth

	Info cluster_info.NewInfo

	PendingTasks cluster_pending_tasks.NewPendingTasks

	PostVotingConfigExclusions cluster_post_voting_config_exclusions.NewPostVotingConfigExclusions

	PutComponentTemplate cluster_put_component_template.NewPutComponentTemplate

	PutSettings cluster_put_settings.NewPutSettings

	RemoteInfo cluster_remote_info.NewRemoteInfo

	Reroute cluster_reroute.NewReroute

	State cluster_state.NewState

	Stats cluster_stats.NewStats
}

type Connector struct {
	CheckIn connector_check_in.NewCheckIn

	Delete connector_delete.NewDelete

	Get connector_get.NewGet

	LastSync connector_last_sync.NewLastSync

	List connector_list.NewList

	Post connector_post.NewPost

	Put connector_put.NewPut

	SecretDelete connector_secret_delete.NewSecretDelete

	SecretGet connector_secret_get.NewSecretGet

	SecretPost connector_secret_post.NewSecretPost

	SecretPut connector_secret_put.NewSecretPut

	SyncJobCancel connector_sync_job_cancel.NewSyncJobCancel

	SyncJobCheckIn connector_sync_job_check_in.NewSyncJobCheckIn

	SyncJobClaim connector_sync_job_claim.NewSyncJobClaim

	SyncJobDelete connector_sync_job_delete.NewSyncJobDelete

	SyncJobError connector_sync_job_error.NewSyncJobError

	SyncJobGet connector_sync_job_get.NewSyncJobGet

	SyncJobList connector_sync_job_list.NewSyncJobList

	SyncJobPost connector_sync_job_post.NewSyncJobPost

	SyncJobUpdateStats connector_sync_job_update_stats.NewSyncJobUpdateStats

	UpdateActiveFiltering connector_update_active_filtering.NewUpdateActiveFiltering

	UpdateApiKeyId connector_update_api_key_id.NewUpdateApiKeyId

	UpdateConfiguration connector_update_configuration.NewUpdateConfiguration

	UpdateError connector_update_error.NewUpdateError

	UpdateFeatures connector_update_features.NewUpdateFeatures

	UpdateFiltering connector_update_filtering.NewUpdateFiltering

	UpdateFilteringValidation connector_update_filtering_validation.NewUpdateFilteringValidation

	UpdateIndexName connector_update_index_name.NewUpdateIndexName

	UpdateName connector_update_name.NewUpdateName

	UpdateNative connector_update_native.NewUpdateNative

	UpdatePipeline connector_update_pipeline.NewUpdatePipeline

	UpdateScheduling connector_update_scheduling.NewUpdateScheduling

	UpdateServiceType connector_update_service_type.NewUpdateServiceType

	UpdateStatus connector_update_status.NewUpdateStatus
}

type Core struct {
	Bulk core_bulk.NewBulk

	CancelReindex core_cancel_reindex.NewCancelReindex

	Capabilities core_capabilities.NewCapabilities

	ClearScroll core_clear_scroll.NewClearScroll

	ClosePointInTime core_close_point_in_time.NewClosePointInTime

	Count core_count.NewCount

	Create core_create.NewCreate

	Delete core_delete.NewDelete

	DeleteByQuery core_delete_by_query.NewDeleteByQuery

	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle

	DeleteScript core_delete_script.NewDeleteScript

	Exists core_exists.NewExists

	ExistsSource core_exists_source.NewExistsSource

	Explain core_explain.NewExplain

	FieldCaps core_field_caps.NewFieldCaps

	Get core_get.NewGet

	GetReindex core_get_reindex.NewGetReindex

	GetScript core_get_script.NewGetScript

	GetScriptContext core_get_script_context.NewGetScriptContext

	GetScriptLanguages core_get_script_languages.NewGetScriptLanguages

	GetSource core_get_source.NewGetSource

	HealthReport core_health_report.NewHealthReport

	Index core_index.NewIndex

	Info core_info.NewInfo

	KnnSearch core_knn_search.NewKnnSearch

	ListReindex core_list_reindex.NewListReindex

	Mget core_mget.NewMget

	Msearch core_msearch.NewMsearch

	MsearchTemplate core_msearch_template.NewMsearchTemplate

	Mtermvectors core_mtermvectors.NewMtermvectors

	OpenPointInTime core_open_point_in_time.NewOpenPointInTime

	Ping core_ping.NewPing

	PutScript core_put_script.NewPutScript

	RankEval core_rank_eval.NewRankEval

	Reindex core_reindex.NewReindex

	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle

	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate

	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute

	Scroll core_scroll.NewScroll

	Search core_search.NewSearch

	SearchMvt core_search_mvt.NewSearchMvt

	SearchShards core_search_shards.NewSearchShards

	SearchTemplate core_search_template.NewSearchTemplate

	TermsEnum core_terms_enum.NewTermsEnum

	Termvectors core_termvectors.NewTermvectors

	Update core_update.NewUpdate

	UpdateByQuery core_update_by_query.NewUpdateByQuery

	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

type DanglingIndices struct {
	DeleteDanglingIndex dangling_indices_delete_dangling_index.NewDeleteDanglingIndex

	ImportDanglingIndex dangling_indices_import_dangling_index.NewImportDanglingIndex

	ListDanglingIndices dangling_indices_list_dangling_indices.NewListDanglingIndices
}

type Enrich struct {
	DeletePolicy enrich_delete_policy.NewDeletePolicy

	ExecutePolicy enrich_execute_policy.NewExecutePolicy

	GetPolicy enrich_get_policy.NewGetPolicy

	PutPolicy enrich_put_policy.NewPutPolicy

	Stats enrich_stats.NewStats
}

type Eql struct {
	Delete eql_delete.NewDelete

	Get eql_get.NewGet

	GetStatus eql_get_status.NewGetStatus

	Search eql_search.NewSearch
}

type Esql struct {
	AsyncQuery esql_async_query.NewAsyncQuery

	AsyncQueryDelete esql_async_query_delete.NewAsyncQueryDelete

	AsyncQueryGet esql_async_query_get.NewAsyncQueryGet

	AsyncQueryStop esql_async_query_stop.NewAsyncQueryStop

	DeleteView esql_delete_view.NewDeleteView

	GetQuery esql_get_query.NewGetQuery

	GetView esql_get_view.NewGetView

	ListQueries esql_list_queries.NewListQueries

	PutView esql_put_view.NewPutView

	Query esql_query.NewQuery
}

type Features struct {
	GetFeatures features_get_features.NewGetFeatures

	ResetFeatures features_reset_features.NewResetFeatures
}

type Fleet struct {
	DeleteSecret fleet_delete_secret.NewDeleteSecret

	GetSecret fleet_get_secret.NewGetSecret

	GlobalCheckpoints fleet_global_checkpoints.NewGlobalCheckpoints

	Msearch fleet_msearch.NewMsearch

	PostSecret fleet_post_secret.NewPostSecret

	Search fleet_search.NewSearch
}

type Graph struct {
	Explore graph_explore.NewExplore
}

type Ilm struct {
	DeleteLifecycle ilm_delete_lifecycle.NewDeleteLifecycle

	ExplainLifecycle ilm_explain_lifecycle.NewExplainLifecycle

	GetLifecycle ilm_get_lifecycle.NewGetLifecycle

	GetStatus ilm_get_status.NewGetStatus

	MigrateToDataTiers ilm_migrate_to_data_tiers.NewMigrateToDataTiers

	MoveToStep ilm_move_to_step.NewMoveToStep

	PutLifecycle ilm_put_lifecycle.NewPutLifecycle

	RemovePolicy ilm_remove_policy.NewRemovePolicy

	Retry ilm_retry.NewRetry

	Start ilm_start.NewStart

	Stop ilm_stop.NewStop
}

type Indices struct {
	AddBlock indices_add_block.NewAddBlock

	Analyze indices_analyze.NewAnalyze

	CancelMigrateReindex indices_cancel_migrate_reindex.NewCancelMigrateReindex

	ClearCache indices_clear_cache.NewClearCache

	Clone indices_clone.NewClone

	Close indices_close.NewClose

	Create indices_create.NewCreate

	CreateDataStream indices_create_data_stream.NewCreateDataStream

	CreateFrom indices_create_from.NewCreateFrom

	DataStreamsStats indices_data_streams_stats.NewDataStreamsStats

	Delete indices_delete.NewDelete

	DeleteAlias indices_delete_alias.NewDeleteAlias

	DeleteDataLifecycle indices_delete_data_lifecycle.NewDeleteDataLifecycle

	DeleteDataStream indices_delete_data_stream.NewDeleteDataStream

	DeleteDataStreamOptions indices_delete_data_stream_options.NewDeleteDataStreamOptions

	DeleteIndexTemplate indices_delete_index_template.NewDeleteIndexTemplate

	DeleteTemplate indices_delete_template.NewDeleteTemplate

	DiskUsage indices_disk_usage.NewDiskUsage

	Downsample indices_downsample.NewDownsample

	Exists indices_exists.NewExists

	ExistsAlias indices_exists_alias.NewExistsAlias

	ExistsIndexTemplate indices_exists_index_template.NewExistsIndexTemplate

	ExistsTemplate indices_exists_template.NewExistsTemplate

	ExplainDataLifecycle indices_explain_data_lifecycle.NewExplainDataLifecycle

	FieldUsageStats indices_field_usage_stats.NewFieldUsageStats

	Flush indices_flush.NewFlush

	Forcemerge indices_forcemerge.NewForcemerge

	Get indices_get.NewGet

	GetAlias indices_get_alias.NewGetAlias

	GetDataLifecycle indices_get_data_lifecycle.NewGetDataLifecycle

	GetDataLifecycleStats indices_get_data_lifecycle_stats.NewGetDataLifecycleStats

	GetDataStream indices_get_data_stream.NewGetDataStream

	GetDataStreamMappings indices_get_data_stream_mappings.NewGetDataStreamMappings

	GetDataStreamOptions indices_get_data_stream_options.NewGetDataStreamOptions

	GetDataStreamSettings indices_get_data_stream_settings.NewGetDataStreamSettings

	GetFieldMapping indices_get_field_mapping.NewGetFieldMapping

	GetIndexTemplate indices_get_index_template.NewGetIndexTemplate

	GetMapping indices_get_mapping.NewGetMapping

	GetMigrateReindexStatus indices_get_migrate_reindex_status.NewGetMigrateReindexStatus

	GetSettings indices_get_settings.NewGetSettings

	GetTemplate indices_get_template.NewGetTemplate

	MigrateReindex indices_migrate_reindex.NewMigrateReindex

	MigrateToDataStream indices_migrate_to_data_stream.NewMigrateToDataStream

	ModifyDataStream indices_modify_data_stream.NewModifyDataStream

	Open indices_open.NewOpen

	PromoteDataStream indices_promote_data_stream.NewPromoteDataStream

	PutAlias indices_put_alias.NewPutAlias

	PutDataLifecycle indices_put_data_lifecycle.NewPutDataLifecycle

	PutDataStreamMappings indices_put_data_stream_mappings.NewPutDataStreamMappings

	PutDataStreamOptions indices_put_data_stream_options.NewPutDataStreamOptions

	PutDataStreamSettings indices_put_data_stream_settings.NewPutDataStreamSettings

	PutIndexTemplate indices_put_index_template.NewPutIndexTemplate

	PutMapping indices_put_mapping.NewPutMapping

	PutSettings indices_put_settings.NewPutSettings

	PutTemplate indices_put_template.NewPutTemplate

	Recovery indices_recovery.NewRecovery

	Refresh indices_refresh.NewRefresh

	ReloadSearchAnalyzers indices_reload_search_analyzers.NewReloadSearchAnalyzers

	RemoveBlock indices_remove_block.NewRemoveBlock

	ResolveCluster indices_resolve_cluster.NewResolveCluster

	ResolveIndex indices_resolve_index.NewResolveIndex

	Rollover indices_rollover.NewRollover

	Segments indices_segments.NewSegments

	ShardStores indices_shard_stores.NewShardStores

	Shrink indices_shrink.NewShrink

	SimulateIndexTemplate indices_simulate_index_template.NewSimulateIndexTemplate

	SimulateTemplate indices_simulate_template.NewSimulateTemplate

	Split indices_split.NewSplit

	Stats indices_stats.NewStats

	UpdateAliases indices_update_aliases.NewUpdateAliases

	ValidateQuery indices_validate_query.NewValidateQuery
}

type Inference struct {
	ChatCompletionUnified inference_chat_completion_unified.NewChatCompletionUnified

	Completion inference_completion.NewCompletion

	Delete inference_delete.NewDelete

	Embedding inference_embedding.NewEmbedding

	Get inference_get.NewGet

	Inference inference_inference.NewInference

	Put inference_put.NewPut

	PutAi21 inference_put_ai21.NewPutAi21

	PutAlibabacloud inference_put_alibabacloud.NewPutAlibabacloud

	PutAmazonbedrock inference_put_amazonbedrock.NewPutAmazonbedrock

	PutAmazonsagemaker inference_put_amazonsagemaker.NewPutAmazonsagemaker

	PutAnthropic inference_put_anthropic.NewPutAnthropic

	PutAzureaistudio inference_put_azureaistudio.NewPutAzureaistudio

	PutAzureopenai inference_put_azureopenai.NewPutAzureopenai

	PutCohere inference_put_cohere.NewPutCohere

	PutContextualai inference_put_contextualai.NewPutContextualai

	PutCustom inference_put_custom.NewPutCustom

	PutDeepseek inference_put_deepseek.NewPutDeepseek

	PutElasticsearch inference_put_elasticsearch.NewPutElasticsearch

	PutElser inference_put_elser.NewPutElser

	PutFireworksai inference_put_fireworksai.NewPutFireworksai

	PutGoogleaistudio inference_put_googleaistudio.NewPutGoogleaistudio

	PutGooglevertexai inference_put_googlevertexai.NewPutGooglevertexai

	PutGroq inference_put_groq.NewPutGroq

	PutHuggingFace inference_put_hugging_face.NewPutHuggingFace

	PutJinaai inference_put_jinaai.NewPutJinaai

	PutLlama inference_put_llama.NewPutLlama

	PutMistral inference_put_mistral.NewPutMistral

	PutNvidia inference_put_nvidia.NewPutNvidia

	PutOpenai inference_put_openai.NewPutOpenai

	PutOpenshiftAi inference_put_openshift_ai.NewPutOpenshiftAi

	PutVoyageai inference_put_voyageai.NewPutVoyageai

	PutWatsonx inference_put_watsonx.NewPutWatsonx

	Rerank inference_rerank.NewRerank

	SparseEmbedding inference_sparse_embedding.NewSparseEmbedding

	StreamCompletion inference_stream_completion.NewStreamCompletion

	TextEmbedding inference_text_embedding.NewTextEmbedding

	Update inference_update.NewUpdate
}

type Ingest struct {
	DeleteGeoipDatabase ingest_delete_geoip_database.NewDeleteGeoipDatabase

	DeleteIpLocationDatabase ingest_delete_ip_location_database.NewDeleteIpLocationDatabase

	DeletePipeline ingest_delete_pipeline.NewDeletePipeline

	GeoIpStats ingest_geo_ip_stats.NewGeoIpStats

	GetGeoipDatabase ingest_get_geoip_database.NewGetGeoipDatabase

	GetIpLocationDatabase ingest_get_ip_location_database.NewGetIpLocationDatabase

	GetPipeline ingest_get_pipeline.NewGetPipeline

	ProcessorGrok ingest_processor_grok.NewProcessorGrok

	PutGeoipDatabase ingest_put_geoip_database.NewPutGeoipDatabase

	PutIpLocationDatabase ingest_put_ip_location_database.NewPutIpLocationDatabase

	PutPipeline ingest_put_pipeline.NewPutPipeline

	Simulate ingest_simulate.NewSimulate
}

type License struct {
	Delete license_delete.NewDelete

	Get license_get.NewGet

	GetBasicStatus license_get_basic_status.NewGetBasicStatus

	GetTrialStatus license_get_trial_status.NewGetTrialStatus

	Post license_post.NewPost

	PostStartBasic license_post_start_basic.NewPostStartBasic

	PostStartTrial license_post_start_trial.NewPostStartTrial
}

type Logstash struct {
	DeletePipeline logstash_delete_pipeline.NewDeletePipeline

	GetPipeline logstash_get_pipeline.NewGetPipeline

	PutPipeline logstash_put_pipeline.NewPutPipeline
}

type Migration struct {
	Deprecations migration_deprecations.NewDeprecations

	GetFeatureUpgradeStatus migration_get_feature_upgrade_status.NewGetFeatureUpgradeStatus

	PostFeatureUpgrade migration_post_feature_upgrade.NewPostFeatureUpgrade
}

type Ml struct {
	ClearTrainedModelDeploymentCache ml_clear_trained_model_deployment_cache.NewClearTrainedModelDeploymentCache

	CloseJob ml_close_job.NewCloseJob

	DeleteCalendar ml_delete_calendar.NewDeleteCalendar

	DeleteCalendarEvent ml_delete_calendar_event.NewDeleteCalendarEvent

	DeleteCalendarJob ml_delete_calendar_job.NewDeleteCalendarJob

	DeleteDataFrameAnalytics ml_delete_data_frame_analytics.NewDeleteDataFrameAnalytics

	DeleteDatafeed ml_delete_datafeed.NewDeleteDatafeed

	DeleteExpiredData ml_delete_expired_data.NewDeleteExpiredData

	DeleteFilter ml_delete_filter.NewDeleteFilter

	DeleteForecast ml_delete_forecast.NewDeleteForecast

	DeleteJob ml_delete_job.NewDeleteJob

	DeleteModelSnapshot ml_delete_model_snapshot.NewDeleteModelSnapshot

	DeleteTrainedModel ml_delete_trained_model.NewDeleteTrainedModel

	DeleteTrainedModelAlias ml_delete_trained_model_alias.NewDeleteTrainedModelAlias

	EstimateModelMemory ml_estimate_model_memory.NewEstimateModelMemory

	EvaluateDataFrame ml_evaluate_data_frame.NewEvaluateDataFrame

	ExplainDataFrameAnalytics ml_explain_data_frame_analytics.NewExplainDataFrameAnalytics

	FlushJob ml_flush_job.NewFlushJob

	Forecast ml_forecast.NewForecast

	GetBuckets ml_get_buckets.NewGetBuckets

	GetCalendarEvents ml_get_calendar_events.NewGetCalendarEvents

	GetCalendars ml_get_calendars.NewGetCalendars

	GetCategories ml_get_categories.NewGetCategories

	GetDataFrameAnalytics ml_get_data_frame_analytics.NewGetDataFrameAnalytics

	GetDataFrameAnalyticsStats ml_get_data_frame_analytics_stats.NewGetDataFrameAnalyticsStats

	GetDatafeedStats ml_get_datafeed_stats.NewGetDatafeedStats

	GetDatafeeds ml_get_datafeeds.NewGetDatafeeds

	GetFilters ml_get_filters.NewGetFilters

	GetInfluencers ml_get_influencers.NewGetInfluencers

	GetJobStats ml_get_job_stats.NewGetJobStats

	GetJobs ml_get_jobs.NewGetJobs

	GetMemoryStats ml_get_memory_stats.NewGetMemoryStats

	GetModelSnapshotUpgradeStats ml_get_model_snapshot_upgrade_stats.NewGetModelSnapshotUpgradeStats

	GetModelSnapshots ml_get_model_snapshots.NewGetModelSnapshots

	GetOverallBuckets ml_get_overall_buckets.NewGetOverallBuckets

	GetRecords ml_get_records.NewGetRecords

	GetTrainedModels ml_get_trained_models.NewGetTrainedModels

	GetTrainedModelsStats ml_get_trained_models_stats.NewGetTrainedModelsStats

	InferTrainedModel ml_infer_trained_model.NewInferTrainedModel

	Info ml_info.NewInfo

	OpenJob ml_open_job.NewOpenJob

	PostCalendarEvents ml_post_calendar_events.NewPostCalendarEvents

	PostData ml_post_data.NewPostData

	PreviewDataFrameAnalytics ml_preview_data_frame_analytics.NewPreviewDataFrameAnalytics

	PreviewDatafeed ml_preview_datafeed.NewPreviewDatafeed

	PutCalendar ml_put_calendar.NewPutCalendar

	PutCalendarJob ml_put_calendar_job.NewPutCalendarJob

	PutDataFrameAnalytics ml_put_data_frame_analytics.NewPutDataFrameAnalytics

	PutDatafeed ml_put_datafeed.NewPutDatafeed

	PutFilter ml_put_filter.NewPutFilter

	PutJob ml_put_job.NewPutJob

	PutTrainedModel ml_put_trained_model.NewPutTrainedModel

	PutTrainedModelAlias ml_put_trained_model_alias.NewPutTrainedModelAlias

	PutTrainedModelDefinitionPart ml_put_trained_model_definition_part.NewPutTrainedModelDefinitionPart

	PutTrainedModelVocabulary ml_put_trained_model_vocabulary.NewPutTrainedModelVocabulary

	ResetJob ml_reset_job.NewResetJob

	RevertModelSnapshot ml_revert_model_snapshot.NewRevertModelSnapshot

	SetUpgradeMode ml_set_upgrade_mode.NewSetUpgradeMode

	StartDataFrameAnalytics ml_start_data_frame_analytics.NewStartDataFrameAnalytics

	StartDatafeed ml_start_datafeed.NewStartDatafeed

	StartTrainedModelDeployment ml_start_trained_model_deployment.NewStartTrainedModelDeployment

	StopDataFrameAnalytics ml_stop_data_frame_analytics.NewStopDataFrameAnalytics

	StopDatafeed ml_stop_datafeed.NewStopDatafeed

	StopTrainedModelDeployment ml_stop_trained_model_deployment.NewStopTrainedModelDeployment

	UpdateDataFrameAnalytics ml_update_data_frame_analytics.NewUpdateDataFrameAnalytics

	UpdateDatafeed ml_update_datafeed.NewUpdateDatafeed

	UpdateFilter ml_update_filter.NewUpdateFilter

	UpdateJob ml_update_job.NewUpdateJob

	UpdateModelSnapshot ml_update_model_snapshot.NewUpdateModelSnapshot

	UpdateTrainedModelDeployment ml_update_trained_model_deployment.NewUpdateTrainedModelDeployment

	UpgradeJobSnapshot ml_upgrade_job_snapshot.NewUpgradeJobSnapshot

	Validate ml_validate.NewValidate

	ValidateDetector ml_validate_detector.NewValidateDetector
}

type Monitoring struct {
	Bulk monitoring_bulk.NewBulk
}

type Nodes struct {
	ClearRepositoriesMeteringArchive nodes_clear_repositories_metering_archive.NewClearRepositoriesMeteringArchive

	GetRepositoriesMeteringInfo nodes_get_repositories_metering_info.NewGetRepositoriesMeteringInfo

	HotThreads nodes_hot_threads.NewHotThreads

	Info nodes_info.NewInfo

	ReloadSecureSettings nodes_reload_secure_settings.NewReloadSecureSettings

	Stats nodes_stats.NewStats

	Usage nodes_usage.NewUsage
}

type Profiling struct {
	Flamegraph profiling_flamegraph.NewFlamegraph

	Stacktraces profiling_stacktraces.NewStacktraces

	Status profiling_status.NewStatus

	TopnFunctions profiling_topn_functions.NewTopnFunctions
}

type Project struct {
	CreateManyRouting project_create_many_routing.NewCreateManyRouting

	CreateRouting project_create_routing.NewCreateRouting

	DeleteRouting project_delete_routing.NewDeleteRouting

	GetManyRouting project_get_many_routing.NewGetManyRouting

	GetRouting project_get_routing.NewGetRouting

	Tags project_tags.NewTags
}

type QueryRules struct {
	DeleteRule query_rules_delete_rule.NewDeleteRule

	DeleteRuleset query_rules_delete_ruleset.NewDeleteRuleset

	GetRule query_rules_get_rule.NewGetRule

	GetRuleset query_rules_get_ruleset.NewGetRuleset

	ListRulesets query_rules_list_rulesets.NewListRulesets

	PutRule query_rules_put_rule.NewPutRule

	PutRuleset query_rules_put_ruleset.NewPutRuleset

	Test query_rules_test.NewTest
}

type Rollup struct {
	DeleteJob rollup_delete_job.NewDeleteJob

	GetJobs rollup_get_jobs.NewGetJobs

	GetRollupCaps rollup_get_rollup_caps.NewGetRollupCaps

	GetRollupIndexCaps rollup_get_rollup_index_caps.NewGetRollupIndexCaps

	PutJob rollup_put_job.NewPutJob

	RollupSearch rollup_rollup_search.NewRollupSearch

	StartJob rollup_start_job.NewStartJob

	StopJob rollup_stop_job.NewStopJob
}

type SearchApplication struct {
	Delete search_application_delete.NewDelete

	DeleteBehavioralAnalytics search_application_delete_behavioral_analytics.NewDeleteBehavioralAnalytics

	Get search_application_get.NewGet

	GetBehavioralAnalytics search_application_get_behavioral_analytics.NewGetBehavioralAnalytics

	List search_application_list.NewList

	PostBehavioralAnalyticsEvent search_application_post_behavioral_analytics_event.NewPostBehavioralAnalyticsEvent

	Put search_application_put.NewPut

	PutBehavioralAnalytics search_application_put_behavioral_analytics.NewPutBehavioralAnalytics

	RenderQuery search_application_render_query.NewRenderQuery

	Search search_application_search.NewSearch
}

type SearchableSnapshots struct {
	CacheStats searchable_snapshots_cache_stats.NewCacheStats

	ClearCache searchable_snapshots_clear_cache.NewClearCache

	Mount searchable_snapshots_mount.NewMount

	Stats searchable_snapshots_stats.NewStats
}

type Security struct {
	ActivateUserProfile security_activate_user_profile.NewActivateUserProfile

	Authenticate security_authenticate.NewAuthenticate

	BulkDeleteRole security_bulk_delete_role.NewBulkDeleteRole

	BulkPutRole security_bulk_put_role.NewBulkPutRole

	BulkUpdateApiKeys security_bulk_update_api_keys.NewBulkUpdateApiKeys

	ChangePassword security_change_password.NewChangePassword

	ClearApiKeyCache security_clear_api_key_cache.NewClearApiKeyCache

	ClearCachedPrivileges security_clear_cached_privileges.NewClearCachedPrivileges

	ClearCachedRealms security_clear_cached_realms.NewClearCachedRealms

	ClearCachedRoles security_clear_cached_roles.NewClearCachedRoles

	ClearCachedServiceTokens security_clear_cached_service_tokens.NewClearCachedServiceTokens

	CloneApiKey security_clone_api_key.NewCloneApiKey

	CreateApiKey security_create_api_key.NewCreateApiKey

	CreateCrossClusterApiKey security_create_cross_cluster_api_key.NewCreateCrossClusterApiKey

	CreateServiceToken security_create_service_token.NewCreateServiceToken

	DelegatePki security_delegate_pki.NewDelegatePki

	DeletePrivileges security_delete_privileges.NewDeletePrivileges

	DeleteRole security_delete_role.NewDeleteRole

	DeleteRoleMapping security_delete_role_mapping.NewDeleteRoleMapping

	DeleteServiceToken security_delete_service_token.NewDeleteServiceToken

	DeleteUser security_delete_user.NewDeleteUser

	DisableUser security_disable_user.NewDisableUser

	DisableUserProfile security_disable_user_profile.NewDisableUserProfile

	EnableUser security_enable_user.NewEnableUser

	EnableUserProfile security_enable_user_profile.NewEnableUserProfile

	EnrollKibana security_enroll_kibana.NewEnrollKibana

	EnrollNode security_enroll_node.NewEnrollNode

	GetApiKey security_get_api_key.NewGetApiKey

	GetBuiltinPrivileges security_get_builtin_privileges.NewGetBuiltinPrivileges

	GetPrivileges security_get_privileges.NewGetPrivileges

	GetRole security_get_role.NewGetRole

	GetRoleMapping security_get_role_mapping.NewGetRoleMapping

	GetServiceAccounts security_get_service_accounts.NewGetServiceAccounts

	GetServiceCredentials security_get_service_credentials.NewGetServiceCredentials

	GetSettings security_get_settings.NewGetSettings

	GetStats security_get_stats.NewGetStats

	GetToken security_get_token.NewGetToken

	GetUser security_get_user.NewGetUser

	GetUserPrivileges security_get_user_privileges.NewGetUserPrivileges

	GetUserProfile security_get_user_profile.NewGetUserProfile

	GrantApiKey security_grant_api_key.NewGrantApiKey

	HasPrivileges security_has_privileges.NewHasPrivileges

	HasPrivilegesUserProfile security_has_privileges_user_profile.NewHasPrivilegesUserProfile

	InvalidateApiKey security_invalidate_api_key.NewInvalidateApiKey

	InvalidateToken security_invalidate_token.NewInvalidateToken

	OidcAuthenticate security_oidc_authenticate.NewOidcAuthenticate

	OidcLogout security_oidc_logout.NewOidcLogout

	OidcPrepareAuthentication security_oidc_prepare_authentication.NewOidcPrepareAuthentication

	PutPrivileges security_put_privileges.NewPutPrivileges

	PutRole security_put_role.NewPutRole

	PutRoleMapping security_put_role_mapping.NewPutRoleMapping

	PutUser security_put_user.NewPutUser

	QueryApiKeys security_query_api_keys.NewQueryApiKeys

	QueryRole security_query_role.NewQueryRole

	QueryUser security_query_user.NewQueryUser

	SamlAuthenticate security_saml_authenticate.NewSamlAuthenticate

	SamlCompleteLogout security_saml_complete_logout.NewSamlCompleteLogout

	SamlInvalidate security_saml_invalidate.NewSamlInvalidate

	SamlLogout security_saml_logout.NewSamlLogout

	SamlPrepareAuthentication security_saml_prepare_authentication.NewSamlPrepareAuthentication

	SamlServiceProviderMetadata security_saml_service_provider_metadata.NewSamlServiceProviderMetadata

	SuggestUserProfiles security_suggest_user_profiles.NewSuggestUserProfiles

	UpdateApiKey security_update_api_key.NewUpdateApiKey

	UpdateCrossClusterApiKey security_update_cross_cluster_api_key.NewUpdateCrossClusterApiKey

	UpdateSettings security_update_settings.NewUpdateSettings

	UpdateUserProfileData security_update_user_profile_data.NewUpdateUserProfileData
}

type Shutdown struct {
	DeleteNode shutdown_delete_node.NewDeleteNode

	GetNode shutdown_get_node.NewGetNode

	PutNode shutdown_put_node.NewPutNode
}

type Simulate struct {
	Ingest simulate_ingest.NewIngest
}

type Slm struct {
	DeleteLifecycle slm_delete_lifecycle.NewDeleteLifecycle

	ExecuteLifecycle slm_execute_lifecycle.NewExecuteLifecycle

	ExecuteRetention slm_execute_retention.NewExecuteRetention

	GetLifecycle slm_get_lifecycle.NewGetLifecycle

	GetStats slm_get_stats.NewGetStats

	GetStatus slm_get_status.NewGetStatus

	PutLifecycle slm_put_lifecycle.NewPutLifecycle

	Start slm_start.NewStart

	Stop slm_stop.NewStop
}

type Snapshot struct {
	CleanupRepository snapshot_cleanup_repository.NewCleanupRepository

	Clone snapshot_clone.NewClone

	Create snapshot_create.NewCreate

	CreateRepository snapshot_create_repository.NewCreateRepository

	Delete snapshot_delete.NewDelete

	DeleteRepository snapshot_delete_repository.NewDeleteRepository

	Get snapshot_get.NewGet

	GetRepository snapshot_get_repository.NewGetRepository

	RepositoryAnalyze snapshot_repository_analyze.NewRepositoryAnalyze

	RepositoryVerifyIntegrity snapshot_repository_verify_integrity.NewRepositoryVerifyIntegrity

	Restore snapshot_restore.NewRestore

	Status snapshot_status.NewStatus

	VerifyRepository snapshot_verify_repository.NewVerifyRepository
}

type Sql struct {
	ClearCursor sql_clear_cursor.NewClearCursor

	DeleteAsync sql_delete_async.NewDeleteAsync

	GetAsync sql_get_async.NewGetAsync

	GetAsyncStatus sql_get_async_status.NewGetAsyncStatus

	Query sql_query.NewQuery

	Translate sql_translate.NewTranslate
}

type Ssl struct {
	Certificates ssl_certificates.NewCertificates
}

type Streams struct {
	LogsDisable streams_logs_disable.NewLogsDisable

	LogsEnable streams_logs_enable.NewLogsEnable

	Status streams_status.NewStatus
}

type Synonyms struct {
	DeleteSynonym synonyms_delete_synonym.NewDeleteSynonym

	DeleteSynonymRule synonyms_delete_synonym_rule.NewDeleteSynonymRule

	GetSynonym synonyms_get_synonym.NewGetSynonym

	GetSynonymRule synonyms_get_synonym_rule.NewGetSynonymRule

	GetSynonymsSets synonyms_get_synonyms_sets.NewGetSynonymsSets

	PutSynonym synonyms_put_synonym.NewPutSynonym

	PutSynonymRule synonyms_put_synonym_rule.NewPutSynonymRule
}

type Tasks struct {
	Cancel tasks_cancel.NewCancel

	Get tasks_get.NewGet

	List tasks_list.NewList
}

type TextStructure struct {
	FindFieldStructure text_structure_find_field_structure.NewFindFieldStructure

	FindMessageStructure text_structure_find_message_structure.NewFindMessageStructure

	FindStructure text_structure_find_structure.NewFindStructure

	TestGrokPattern text_structure_test_grok_pattern.NewTestGrokPattern
}

type Transform struct {
	DeleteTransform transform_delete_transform.NewDeleteTransform

	GetNodeStats transform_get_node_stats.NewGetNodeStats

	GetTransform transform_get_transform.NewGetTransform

	GetTransformStats transform_get_transform_stats.NewGetTransformStats

	PreviewTransform transform_preview_transform.NewPreviewTransform

	PutTransform transform_put_transform.NewPutTransform

	ResetTransform transform_reset_transform.NewResetTransform

	ScheduleNowTransform transform_schedule_now_transform.NewScheduleNowTransform

	SetUpgradeMode transform_set_upgrade_mode.NewSetUpgradeMode

	StartTransform transform_start_transform.NewStartTransform

	StopTransform transform_stop_transform.NewStopTransform

	UpdateTransform transform_update_transform.NewUpdateTransform

	UpgradeTransforms transform_upgrade_transforms.NewUpgradeTransforms
}

type Watcher struct {
	AckWatch watcher_ack_watch.NewAckWatch

	ActivateWatch watcher_activate_watch.NewActivateWatch

	DeactivateWatch watcher_deactivate_watch.NewDeactivateWatch

	DeleteWatch watcher_delete_watch.NewDeleteWatch

	ExecuteWatch watcher_execute_watch.NewExecuteWatch

	GetSettings watcher_get_settings.NewGetSettings

	GetWatch watcher_get_watch.NewGetWatch

	PutWatch watcher_put_watch.NewPutWatch

	QueryWatches watcher_query_watches.NewQueryWatches

	Start watcher_start.NewStart

	Stats watcher_stats.NewStats

	Stop watcher_stop.NewStop

	UpdateSettings watcher_update_settings.NewUpdateSettings
}

type Xpack struct {
	Info xpack_info.NewInfo

	Usage xpack_usage.NewUsage
}

type API struct {
	AsyncSearch         AsyncSearch
	Autoscaling         Autoscaling
	Cat                 Cat
	Ccr                 Ccr
	Cluster             Cluster
	Connector           Connector
	Core                Core
	DanglingIndices     DanglingIndices
	Enrich              Enrich
	Eql                 Eql
	Esql                Esql
	Features            Features
	Fleet               Fleet
	Graph               Graph
	Ilm                 Ilm
	Indices             Indices
	Inference           Inference
	Ingest              Ingest
	License             License
	Logstash            Logstash
	Migration           Migration
	Ml                  Ml
	Monitoring          Monitoring
	Nodes               Nodes
	Profiling           Profiling
	Project             Project
	QueryRules          QueryRules
	Rollup              Rollup
	SearchApplication   SearchApplication
	SearchableSnapshots SearchableSnapshots
	Security            Security
	Shutdown            Shutdown
	Simulate            Simulate
	Slm                 Slm
	Snapshot            Snapshot
	Sql                 Sql
	Ssl                 Ssl
	Streams             Streams
	Synonyms            Synonyms
	Tasks               Tasks
	TextStructure       TextStructure
	Transform           Transform
	Watcher             Watcher
	Xpack               Xpack

	Bulk core_bulk.NewBulk

	CancelReindex core_cancel_reindex.NewCancelReindex

	Capabilities core_capabilities.NewCapabilities

	ClearScroll core_clear_scroll.NewClearScroll

	ClosePointInTime core_close_point_in_time.NewClosePointInTime

	Count core_count.NewCount

	Create core_create.NewCreate

	Delete core_delete.NewDelete

	DeleteByQuery core_delete_by_query.NewDeleteByQuery

	DeleteByQueryRethrottle core_delete_by_query_rethrottle.NewDeleteByQueryRethrottle

	DeleteScript core_delete_script.NewDeleteScript

	Exists core_exists.NewExists

	ExistsSource core_exists_source.NewExistsSource

	Explain core_explain.NewExplain

	FieldCaps core_field_caps.NewFieldCaps

	Get core_get.NewGet

	GetReindex core_get_reindex.NewGetReindex

	GetScript core_get_script.NewGetScript

	GetScriptContext core_get_script_context.NewGetScriptContext

	GetScriptLanguages core_get_script_languages.NewGetScriptLanguages

	GetSource core_get_source.NewGetSource

	HealthReport core_health_report.NewHealthReport

	Index core_index.NewIndex

	Info core_info.NewInfo

	KnnSearch core_knn_search.NewKnnSearch

	ListReindex core_list_reindex.NewListReindex

	Mget core_mget.NewMget

	Msearch core_msearch.NewMsearch

	MsearchTemplate core_msearch_template.NewMsearchTemplate

	Mtermvectors core_mtermvectors.NewMtermvectors

	OpenPointInTime core_open_point_in_time.NewOpenPointInTime

	Ping core_ping.NewPing

	PutScript core_put_script.NewPutScript

	RankEval core_rank_eval.NewRankEval

	Reindex core_reindex.NewReindex

	ReindexRethrottle core_reindex_rethrottle.NewReindexRethrottle

	RenderSearchTemplate core_render_search_template.NewRenderSearchTemplate

	ScriptsPainlessExecute core_scripts_painless_execute.NewScriptsPainlessExecute

	Scroll core_scroll.NewScroll

	Search core_search.NewSearch

	SearchMvt core_search_mvt.NewSearchMvt

	SearchShards core_search_shards.NewSearchShards

	SearchTemplate core_search_template.NewSearchTemplate

	TermsEnum core_terms_enum.NewTermsEnum

	Termvectors core_termvectors.NewTermvectors

	Update core_update.NewUpdate

	UpdateByQuery core_update_by_query.NewUpdateByQuery

	UpdateByQueryRethrottle core_update_by_query_rethrottle.NewUpdateByQueryRethrottle
}

func New(tp elastictransport.Interface) *API { _ = "STUB: not implemented"; return nil }

type MethodAsyncSearch struct {
	tp elastictransport.Interface
}

type MethodAutoscaling struct {
	tp elastictransport.Interface
}

type MethodCat struct {
	tp elastictransport.Interface
}

type MethodCcr struct {
	tp elastictransport.Interface
}

type MethodCluster struct {
	tp elastictransport.Interface
}

type MethodConnector struct {
	tp elastictransport.Interface
}

type MethodCore struct {
	tp elastictransport.Interface
}

type MethodDanglingIndices struct {
	tp elastictransport.Interface
}

type MethodEnrich struct {
	tp elastictransport.Interface
}

type MethodEql struct {
	tp elastictransport.Interface
}

type MethodEsql struct {
	tp elastictransport.Interface
}

type MethodFeatures struct {
	tp elastictransport.Interface
}

type MethodFleet struct {
	tp elastictransport.Interface
}

type MethodGraph struct {
	tp elastictransport.Interface
}

type MethodIlm struct {
	tp elastictransport.Interface
}

type MethodIndices struct {
	tp elastictransport.Interface
}

type MethodInference struct {
	tp elastictransport.Interface
}

type MethodIngest struct {
	tp elastictransport.Interface
}

type MethodLicense struct {
	tp elastictransport.Interface
}

type MethodLogstash struct {
	tp elastictransport.Interface
}

type MethodMigration struct {
	tp elastictransport.Interface
}

type MethodMl struct {
	tp elastictransport.Interface
}

type MethodMonitoring struct {
	tp elastictransport.Interface
}

type MethodNodes struct {
	tp elastictransport.Interface
}

type MethodProfiling struct {
	tp elastictransport.Interface
}

type MethodProject struct {
	tp elastictransport.Interface
}

type MethodQueryRules struct {
	tp elastictransport.Interface
}

type MethodRollup struct {
	tp elastictransport.Interface
}

type MethodSearchApplication struct {
	tp elastictransport.Interface
}

type MethodSearchableSnapshots struct {
	tp elastictransport.Interface
}

type MethodSecurity struct {
	tp elastictransport.Interface
}

type MethodShutdown struct {
	tp elastictransport.Interface
}

type MethodSimulate struct {
	tp elastictransport.Interface
}

type MethodSlm struct {
	tp elastictransport.Interface
}

type MethodSnapshot struct {
	tp elastictransport.Interface
}

type MethodSql struct {
	tp elastictransport.Interface
}

type MethodSsl struct {
	tp elastictransport.Interface
}

type MethodStreams struct {
	tp elastictransport.Interface
}

type MethodSynonyms struct {
	tp elastictransport.Interface
}

type MethodTasks struct {
	tp elastictransport.Interface
}

type MethodTextStructure struct {
	tp elastictransport.Interface
}

type MethodTransform struct {
	tp elastictransport.Interface
}

type MethodWatcher struct {
	tp elastictransport.Interface
}

type MethodXpack struct {
	tp elastictransport.Interface
}

type MethodAPI struct {
	tp                  elastictransport.Interface
	AsyncSearch         MethodAsyncSearch
	Autoscaling         MethodAutoscaling
	Cat                 MethodCat
	Ccr                 MethodCcr
	Cluster             MethodCluster
	Connector           MethodConnector
	Core                MethodCore
	DanglingIndices     MethodDanglingIndices
	Enrich              MethodEnrich
	Eql                 MethodEql
	Esql                MethodEsql
	Features            MethodFeatures
	Fleet               MethodFleet
	Graph               MethodGraph
	Ilm                 MethodIlm
	Indices             MethodIndices
	Inference           MethodInference
	Ingest              MethodIngest
	License             MethodLicense
	Logstash            MethodLogstash
	Migration           MethodMigration
	Ml                  MethodMl
	Monitoring          MethodMonitoring
	Nodes               MethodNodes
	Profiling           MethodProfiling
	Project             MethodProject
	QueryRules          MethodQueryRules
	Rollup              MethodRollup
	SearchApplication   MethodSearchApplication
	SearchableSnapshots MethodSearchableSnapshots
	Security            MethodSecurity
	Shutdown            MethodShutdown
	Simulate            MethodSimulate
	Slm                 MethodSlm
	Snapshot            MethodSnapshot
	Sql                 MethodSql
	Ssl                 MethodSsl
	Streams             MethodStreams
	Synonyms            MethodSynonyms
	Tasks               MethodTasks
	TextStructure       MethodTextStructure
	Transform           MethodTransform
	Watcher             MethodWatcher
	Xpack               MethodXpack
}

func (p *MethodAPI) Bulk() *core_bulk.Bulk { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) CancelReindex(taskid string) *core_cancel_reindex.CancelReindex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Capabilities() *core_capabilities.Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) ClearScroll() *core_clear_scroll.ClearScroll {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) ClosePointInTime() *core_close_point_in_time.ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Count() *core_count.Count { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) Create(index, id string) *core_create.Create {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Delete(index, id string) *core_delete.Delete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) DeleteByQuery(index string) *core_delete_by_query.DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) DeleteByQueryRethrottle(taskid string) *core_delete_by_query_rethrottle.DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) DeleteScript(id string) *core_delete_script.DeleteScript {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Exists(index, id string) *core_exists.Exists {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) ExistsSource(index, id string) *core_exists_source.ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Explain(index, id string) *core_explain.Explain {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) FieldCaps() *core_field_caps.FieldCaps { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) Get(index, id string) *core_get.Get { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) GetReindex(taskid string) *core_get_reindex.GetReindex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) GetScript(id string) *core_get_script.GetScript {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) GetScriptContext() *core_get_script_context.GetScriptContext {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) GetScriptLanguages() *core_get_script_languages.GetScriptLanguages {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) GetSource(index, id string) *core_get_source.GetSource {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) HealthReport() *core_health_report.HealthReport {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Index(index string) *core_index.Index { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) Info() *core_info.Info { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) KnnSearch(index string) *core_knn_search.KnnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) ListReindex() *core_list_reindex.ListReindex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Mget() *core_mget.Mget { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) Msearch() *core_msearch.Msearch { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) MsearchTemplate() *core_msearch_template.MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Mtermvectors() *core_mtermvectors.Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) OpenPointInTime(index string) *core_open_point_in_time.OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Ping() *core_ping.Ping { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) PutScript(id string) *core_put_script.PutScript {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) RankEval() *core_rank_eval.RankEval { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) Reindex() *core_reindex.Reindex { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) ReindexRethrottle(taskid string) *core_reindex_rethrottle.ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) RenderSearchTemplate() *core_render_search_template.RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) ScriptsPainlessExecute() *core_scripts_painless_execute.ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Scroll() *core_scroll.Scroll { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) Search() *core_search.Search { _ = "STUB: not implemented"; return nil }

func (p *MethodAPI) SearchMvt(index, field, zoom, x, y string) *core_search_mvt.SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) SearchShards() *core_search_shards.SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) SearchTemplate() *core_search_template.SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) TermsEnum(index string) *core_terms_enum.TermsEnum {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Termvectors(index string) *core_termvectors.Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) Update(index, id string) *core_update.Update {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) UpdateByQuery(index string) *core_update_by_query.UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAPI) UpdateByQueryRethrottle(taskid string) *core_update_by_query_rethrottle.UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAsyncSearch) Delete(id string) *async_search_delete.Delete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAsyncSearch) Get(id string) *async_search_get.Get {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAsyncSearch) Status(id string) *async_search_status.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAsyncSearch) Submit() *async_search_submit.Submit {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAutoscaling) DeleteAutoscalingPolicy(name string) *autoscaling_delete_autoscaling_policy.DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAutoscaling) GetAutoscalingCapacity() *autoscaling_get_autoscaling_capacity.GetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAutoscaling) GetAutoscalingPolicy(name string) *autoscaling_get_autoscaling_policy.GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodAutoscaling) PutAutoscalingPolicy(name string) *autoscaling_put_autoscaling_policy.PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCat) Aliases() *cat_aliases.Aliases { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Allocation() *cat_allocation.Allocation { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) CircuitBreaker() *cat_circuit_breaker.CircuitBreaker {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCat) ComponentTemplates() *cat_component_templates.ComponentTemplates {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCat) Count() *cat_count.Count { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Fielddata() *cat_fielddata.Fielddata { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Health() *cat_health.Health { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Help() *cat_help.Help { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Indices() *cat_indices.Indices { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Master() *cat_master.Master { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) MlDataFrameAnalytics() *cat_ml_data_frame_analytics.MlDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCat) MlDatafeeds() *cat_ml_datafeeds.MlDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCat) MlJobs() *cat_ml_jobs.MlJobs { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) MlTrainedModels() *cat_ml_trained_models.MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCat) Nodeattrs() *cat_nodeattrs.Nodeattrs { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Nodes() *cat_nodes.Nodes { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) PendingTasks() *cat_pending_tasks.PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCat) Plugins() *cat_plugins.Plugins { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Recovery() *cat_recovery.Recovery { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Repositories() *cat_repositories.Repositories {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCat) Segments() *cat_segments.Segments { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Shards() *cat_shards.Shards { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Snapshots() *cat_snapshots.Snapshots { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Tasks() *cat_tasks.Tasks { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Templates() *cat_templates.Templates { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) ThreadPool() *cat_thread_pool.ThreadPool { _ = "STUB: not implemented"; return nil }

func (p *MethodCat) Transforms() *cat_transforms.Transforms { _ = "STUB: not implemented"; return nil }

func (p *MethodCcr) DeleteAutoFollowPattern(name string) *ccr_delete_auto_follow_pattern.DeleteAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) Follow(index string) *ccr_follow.Follow { _ = "STUB: not implemented"; return nil }

func (p *MethodCcr) FollowInfo(index string) *ccr_follow_info.FollowInfo {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) FollowStats(index string) *ccr_follow_stats.FollowStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) ForgetFollower(index string) *ccr_forget_follower.ForgetFollower {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) GetAutoFollowPattern() *ccr_get_auto_follow_pattern.GetAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) PauseAutoFollowPattern(name string) *ccr_pause_auto_follow_pattern.PauseAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) PauseFollow(index string) *ccr_pause_follow.PauseFollow {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) PutAutoFollowPattern(name string) *ccr_put_auto_follow_pattern.PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) ResumeAutoFollowPattern(name string) *ccr_resume_auto_follow_pattern.ResumeAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) ResumeFollow(index string) *ccr_resume_follow.ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCcr) Stats() *ccr_stats.Stats { _ = "STUB: not implemented"; return nil }

func (p *MethodCcr) Unfollow(index string) *ccr_unfollow.Unfollow {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) AllocationExplain() *cluster_allocation_explain.AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) DeleteComponentTemplate(name string) *cluster_delete_component_template.DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) DeleteVotingConfigExclusions() *cluster_delete_voting_config_exclusions.DeleteVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) ExistsComponentTemplate(name string) *cluster_exists_component_template.ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) GetComponentTemplate() *cluster_get_component_template.GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) GetSettings() *cluster_get_settings.GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) Health() *cluster_health.Health { _ = "STUB: not implemented"; return nil }

func (p *MethodCluster) Info(target string) *cluster_info.Info {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) PendingTasks() *cluster_pending_tasks.PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) PostVotingConfigExclusions() *cluster_post_voting_config_exclusions.PostVotingConfigExclusions {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) PutComponentTemplate(name string) *cluster_put_component_template.PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) PutSettings() *cluster_put_settings.PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) RemoteInfo() *cluster_remote_info.RemoteInfo {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCluster) Reroute() *cluster_reroute.Reroute { _ = "STUB: not implemented"; return nil }

func (p *MethodCluster) State() *cluster_state.State { _ = "STUB: not implemented"; return nil }

func (p *MethodCluster) Stats() *cluster_stats.Stats { _ = "STUB: not implemented"; return nil }

func (p *MethodConnector) CheckIn(connectorid string) *connector_check_in.CheckIn {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) Delete(connectorid string) *connector_delete.Delete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) Get(connectorid string) *connector_get.Get {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) LastSync(connectorid string) *connector_last_sync.LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) List() *connector_list.List { _ = "STUB: not implemented"; return nil }

func (p *MethodConnector) Post() *connector_post.Post { _ = "STUB: not implemented"; return nil }

func (p *MethodConnector) Put() *connector_put.Put { _ = "STUB: not implemented"; return nil }

func (p *MethodConnector) SecretDelete(id string) *connector_secret_delete.SecretDelete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SecretGet(id string) *connector_secret_get.SecretGet {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SecretPost() *connector_secret_post.SecretPost {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SecretPut(id string) *connector_secret_put.SecretPut {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobCancel(connectorsyncjobid string) *connector_sync_job_cancel.SyncJobCancel {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobCheckIn(connectorsyncjobid string) *connector_sync_job_check_in.SyncJobCheckIn {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobClaim(connectorsyncjobid string) *connector_sync_job_claim.SyncJobClaim {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobDelete(connectorsyncjobid string) *connector_sync_job_delete.SyncJobDelete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobError(connectorsyncjobid string) *connector_sync_job_error.SyncJobError {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobGet(connectorsyncjobid string) *connector_sync_job_get.SyncJobGet {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobList() *connector_sync_job_list.SyncJobList {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobPost() *connector_sync_job_post.SyncJobPost {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) SyncJobUpdateStats(connectorsyncjobid string) *connector_sync_job_update_stats.SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateActiveFiltering(connectorid string) *connector_update_active_filtering.UpdateActiveFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateApiKeyId(connectorid string) *connector_update_api_key_id.UpdateApiKeyId {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateConfiguration(connectorid string) *connector_update_configuration.UpdateConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateError(connectorid string) *connector_update_error.UpdateError {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateFeatures(connectorid string) *connector_update_features.UpdateFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateFiltering(connectorid string) *connector_update_filtering.UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateFilteringValidation(connectorid string) *connector_update_filtering_validation.UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateIndexName(connectorid string) *connector_update_index_name.UpdateIndexName {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateName(connectorid string) *connector_update_name.UpdateName {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateNative(connectorid string) *connector_update_native.UpdateNative {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdatePipeline(connectorid string) *connector_update_pipeline.UpdatePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateScheduling(connectorid string) *connector_update_scheduling.UpdateScheduling {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateServiceType(connectorid string) *connector_update_service_type.UpdateServiceType {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodConnector) UpdateStatus(connectorid string) *connector_update_status.UpdateStatus {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Bulk() *core_bulk.Bulk { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) CancelReindex(taskid string) *core_cancel_reindex.CancelReindex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Capabilities() *core_capabilities.Capabilities {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) ClearScroll() *core_clear_scroll.ClearScroll {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) ClosePointInTime() *core_close_point_in_time.ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Count() *core_count.Count { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) Create(index, id string) *core_create.Create {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Delete(index, id string) *core_delete.Delete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) DeleteByQuery(index string) *core_delete_by_query.DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) DeleteByQueryRethrottle(taskid string) *core_delete_by_query_rethrottle.DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) DeleteScript(id string) *core_delete_script.DeleteScript {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Exists(index, id string) *core_exists.Exists {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) ExistsSource(index, id string) *core_exists_source.ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Explain(index, id string) *core_explain.Explain {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) FieldCaps() *core_field_caps.FieldCaps { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) Get(index, id string) *core_get.Get { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) GetReindex(taskid string) *core_get_reindex.GetReindex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) GetScript(id string) *core_get_script.GetScript {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) GetScriptContext() *core_get_script_context.GetScriptContext {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) GetScriptLanguages() *core_get_script_languages.GetScriptLanguages {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) GetSource(index, id string) *core_get_source.GetSource {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) HealthReport() *core_health_report.HealthReport {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Index(index string) *core_index.Index { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) Info() *core_info.Info { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) KnnSearch(index string) *core_knn_search.KnnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) ListReindex() *core_list_reindex.ListReindex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Mget() *core_mget.Mget { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) Msearch() *core_msearch.Msearch { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) MsearchTemplate() *core_msearch_template.MsearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Mtermvectors() *core_mtermvectors.Mtermvectors {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) OpenPointInTime(index string) *core_open_point_in_time.OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Ping() *core_ping.Ping { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) PutScript(id string) *core_put_script.PutScript {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) RankEval() *core_rank_eval.RankEval { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) Reindex() *core_reindex.Reindex { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) ReindexRethrottle(taskid string) *core_reindex_rethrottle.ReindexRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) RenderSearchTemplate() *core_render_search_template.RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) ScriptsPainlessExecute() *core_scripts_painless_execute.ScriptsPainlessExecute {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Scroll() *core_scroll.Scroll { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) Search() *core_search.Search { _ = "STUB: not implemented"; return nil }

func (p *MethodCore) SearchMvt(index, field, zoom, x, y string) *core_search_mvt.SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) SearchShards() *core_search_shards.SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) SearchTemplate() *core_search_template.SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) TermsEnum(index string) *core_terms_enum.TermsEnum {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Termvectors(index string) *core_termvectors.Termvectors {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) Update(index, id string) *core_update.Update {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) UpdateByQuery(index string) *core_update_by_query.UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodCore) UpdateByQueryRethrottle(taskid string) *core_update_by_query_rethrottle.UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodDanglingIndices) DeleteDanglingIndex(indexuuid string) *dangling_indices_delete_dangling_index.DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodDanglingIndices) ImportDanglingIndex(indexuuid string) *dangling_indices_import_dangling_index.ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodDanglingIndices) ListDanglingIndices() *dangling_indices_list_dangling_indices.ListDanglingIndices {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEnrich) DeletePolicy(name string) *enrich_delete_policy.DeletePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEnrich) ExecutePolicy(name string) *enrich_execute_policy.ExecutePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEnrich) GetPolicy() *enrich_get_policy.GetPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEnrich) PutPolicy(name string) *enrich_put_policy.PutPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEnrich) Stats() *enrich_stats.Stats { _ = "STUB: not implemented"; return nil }

func (p *MethodEql) Delete(id string) *eql_delete.Delete { _ = "STUB: not implemented"; return nil }

func (p *MethodEql) Get(id string) *eql_get.Get { _ = "STUB: not implemented"; return nil }

func (p *MethodEql) GetStatus(id string) *eql_get_status.GetStatus {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEql) Search(index string) *eql_search.Search { _ = "STUB: not implemented"; return nil }

func (p *MethodEsql) AsyncQuery() *esql_async_query.AsyncQuery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEsql) AsyncQueryDelete(id string) *esql_async_query_delete.AsyncQueryDelete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEsql) AsyncQueryGet(id string) *esql_async_query_get.AsyncQueryGet {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEsql) AsyncQueryStop(id string) *esql_async_query_stop.AsyncQueryStop {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEsql) DeleteView(name string) *esql_delete_view.DeleteView {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEsql) GetQuery(id string) *esql_get_query.GetQuery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEsql) GetView() *esql_get_view.GetView { _ = "STUB: not implemented"; return nil }

func (p *MethodEsql) ListQueries() *esql_list_queries.ListQueries {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEsql) PutView(name string) *esql_put_view.PutView {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodEsql) Query() *esql_query.Query { _ = "STUB: not implemented"; return nil }

func (p *MethodFeatures) GetFeatures() *features_get_features.GetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodFeatures) ResetFeatures() *features_reset_features.ResetFeatures {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodFleet) DeleteSecret(id string) *fleet_delete_secret.DeleteSecret {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodFleet) GetSecret(id string) *fleet_get_secret.GetSecret {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodFleet) GlobalCheckpoints(index string) *fleet_global_checkpoints.GlobalCheckpoints {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodFleet) Msearch() *fleet_msearch.Msearch { _ = "STUB: not implemented"; return nil }

func (p *MethodFleet) PostSecret() *fleet_post_secret.PostSecret {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodFleet) Search(index string) *fleet_search.Search {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodGraph) Explore(index string) *graph_explore.Explore {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIlm) DeleteLifecycle(policy string) *ilm_delete_lifecycle.DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIlm) ExplainLifecycle(index string) *ilm_explain_lifecycle.ExplainLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIlm) GetLifecycle() *ilm_get_lifecycle.GetLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIlm) GetStatus() *ilm_get_status.GetStatus { _ = "STUB: not implemented"; return nil }

func (p *MethodIlm) MigrateToDataTiers() *ilm_migrate_to_data_tiers.MigrateToDataTiers {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIlm) MoveToStep(index string) *ilm_move_to_step.MoveToStep {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIlm) PutLifecycle(policy string) *ilm_put_lifecycle.PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIlm) RemovePolicy(index string) *ilm_remove_policy.RemovePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIlm) Retry(index string) *ilm_retry.Retry { _ = "STUB: not implemented"; return nil }

func (p *MethodIlm) Start() *ilm_start.Start { _ = "STUB: not implemented"; return nil }

func (p *MethodIlm) Stop() *ilm_stop.Stop { _ = "STUB: not implemented"; return nil }

func (p *MethodIndices) AddBlock(index, block string) *indices_add_block.AddBlock {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Analyze() *indices_analyze.Analyze { _ = "STUB: not implemented"; return nil }

func (p *MethodIndices) CancelMigrateReindex(index string) *indices_cancel_migrate_reindex.CancelMigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ClearCache() *indices_clear_cache.ClearCache {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Clone(index, target string) *indices_clone.Clone {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Close(index string) *indices_close.Close {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Create(index string) *indices_create.Create {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) CreateDataStream(name string) *indices_create_data_stream.CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) CreateFrom(source, dest string) *indices_create_from.CreateFrom {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) DataStreamsStats() *indices_data_streams_stats.DataStreamsStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Delete(index string) *indices_delete.Delete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) DeleteAlias(index, name string) *indices_delete_alias.DeleteAlias {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) DeleteDataLifecycle(name string) *indices_delete_data_lifecycle.DeleteDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) DeleteDataStream(name string) *indices_delete_data_stream.DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) DeleteDataStreamOptions(name string) *indices_delete_data_stream_options.DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) DeleteIndexTemplate(name string) *indices_delete_index_template.DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) DeleteTemplate(name string) *indices_delete_template.DeleteTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) DiskUsage(index string) *indices_disk_usage.DiskUsage {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Downsample(index, targetindex string) *indices_downsample.Downsample {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Exists(index string) *indices_exists.Exists {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ExistsAlias(name string) *indices_exists_alias.ExistsAlias {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ExistsIndexTemplate(name string) *indices_exists_index_template.ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ExistsTemplate(name string) *indices_exists_template.ExistsTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ExplainDataLifecycle(index string) *indices_explain_data_lifecycle.ExplainDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) FieldUsageStats(index string) *indices_field_usage_stats.FieldUsageStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Flush() *indices_flush.Flush { _ = "STUB: not implemented"; return nil }

func (p *MethodIndices) Forcemerge() *indices_forcemerge.Forcemerge {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Get(index string) *indices_get.Get { _ = "STUB: not implemented"; return nil }

func (p *MethodIndices) GetAlias() *indices_get_alias.GetAlias {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetDataLifecycle(name string) *indices_get_data_lifecycle.GetDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetDataLifecycleStats() *indices_get_data_lifecycle_stats.GetDataLifecycleStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetDataStream() *indices_get_data_stream.GetDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetDataStreamMappings(name string) *indices_get_data_stream_mappings.GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetDataStreamOptions(name string) *indices_get_data_stream_options.GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetDataStreamSettings(name string) *indices_get_data_stream_settings.GetDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetFieldMapping(fields string) *indices_get_field_mapping.GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetIndexTemplate() *indices_get_index_template.GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetMapping() *indices_get_mapping.GetMapping {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetMigrateReindexStatus(index string) *indices_get_migrate_reindex_status.GetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetSettings() *indices_get_settings.GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) GetTemplate() *indices_get_template.GetTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) MigrateReindex() *indices_migrate_reindex.MigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) MigrateToDataStream(name string) *indices_migrate_to_data_stream.MigrateToDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ModifyDataStream() *indices_modify_data_stream.ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Open(index string) *indices_open.Open {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PromoteDataStream(name string) *indices_promote_data_stream.PromoteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutAlias(index, name string) *indices_put_alias.PutAlias {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutDataLifecycle(name string) *indices_put_data_lifecycle.PutDataLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutDataStreamMappings(name string) *indices_put_data_stream_mappings.PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutDataStreamOptions(name string) *indices_put_data_stream_options.PutDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutDataStreamSettings(name string) *indices_put_data_stream_settings.PutDataStreamSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutIndexTemplate(name string) *indices_put_index_template.PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutMapping(index string) *indices_put_mapping.PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutSettings() *indices_put_settings.PutSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) PutTemplate(name string) *indices_put_template.PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Recovery() *indices_recovery.Recovery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Refresh() *indices_refresh.Refresh { _ = "STUB: not implemented"; return nil }

func (p *MethodIndices) ReloadSearchAnalyzers(index string) *indices_reload_search_analyzers.ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) RemoveBlock(index, block string) *indices_remove_block.RemoveBlock {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ResolveCluster() *indices_resolve_cluster.ResolveCluster {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ResolveIndex(name string) *indices_resolve_index.ResolveIndex {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Rollover(alias string) *indices_rollover.Rollover {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Segments() *indices_segments.Segments {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ShardStores() *indices_shard_stores.ShardStores {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Shrink(index, target string) *indices_shrink.Shrink {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) SimulateIndexTemplate(name string) *indices_simulate_index_template.SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) SimulateTemplate() *indices_simulate_template.SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Split(index, target string) *indices_split.Split {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) Stats() *indices_stats.Stats { _ = "STUB: not implemented"; return nil }

func (p *MethodIndices) UpdateAliases() *indices_update_aliases.UpdateAliases {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIndices) ValidateQuery() *indices_validate_query.ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) ChatCompletionUnified(inferenceid string) *inference_chat_completion_unified.ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) Completion(inferenceid string) *inference_completion.Completion {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) Delete(inferenceid string) *inference_delete.Delete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) Embedding(inferenceid string) *inference_embedding.Embedding {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) Get() *inference_get.Get { _ = "STUB: not implemented"; return nil }

func (p *MethodInference) Inference(inferenceid string) *inference_inference.Inference {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) Put(inferenceid string) *inference_put.Put {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutAi21(tasktype, ai21inferenceid string) *inference_put_ai21.PutAi21 {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutAlibabacloud(tasktype, alibabacloudinferenceid string) *inference_put_alibabacloud.PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutAmazonbedrock(tasktype, amazonbedrockinferenceid string) *inference_put_amazonbedrock.PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutAmazonsagemaker(tasktype, amazonsagemakerinferenceid string) *inference_put_amazonsagemaker.PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutAnthropic(tasktype, anthropicinferenceid string) *inference_put_anthropic.PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutAzureaistudio(tasktype, azureaistudioinferenceid string) *inference_put_azureaistudio.PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutAzureopenai(tasktype, azureopenaiinferenceid string) *inference_put_azureopenai.PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutCohere(tasktype, cohereinferenceid string) *inference_put_cohere.PutCohere {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutContextualai(tasktype, contextualaiinferenceid string) *inference_put_contextualai.PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutCustom(tasktype, custominferenceid string) *inference_put_custom.PutCustom {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutDeepseek(tasktype, deepseekinferenceid string) *inference_put_deepseek.PutDeepseek {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutElasticsearch(tasktype, elasticsearchinferenceid string) *inference_put_elasticsearch.PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutElser(tasktype, elserinferenceid string) *inference_put_elser.PutElser {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutFireworksai(tasktype, fireworksaiinferenceid string) *inference_put_fireworksai.PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutGoogleaistudio(tasktype, googleaistudioinferenceid string) *inference_put_googleaistudio.PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutGooglevertexai(tasktype, googlevertexaiinferenceid string) *inference_put_googlevertexai.PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutGroq(tasktype, groqinferenceid string) *inference_put_groq.PutGroq {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutHuggingFace(tasktype, huggingfaceinferenceid string) *inference_put_hugging_face.PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutJinaai(tasktype, jinaaiinferenceid string) *inference_put_jinaai.PutJinaai {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutLlama(tasktype, llamainferenceid string) *inference_put_llama.PutLlama {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutMistral(tasktype, mistralinferenceid string) *inference_put_mistral.PutMistral {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutNvidia(tasktype, nvidiainferenceid string) *inference_put_nvidia.PutNvidia {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutOpenai(tasktype, openaiinferenceid string) *inference_put_openai.PutOpenai {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutOpenshiftAi(tasktype, openshiftaiinferenceid string) *inference_put_openshift_ai.PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutVoyageai(tasktype, voyageaiinferenceid string) *inference_put_voyageai.PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) PutWatsonx(tasktype, watsonxinferenceid string) *inference_put_watsonx.PutWatsonx {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) Rerank(inferenceid string) *inference_rerank.Rerank {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) SparseEmbedding(inferenceid string) *inference_sparse_embedding.SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) StreamCompletion(inferenceid string) *inference_stream_completion.StreamCompletion {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) TextEmbedding(inferenceid string) *inference_text_embedding.TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodInference) Update(inferenceid string) *inference_update.Update {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) DeleteGeoipDatabase(id string) *ingest_delete_geoip_database.DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) DeleteIpLocationDatabase(id string) *ingest_delete_ip_location_database.DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) DeletePipeline(id string) *ingest_delete_pipeline.DeletePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) GeoIpStats() *ingest_geo_ip_stats.GeoIpStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) GetGeoipDatabase() *ingest_get_geoip_database.GetGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) GetIpLocationDatabase() *ingest_get_ip_location_database.GetIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) GetPipeline() *ingest_get_pipeline.GetPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) ProcessorGrok() *ingest_processor_grok.ProcessorGrok {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) PutGeoipDatabase(id string) *ingest_put_geoip_database.PutGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) PutIpLocationDatabase(id string) *ingest_put_ip_location_database.PutIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) PutPipeline(id string) *ingest_put_pipeline.PutPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodIngest) Simulate() *ingest_simulate.Simulate { _ = "STUB: not implemented"; return nil }

func (p *MethodLicense) Delete() *license_delete.Delete { _ = "STUB: not implemented"; return nil }

func (p *MethodLicense) Get() *license_get.Get { _ = "STUB: not implemented"; return nil }

func (p *MethodLicense) GetBasicStatus() *license_get_basic_status.GetBasicStatus {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodLicense) GetTrialStatus() *license_get_trial_status.GetTrialStatus {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodLicense) Post() *license_post.Post { _ = "STUB: not implemented"; return nil }

func (p *MethodLicense) PostStartBasic() *license_post_start_basic.PostStartBasic {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodLicense) PostStartTrial() *license_post_start_trial.PostStartTrial {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodLogstash) DeletePipeline(id string) *logstash_delete_pipeline.DeletePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodLogstash) GetPipeline() *logstash_get_pipeline.GetPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodLogstash) PutPipeline(id string) *logstash_put_pipeline.PutPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMigration) Deprecations() *migration_deprecations.Deprecations {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMigration) GetFeatureUpgradeStatus() *migration_get_feature_upgrade_status.GetFeatureUpgradeStatus {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMigration) PostFeatureUpgrade() *migration_post_feature_upgrade.PostFeatureUpgrade {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) ClearTrainedModelDeploymentCache(modelid string) *ml_clear_trained_model_deployment_cache.ClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) CloseJob(jobid string) *ml_close_job.CloseJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteCalendar(calendarid string) *ml_delete_calendar.DeleteCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteCalendarEvent(calendarid, eventid string) *ml_delete_calendar_event.DeleteCalendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteCalendarJob(calendarid, jobid string) *ml_delete_calendar_job.DeleteCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteDataFrameAnalytics(id string) *ml_delete_data_frame_analytics.DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteDatafeed(datafeedid string) *ml_delete_datafeed.DeleteDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteExpiredData() *ml_delete_expired_data.DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteFilter(filterid string) *ml_delete_filter.DeleteFilter {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteForecast(jobid string) *ml_delete_forecast.DeleteForecast {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteJob(jobid string) *ml_delete_job.DeleteJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteModelSnapshot(jobid, snapshotid string) *ml_delete_model_snapshot.DeleteModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteTrainedModel(modelid string) *ml_delete_trained_model.DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) DeleteTrainedModelAlias(modelid, modelalias string) *ml_delete_trained_model_alias.DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) EstimateModelMemory() *ml_estimate_model_memory.EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) EvaluateDataFrame() *ml_evaluate_data_frame.EvaluateDataFrame {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) ExplainDataFrameAnalytics() *ml_explain_data_frame_analytics.ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) FlushJob(jobid string) *ml_flush_job.FlushJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) Forecast(jobid string) *ml_forecast.Forecast {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetBuckets(jobid string) *ml_get_buckets.GetBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetCalendarEvents(calendarid string) *ml_get_calendar_events.GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetCalendars() *ml_get_calendars.GetCalendars {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetCategories(jobid string) *ml_get_categories.GetCategories {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetDataFrameAnalytics() *ml_get_data_frame_analytics.GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetDataFrameAnalyticsStats() *ml_get_data_frame_analytics_stats.GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetDatafeedStats() *ml_get_datafeed_stats.GetDatafeedStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetDatafeeds() *ml_get_datafeeds.GetDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetFilters() *ml_get_filters.GetFilters { _ = "STUB: not implemented"; return nil }

func (p *MethodMl) GetInfluencers(jobid string) *ml_get_influencers.GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetJobStats() *ml_get_job_stats.GetJobStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetJobs() *ml_get_jobs.GetJobs { _ = "STUB: not implemented"; return nil }

func (p *MethodMl) GetMemoryStats() *ml_get_memory_stats.GetMemoryStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetModelSnapshotUpgradeStats(jobid, snapshotid string) *ml_get_model_snapshot_upgrade_stats.GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetModelSnapshots(jobid string) *ml_get_model_snapshots.GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetOverallBuckets(jobid string) *ml_get_overall_buckets.GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetRecords(jobid string) *ml_get_records.GetRecords {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetTrainedModels() *ml_get_trained_models.GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) GetTrainedModelsStats() *ml_get_trained_models_stats.GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) InferTrainedModel(modelid string) *ml_infer_trained_model.InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) Info() *ml_info.Info { _ = "STUB: not implemented"; return nil }

func (p *MethodMl) OpenJob(jobid string) *ml_open_job.OpenJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PostCalendarEvents(calendarid string) *ml_post_calendar_events.PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PostData(jobid string) *ml_post_data.PostData {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PreviewDataFrameAnalytics() *ml_preview_data_frame_analytics.PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PreviewDatafeed() *ml_preview_datafeed.PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutCalendar(calendarid string) *ml_put_calendar.PutCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutCalendarJob(calendarid, jobid string) *ml_put_calendar_job.PutCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutDataFrameAnalytics(id string) *ml_put_data_frame_analytics.PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutDatafeed(datafeedid string) *ml_put_datafeed.PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutFilter(filterid string) *ml_put_filter.PutFilter {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutJob(jobid string) *ml_put_job.PutJob { _ = "STUB: not implemented"; return nil }

func (p *MethodMl) PutTrainedModel(modelid string) *ml_put_trained_model.PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutTrainedModelAlias(modelid, modelalias string) *ml_put_trained_model_alias.PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutTrainedModelDefinitionPart(modelid, part string) *ml_put_trained_model_definition_part.PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) PutTrainedModelVocabulary(modelid string) *ml_put_trained_model_vocabulary.PutTrainedModelVocabulary {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) ResetJob(jobid string) *ml_reset_job.ResetJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) RevertModelSnapshot(jobid, snapshotid string) *ml_revert_model_snapshot.RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) SetUpgradeMode() *ml_set_upgrade_mode.SetUpgradeMode {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) StartDataFrameAnalytics(id string) *ml_start_data_frame_analytics.StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) StartDatafeed(datafeedid string) *ml_start_datafeed.StartDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) StartTrainedModelDeployment(modelid string) *ml_start_trained_model_deployment.StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) StopDataFrameAnalytics(id string) *ml_stop_data_frame_analytics.StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) StopDatafeed(datafeedid string) *ml_stop_datafeed.StopDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) StopTrainedModelDeployment(modelid string) *ml_stop_trained_model_deployment.StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) UpdateDataFrameAnalytics(id string) *ml_update_data_frame_analytics.UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) UpdateDatafeed(datafeedid string) *ml_update_datafeed.UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) UpdateFilter(filterid string) *ml_update_filter.UpdateFilter {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) UpdateJob(jobid string) *ml_update_job.UpdateJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) UpdateModelSnapshot(jobid, snapshotid string) *ml_update_model_snapshot.UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) UpdateTrainedModelDeployment(modelid string) *ml_update_trained_model_deployment.UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) UpgradeJobSnapshot(jobid, snapshotid string) *ml_upgrade_job_snapshot.UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMl) Validate() *ml_validate.Validate { _ = "STUB: not implemented"; return nil }

func (p *MethodMl) ValidateDetector() *ml_validate_detector.ValidateDetector {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodMonitoring) Bulk() *monitoring_bulk.Bulk { _ = "STUB: not implemented"; return nil }

func (p *MethodNodes) ClearRepositoriesMeteringArchive(nodeid, maxarchiveversion string) *nodes_clear_repositories_metering_archive.ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodNodes) GetRepositoriesMeteringInfo(nodeid string) *nodes_get_repositories_metering_info.GetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodNodes) HotThreads() *nodes_hot_threads.HotThreads {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodNodes) Info() *nodes_info.Info { _ = "STUB: not implemented"; return nil }

func (p *MethodNodes) ReloadSecureSettings() *nodes_reload_secure_settings.ReloadSecureSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodNodes) Stats() *nodes_stats.Stats { _ = "STUB: not implemented"; return nil }

func (p *MethodNodes) Usage() *nodes_usage.Usage { _ = "STUB: not implemented"; return nil }

func (p *MethodProfiling) Flamegraph() *profiling_flamegraph.Flamegraph {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodProfiling) Stacktraces() *profiling_stacktraces.Stacktraces {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodProfiling) Status() *profiling_status.Status { _ = "STUB: not implemented"; return nil }

func (p *MethodProfiling) TopnFunctions() *profiling_topn_functions.TopnFunctions {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodProject) CreateManyRouting() *project_create_many_routing.CreateManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodProject) CreateRouting(name string) *project_create_routing.CreateRouting {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodProject) DeleteRouting(name string) *project_delete_routing.DeleteRouting {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodProject) GetManyRouting() *project_get_many_routing.GetManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodProject) GetRouting(name string) *project_get_routing.GetRouting {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodProject) Tags() *project_tags.Tags { _ = "STUB: not implemented"; return nil }

func (p *MethodQueryRules) DeleteRule(rulesetid, ruleid string) *query_rules_delete_rule.DeleteRule {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodQueryRules) DeleteRuleset(rulesetid string) *query_rules_delete_ruleset.DeleteRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodQueryRules) GetRule(rulesetid, ruleid string) *query_rules_get_rule.GetRule {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodQueryRules) GetRuleset(rulesetid string) *query_rules_get_ruleset.GetRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodQueryRules) ListRulesets() *query_rules_list_rulesets.ListRulesets {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodQueryRules) PutRule(rulesetid, ruleid string) *query_rules_put_rule.PutRule {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodQueryRules) PutRuleset(rulesetid string) *query_rules_put_ruleset.PutRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodQueryRules) Test(rulesetid string) *query_rules_test.Test {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodRollup) DeleteJob(id string) *rollup_delete_job.DeleteJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodRollup) GetJobs() *rollup_get_jobs.GetJobs { _ = "STUB: not implemented"; return nil }

func (p *MethodRollup) GetRollupCaps() *rollup_get_rollup_caps.GetRollupCaps {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodRollup) GetRollupIndexCaps(index string) *rollup_get_rollup_index_caps.GetRollupIndexCaps {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodRollup) PutJob(id string) *rollup_put_job.PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodRollup) RollupSearch(index string) *rollup_rollup_search.RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodRollup) StartJob(id string) *rollup_start_job.StartJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodRollup) StopJob(id string) *rollup_stop_job.StopJob {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) Delete(name string) *search_application_delete.Delete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) DeleteBehavioralAnalytics(name string) *search_application_delete_behavioral_analytics.DeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) Get(name string) *search_application_get.Get {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) GetBehavioralAnalytics() *search_application_get_behavioral_analytics.GetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) List() *search_application_list.List {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) PostBehavioralAnalyticsEvent(collectionname, eventtype string) *search_application_post_behavioral_analytics_event.PostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) Put(name string) *search_application_put.Put {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) PutBehavioralAnalytics(name string) *search_application_put_behavioral_analytics.PutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) RenderQuery(name string) *search_application_render_query.RenderQuery {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchApplication) Search(name string) *search_application_search.Search {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchableSnapshots) CacheStats() *searchable_snapshots_cache_stats.CacheStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchableSnapshots) ClearCache() *searchable_snapshots_clear_cache.ClearCache {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchableSnapshots) Mount(repository, snapshot string) *searchable_snapshots_mount.Mount {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSearchableSnapshots) Stats() *searchable_snapshots_stats.Stats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) ActivateUserProfile() *security_activate_user_profile.ActivateUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) Authenticate() *security_authenticate.Authenticate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) BulkDeleteRole() *security_bulk_delete_role.BulkDeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) BulkPutRole() *security_bulk_put_role.BulkPutRole {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) BulkUpdateApiKeys() *security_bulk_update_api_keys.BulkUpdateApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) ChangePassword() *security_change_password.ChangePassword {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) ClearApiKeyCache(ids string) *security_clear_api_key_cache.ClearApiKeyCache {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) ClearCachedPrivileges(application string) *security_clear_cached_privileges.ClearCachedPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) ClearCachedRealms(realms string) *security_clear_cached_realms.ClearCachedRealms {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) ClearCachedRoles(name string) *security_clear_cached_roles.ClearCachedRoles {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) ClearCachedServiceTokens(namespace, service, name string) *security_clear_cached_service_tokens.ClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) CloneApiKey() *security_clone_api_key.CloneApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) CreateApiKey() *security_create_api_key.CreateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) CreateCrossClusterApiKey() *security_create_cross_cluster_api_key.CreateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) CreateServiceToken(namespace, service string) *security_create_service_token.CreateServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) DelegatePki() *security_delegate_pki.DelegatePki {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) DeletePrivileges(application, name string) *security_delete_privileges.DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) DeleteRole(name string) *security_delete_role.DeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) DeleteRoleMapping(name string) *security_delete_role_mapping.DeleteRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) DeleteServiceToken(namespace, service, name string) *security_delete_service_token.DeleteServiceToken {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) DeleteUser(username string) *security_delete_user.DeleteUser {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) DisableUser(username string) *security_disable_user.DisableUser {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) DisableUserProfile(uid string) *security_disable_user_profile.DisableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) EnableUser(username string) *security_enable_user.EnableUser {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) EnableUserProfile(uid string) *security_enable_user_profile.EnableUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) EnrollKibana() *security_enroll_kibana.EnrollKibana {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) EnrollNode() *security_enroll_node.EnrollNode {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetApiKey() *security_get_api_key.GetApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetBuiltinPrivileges() *security_get_builtin_privileges.GetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetPrivileges() *security_get_privileges.GetPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetRole() *security_get_role.GetRole {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetRoleMapping() *security_get_role_mapping.GetRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetServiceAccounts() *security_get_service_accounts.GetServiceAccounts {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetServiceCredentials(namespace, service string) *security_get_service_credentials.GetServiceCredentials {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetSettings() *security_get_settings.GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetStats() *security_get_stats.GetStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetToken() *security_get_token.GetToken {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetUser() *security_get_user.GetUser {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetUserPrivileges() *security_get_user_privileges.GetUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GetUserProfile(uid string) *security_get_user_profile.GetUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) GrantApiKey() *security_grant_api_key.GrantApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) HasPrivileges() *security_has_privileges.HasPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) HasPrivilegesUserProfile() *security_has_privileges_user_profile.HasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) InvalidateApiKey() *security_invalidate_api_key.InvalidateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) InvalidateToken() *security_invalidate_token.InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) OidcAuthenticate() *security_oidc_authenticate.OidcAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) OidcLogout() *security_oidc_logout.OidcLogout {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) OidcPrepareAuthentication() *security_oidc_prepare_authentication.OidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) PutPrivileges() *security_put_privileges.PutPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) PutRole(name string) *security_put_role.PutRole {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) PutRoleMapping(name string) *security_put_role_mapping.PutRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) PutUser(username string) *security_put_user.PutUser {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) QueryApiKeys() *security_query_api_keys.QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) QueryRole() *security_query_role.QueryRole {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) QueryUser() *security_query_user.QueryUser {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) SamlAuthenticate() *security_saml_authenticate.SamlAuthenticate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) SamlCompleteLogout() *security_saml_complete_logout.SamlCompleteLogout {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) SamlInvalidate() *security_saml_invalidate.SamlInvalidate {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) SamlLogout() *security_saml_logout.SamlLogout {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) SamlPrepareAuthentication() *security_saml_prepare_authentication.SamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) SamlServiceProviderMetadata(realmname string) *security_saml_service_provider_metadata.SamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) SuggestUserProfiles() *security_suggest_user_profiles.SuggestUserProfiles {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) UpdateApiKey(id string) *security_update_api_key.UpdateApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) UpdateCrossClusterApiKey(id string) *security_update_cross_cluster_api_key.UpdateCrossClusterApiKey {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) UpdateSettings() *security_update_settings.UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSecurity) UpdateUserProfileData(uid string) *security_update_user_profile_data.UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodShutdown) DeleteNode(nodeid string) *shutdown_delete_node.DeleteNode {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodShutdown) GetNode() *shutdown_get_node.GetNode {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodShutdown) PutNode(nodeid string) *shutdown_put_node.PutNode {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSimulate) Ingest() *simulate_ingest.Ingest { _ = "STUB: not implemented"; return nil }

func (p *MethodSlm) DeleteLifecycle(policyid string) *slm_delete_lifecycle.DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSlm) ExecuteLifecycle(policyid string) *slm_execute_lifecycle.ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSlm) ExecuteRetention() *slm_execute_retention.ExecuteRetention {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSlm) GetLifecycle() *slm_get_lifecycle.GetLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSlm) GetStats() *slm_get_stats.GetStats { _ = "STUB: not implemented"; return nil }

func (p *MethodSlm) GetStatus() *slm_get_status.GetStatus { _ = "STUB: not implemented"; return nil }

func (p *MethodSlm) PutLifecycle(policyid string) *slm_put_lifecycle.PutLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSlm) Start() *slm_start.Start { _ = "STUB: not implemented"; return nil }

func (p *MethodSlm) Stop() *slm_stop.Stop { _ = "STUB: not implemented"; return nil }

func (p *MethodSnapshot) CleanupRepository(repository string) *snapshot_cleanup_repository.CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) Clone(repository, snapshot, targetsnapshot string) *snapshot_clone.Clone {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) Create(repository, snapshot string) *snapshot_create.Create {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) CreateRepository(repository string) *snapshot_create_repository.CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) Delete(repository, snapshot string) *snapshot_delete.Delete {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) DeleteRepository(repository string) *snapshot_delete_repository.DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) Get(repository, snapshot string) *snapshot_get.Get {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) GetRepository() *snapshot_get_repository.GetRepository {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) RepositoryAnalyze(repository string) *snapshot_repository_analyze.RepositoryAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) RepositoryVerifyIntegrity(repository string) *snapshot_repository_verify_integrity.RepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) Restore(repository, snapshot string) *snapshot_restore.Restore {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSnapshot) Status() *snapshot_status.Status { _ = "STUB: not implemented"; return nil }

func (p *MethodSnapshot) VerifyRepository(repository string) *snapshot_verify_repository.VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSql) ClearCursor() *sql_clear_cursor.ClearCursor {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSql) DeleteAsync(id string) *sql_delete_async.DeleteAsync {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSql) GetAsync(id string) *sql_get_async.GetAsync {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSql) GetAsyncStatus(id string) *sql_get_async_status.GetAsyncStatus {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSql) Query() *sql_query.Query { _ = "STUB: not implemented"; return nil }

func (p *MethodSql) Translate() *sql_translate.Translate { _ = "STUB: not implemented"; return nil }

func (p *MethodSsl) Certificates() *ssl_certificates.Certificates {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodStreams) LogsDisable(name string) *streams_logs_disable.LogsDisable {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodStreams) LogsEnable(name string) *streams_logs_enable.LogsEnable {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodStreams) Status() *streams_status.Status { _ = "STUB: not implemented"; return nil }

func (p *MethodSynonyms) DeleteSynonym(id string) *synonyms_delete_synonym.DeleteSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSynonyms) DeleteSynonymRule(setid, ruleid string) *synonyms_delete_synonym_rule.DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSynonyms) GetSynonym(id string) *synonyms_get_synonym.GetSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSynonyms) GetSynonymRule(setid, ruleid string) *synonyms_get_synonym_rule.GetSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSynonyms) GetSynonymsSets() *synonyms_get_synonyms_sets.GetSynonymsSets {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSynonyms) PutSynonym(id string) *synonyms_put_synonym.PutSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodSynonyms) PutSynonymRule(setid, ruleid string) *synonyms_put_synonym_rule.PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTasks) Cancel() *tasks_cancel.Cancel { _ = "STUB: not implemented"; return nil }

func (p *MethodTasks) Get(taskid string) *tasks_get.Get { _ = "STUB: not implemented"; return nil }

func (p *MethodTasks) List() *tasks_list.List { _ = "STUB: not implemented"; return nil }

func (p *MethodTextStructure) FindFieldStructure() *text_structure_find_field_structure.FindFieldStructure {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTextStructure) FindMessageStructure() *text_structure_find_message_structure.FindMessageStructure {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTextStructure) FindStructure() *text_structure_find_structure.FindStructure {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTextStructure) TestGrokPattern() *text_structure_test_grok_pattern.TestGrokPattern {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) DeleteTransform(transformid string) *transform_delete_transform.DeleteTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) GetNodeStats() *transform_get_node_stats.GetNodeStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) GetTransform() *transform_get_transform.GetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) GetTransformStats(transformid string) *transform_get_transform_stats.GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) PreviewTransform() *transform_preview_transform.PreviewTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) PutTransform(transformid string) *transform_put_transform.PutTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) ResetTransform(transformid string) *transform_reset_transform.ResetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) ScheduleNowTransform(transformid string) *transform_schedule_now_transform.ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) SetUpgradeMode() *transform_set_upgrade_mode.SetUpgradeMode {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) StartTransform(transformid string) *transform_start_transform.StartTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) StopTransform(transformid string) *transform_stop_transform.StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) UpdateTransform(transformid string) *transform_update_transform.UpdateTransform {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodTransform) UpgradeTransforms() *transform_upgrade_transforms.UpgradeTransforms {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) AckWatch(watchid string) *watcher_ack_watch.AckWatch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) ActivateWatch(watchid string) *watcher_activate_watch.ActivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) DeactivateWatch(watchid string) *watcher_deactivate_watch.DeactivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) DeleteWatch(id string) *watcher_delete_watch.DeleteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) ExecuteWatch() *watcher_execute_watch.ExecuteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) GetSettings() *watcher_get_settings.GetSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) GetWatch(id string) *watcher_get_watch.GetWatch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) PutWatch(id string) *watcher_put_watch.PutWatch {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) QueryWatches() *watcher_query_watches.QueryWatches {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodWatcher) Start() *watcher_start.Start { _ = "STUB: not implemented"; return nil }

func (p *MethodWatcher) Stats() *watcher_stats.Stats { _ = "STUB: not implemented"; return nil }

func (p *MethodWatcher) Stop() *watcher_stop.Stop { _ = "STUB: not implemented"; return nil }

func (p *MethodWatcher) UpdateSettings() *watcher_update_settings.UpdateSettings {
	_ = "STUB: not implemented"
	return nil
}

func (p *MethodXpack) Info() *xpack_info.Info { _ = "STUB: not implemented"; return nil }

func (p *MethodXpack) Usage() *xpack_usage.Usage { _ = "STUB: not implemented"; return nil }

func NewMethodAPI(tp elastictransport.Interface) *MethodAPI { _ = "STUB: not implemented"; return nil }
