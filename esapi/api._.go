package esapi

type API struct {
	Cat         *Cat
	Cluster     *Cluster
	Indices     *Indices
	Ingest      *Ingest
	Nodes       *Nodes
	Remote      *Remote
	Snapshot    *Snapshot
	Tasks       *Tasks
	AsyncSearch *AsyncSearch
	CCR         *CCR
	ILM         *ILM
	License     *License
	Migration   *Migration
	ML          *ML
	Monitoring  *Monitoring
	Rollup      *Rollup
	Security    *Security
	SQL         *SQL
	SSL         *SSL
	Watcher     *Watcher
	XPack       *XPack

	AutoscalingDeleteAutoscalingPolicy            AutoscalingDeleteAutoscalingPolicy
	AutoscalingGetAutoscalingCapacity             AutoscalingGetAutoscalingCapacity
	AutoscalingGetAutoscalingPolicy               AutoscalingGetAutoscalingPolicy
	AutoscalingPutAutoscalingPolicy               AutoscalingPutAutoscalingPolicy
	Bulk                                          Bulk
	Capabilities                                  Capabilities
	ClearScroll                                   ClearScroll
	ClosePointInTime                              ClosePointInTime
	ConnectorCheckIn                              ConnectorCheckIn
	ConnectorDelete                               ConnectorDelete
	ConnectorGet                                  ConnectorGet
	ConnectorLastSync                             ConnectorLastSync
	ConnectorList                                 ConnectorList
	ConnectorPost                                 ConnectorPost
	ConnectorPut                                  ConnectorPut
	ConnectorSecretDelete                         ConnectorSecretDelete
	ConnectorSecretGet                            ConnectorSecretGet
	ConnectorSecretPost                           ConnectorSecretPost
	ConnectorSecretPut                            ConnectorSecretPut
	ConnectorSyncJobCancel                        ConnectorSyncJobCancel
	ConnectorSyncJobCheckIn                       ConnectorSyncJobCheckIn
	ConnectorSyncJobClaim                         ConnectorSyncJobClaim
	ConnectorSyncJobDelete                        ConnectorSyncJobDelete
	ConnectorSyncJobError                         ConnectorSyncJobError
	ConnectorSyncJobGet                           ConnectorSyncJobGet
	ConnectorSyncJobList                          ConnectorSyncJobList
	ConnectorSyncJobPost                          ConnectorSyncJobPost
	ConnectorSyncJobUpdateStats                   ConnectorSyncJobUpdateStats
	ConnectorUpdateAPIKeyDocumentID               ConnectorUpdateAPIKeyDocumentID
	ConnectorUpdateActiveFiltering                ConnectorUpdateActiveFiltering
	ConnectorUpdateConfiguration                  ConnectorUpdateConfiguration
	ConnectorUpdateError                          ConnectorUpdateError
	ConnectorUpdateFeatures                       ConnectorUpdateFeatures
	ConnectorUpdateFiltering                      ConnectorUpdateFiltering
	ConnectorUpdateFilteringValidation            ConnectorUpdateFilteringValidation
	ConnectorUpdateIndexName                      ConnectorUpdateIndexName
	ConnectorUpdateName                           ConnectorUpdateName
	ConnectorUpdateNative                         ConnectorUpdateNative
	ConnectorUpdatePipeline                       ConnectorUpdatePipeline
	ConnectorUpdateScheduling                     ConnectorUpdateScheduling
	ConnectorUpdateServiceDocumentType            ConnectorUpdateServiceDocumentType
	ConnectorUpdateStatus                         ConnectorUpdateStatus
	Count                                         Count
	Create                                        Create
	DanglingIndicesDeleteDanglingIndex            DanglingIndicesDeleteDanglingIndex
	DanglingIndicesImportDanglingIndex            DanglingIndicesImportDanglingIndex
	DanglingIndicesListDanglingIndices            DanglingIndicesListDanglingIndices
	DeleteByQuery                                 DeleteByQuery
	DeleteByQueryRethrottle                       DeleteByQueryRethrottle
	Delete                                        Delete
	DeleteScript                                  DeleteScript
	EnrichDeletePolicy                            EnrichDeletePolicy
	EnrichExecutePolicy                           EnrichExecutePolicy
	EnrichGetPolicy                               EnrichGetPolicy
	EnrichPutPolicy                               EnrichPutPolicy
	EnrichStats                                   EnrichStats
	EqlDelete                                     EqlDelete
	EqlGet                                        EqlGet
	EqlGetStatus                                  EqlGetStatus
	EqlSearch                                     EqlSearch
	EsqlAsyncQueryDelete                          EsqlAsyncQueryDelete
	EsqlAsyncQueryGet                             EsqlAsyncQueryGet
	EsqlAsyncQuery                                EsqlAsyncQuery
	EsqlAsyncQueryStop                            EsqlAsyncQueryStop
	EsqlDeleteView                                EsqlDeleteView
	EsqlGetQuery                                  EsqlGetQuery
	EsqlGetView                                   EsqlGetView
	EsqlListQueries                               EsqlListQueries
	EsqlPutView                                   EsqlPutView
	EsqlQuery                                     EsqlQuery
	Exists                                        Exists
	ExistsSource                                  ExistsSource
	Explain                                       Explain
	FeaturesGetFeatures                           FeaturesGetFeatures
	FeaturesResetFeatures                         FeaturesResetFeatures
	FieldCaps                                     FieldCaps
	FleetDeleteSecret                             FleetDeleteSecret
	FleetGetSecret                                FleetGetSecret
	FleetGlobalCheckpoints                        FleetGlobalCheckpoints
	FleetMsearch                                  FleetMsearch
	FleetPostSecret                               FleetPostSecret
	FleetSearch                                   FleetSearch
	Get                                           Get
	GetScriptContext                              GetScriptContext
	GetScriptLanguages                            GetScriptLanguages
	GetScript                                     GetScript
	GetSource                                     GetSource
	GraphExplore                                  GraphExplore
	HealthReport                                  HealthReport
	Index                                         Index
	InferenceChatCompletionUnified                InferenceChatCompletionUnified
	InferenceCompletion                           InferenceCompletion
	InferenceDelete                               InferenceDelete
	InferenceEmbedding                            InferenceEmbedding
	InferenceGet                                  InferenceGet
	InferenceInference                            InferenceInference
	InferencePutAi21                              InferencePutAi21
	InferencePutAlibabacloud                      InferencePutAlibabacloud
	InferencePutAmazonbedrock                     InferencePutAmazonbedrock
	InferencePutAmazonsagemaker                   InferencePutAmazonsagemaker
	InferencePutAnthropic                         InferencePutAnthropic
	InferencePutAzureaistudio                     InferencePutAzureaistudio
	InferencePutAzureopenai                       InferencePutAzureopenai
	InferencePutCohere                            InferencePutCohere
	InferencePutContextualai                      InferencePutContextualai
	InferencePutCustom                            InferencePutCustom
	InferencePutDeepseek                          InferencePutDeepseek
	InferencePutElasticsearch                     InferencePutElasticsearch
	InferencePutElser                             InferencePutElser
	InferencePutFireworksai                       InferencePutFireworksai
	InferencePutGoogleaistudio                    InferencePutGoogleaistudio
	InferencePutGooglevertexai                    InferencePutGooglevertexai
	InferencePutGroq                              InferencePutGroq
	InferencePutHuggingFace                       InferencePutHuggingFace
	InferencePutJinaai                            InferencePutJinaai
	InferencePutLlama                             InferencePutLlama
	InferencePutMistral                           InferencePutMistral
	InferencePutNvidia                            InferencePutNvidia
	InferencePutOpenai                            InferencePutOpenai
	InferencePutOpenshiftAi                       InferencePutOpenshiftAi
	InferencePut                                  InferencePut
	InferencePutVoyageai                          InferencePutVoyageai
	InferencePutWatsonx                           InferencePutWatsonx
	InferenceRerank                               InferenceRerank
	InferenceSparseEmbedding                      InferenceSparseEmbedding
	InferenceStreamCompletion                     InferenceStreamCompletion
	InferenceTextEmbedding                        InferenceTextEmbedding
	InferenceUpdate                               InferenceUpdate
	Info                                          Info
	LogstashDeletePipeline                        LogstashDeletePipeline
	LogstashGetPipeline                           LogstashGetPipeline
	LogstashPutPipeline                           LogstashPutPipeline
	Mget                                          Mget
	Msearch                                       Msearch
	MsearchTemplate                               MsearchTemplate
	Mtermvectors                                  Mtermvectors
	OpenPointInTime                               OpenPointInTime
	Ping                                          Ping
	ProfilingFlamegraph                           ProfilingFlamegraph
	ProfilingStacktraces                          ProfilingStacktraces
	ProfilingStatus                               ProfilingStatus
	ProfilingTopnFunctions                        ProfilingTopnFunctions
	ProjectCreateManyRouting                      ProjectCreateManyRouting
	ProjectCreateRouting                          ProjectCreateRouting
	ProjectDeleteRouting                          ProjectDeleteRouting
	ProjectGetManyRouting                         ProjectGetManyRouting
	ProjectGetRouting                             ProjectGetRouting
	ProjectTags                                   ProjectTags
	PutScript                                     PutScript
	QueryRulesDeleteRule                          QueryRulesDeleteRule
	QueryRulesDeleteRuleset                       QueryRulesDeleteRuleset
	QueryRulesGetRule                             QueryRulesGetRule
	QueryRulesGetRuleset                          QueryRulesGetRuleset
	QueryRulesListRulesets                        QueryRulesListRulesets
	QueryRulesPutRule                             QueryRulesPutRule
	QueryRulesPutRuleset                          QueryRulesPutRuleset
	QueryRulesTest                                QueryRulesTest
	RankEval                                      RankEval
	ReindexCancel                                 ReindexCancel
	ReindexGet                                    ReindexGet
	ReindexList                                   ReindexList
	Reindex                                       Reindex
	ReindexRethrottle                             ReindexRethrottle
	RenderSearchTemplate                          RenderSearchTemplate
	ScriptsPainlessExecute                        ScriptsPainlessExecute
	Scroll                                        Scroll
	SearchApplicationDeleteBehavioralAnalytics    SearchApplicationDeleteBehavioralAnalytics
	SearchApplicationDelete                       SearchApplicationDelete
	SearchApplicationGetBehavioralAnalytics       SearchApplicationGetBehavioralAnalytics
	SearchApplicationGet                          SearchApplicationGet
	SearchApplicationList                         SearchApplicationList
	SearchApplicationPostBehavioralAnalyticsEvent SearchApplicationPostBehavioralAnalyticsEvent
	SearchApplicationPutBehavioralAnalytics       SearchApplicationPutBehavioralAnalytics
	SearchApplicationPut                          SearchApplicationPut
	SearchApplicationRenderQuery                  SearchApplicationRenderQuery
	SearchApplicationSearch                       SearchApplicationSearch
	SearchMvt                                     SearchMvt
	Search                                        Search
	SearchShards                                  SearchShards
	SearchTemplate                                SearchTemplate
	SearchableSnapshotsCacheStats                 SearchableSnapshotsCacheStats
	SearchableSnapshotsClearCache                 SearchableSnapshotsClearCache
	SearchableSnapshotsMount                      SearchableSnapshotsMount
	SearchableSnapshotsStats                      SearchableSnapshotsStats
	ShutdownDeleteNode                            ShutdownDeleteNode
	ShutdownGetNode                               ShutdownGetNode
	ShutdownPutNode                               ShutdownPutNode
	SimulateIngest                                SimulateIngest
	SlmDeleteLifecycle                            SlmDeleteLifecycle
	SlmExecuteLifecycle                           SlmExecuteLifecycle
	SlmExecuteRetention                           SlmExecuteRetention
	SlmGetLifecycle                               SlmGetLifecycle
	SlmGetStats                                   SlmGetStats
	SlmGetStatus                                  SlmGetStatus
	SlmPutLifecycle                               SlmPutLifecycle
	SlmStart                                      SlmStart
	SlmStop                                       SlmStop
	StreamsLogsDisable                            StreamsLogsDisable
	StreamsLogsEnable                             StreamsLogsEnable
	StreamsStatus                                 StreamsStatus
	SynonymsDeleteSynonym                         SynonymsDeleteSynonym
	SynonymsDeleteSynonymRule                     SynonymsDeleteSynonymRule
	SynonymsGetSynonym                            SynonymsGetSynonym
	SynonymsGetSynonymRule                        SynonymsGetSynonymRule
	SynonymsGetSynonymsSets                       SynonymsGetSynonymsSets
	SynonymsPutSynonym                            SynonymsPutSynonym
	SynonymsPutSynonymRule                        SynonymsPutSynonymRule
	TermsEnum                                     TermsEnum
	Termvectors                                   Termvectors
	TextStructureFindFieldStructure               TextStructureFindFieldStructure
	TextStructureFindMessageStructure             TextStructureFindMessageStructure
	TextStructureFindStructure                    TextStructureFindStructure
	TextStructureTestGrokPattern                  TextStructureTestGrokPattern
	TransformDeleteTransform                      TransformDeleteTransform
	TransformGetNodeStats                         TransformGetNodeStats
	TransformGetTransform                         TransformGetTransform
	TransformGetTransformStats                    TransformGetTransformStats
	TransformPreviewTransform                     TransformPreviewTransform
	TransformPutTransform                         TransformPutTransform
	TransformResetTransform                       TransformResetTransform
	TransformScheduleNowTransform                 TransformScheduleNowTransform
	TransformSetUpgradeMode                       TransformSetUpgradeMode
	TransformStartTransform                       TransformStartTransform
	TransformStopTransform                        TransformStopTransform
	TransformUpdateTransform                      TransformUpdateTransform
	TransformUpgradeTransforms                    TransformUpgradeTransforms
	UpdateByQuery                                 UpdateByQuery
	UpdateByQueryRethrottle                       UpdateByQueryRethrottle
	Update                                        Update
}

type Cat struct {
	Aliases              CatAliases
	Allocation           CatAllocation
	CircuitBreaker       CatCircuitBreaker
	ComponentTemplates   CatComponentTemplates
	Count                CatCount
	Fielddata            CatFielddata
	Health               CatHealth
	Help                 CatHelp
	Indices              CatIndices
	MLDataFrameAnalytics CatMLDataFrameAnalytics
	MLDatafeeds          CatMLDatafeeds
	MLJobs               CatMLJobs
	MLTrainedModels      CatMLTrainedModels
	Master               CatMaster
	Nodeattrs            CatNodeattrs
	Nodes                CatNodes
	PendingTasks         CatPendingTasks
	Plugins              CatPlugins
	Recovery             CatRecovery
	Repositories         CatRepositories
	Segments             CatSegments
	Shards               CatShards
	Snapshots            CatSnapshots
	Tasks                CatTasks
	Templates            CatTemplates
	ThreadPool           CatThreadPool
	Transforms           CatTransforms
}

type Cluster struct {
	AllocationExplain            ClusterAllocationExplain
	DeleteComponentTemplate      ClusterDeleteComponentTemplate
	DeleteVotingConfigExclusions ClusterDeleteVotingConfigExclusions
	ExistsComponentTemplate      ClusterExistsComponentTemplate
	GetComponentTemplate         ClusterGetComponentTemplate
	GetSettings                  ClusterGetSettings
	Health                       ClusterHealth
	Info                         ClusterInfo
	PendingTasks                 ClusterPendingTasks
	PostVotingConfigExclusions   ClusterPostVotingConfigExclusions
	PutComponentTemplate         ClusterPutComponentTemplate
	PutSettings                  ClusterPutSettings
	RemoteInfo                   ClusterRemoteInfo
	Reroute                      ClusterReroute
	State                        ClusterState
	Stats                        ClusterStats
}

type Indices struct {
	AddBlock                IndicesAddBlock
	Analyze                 IndicesAnalyze
	CancelMigrateReindex    IndicesCancelMigrateReindex
	ClearCache              IndicesClearCache
	Clone                   IndicesClone
	Close                   IndicesClose
	CreateDataStream        IndicesCreateDataStream
	CreateFrom              IndicesCreateFrom
	Create                  IndicesCreate
	DataStreamsStats        IndicesDataStreamsStats
	DeleteAlias             IndicesDeleteAlias
	DeleteDataLifecycle     IndicesDeleteDataLifecycle
	DeleteDataStreamOptions IndicesDeleteDataStreamOptions
	DeleteDataStream        IndicesDeleteDataStream
	DeleteIndexTemplate     IndicesDeleteIndexTemplate
	Delete                  IndicesDelete
	DeleteTemplate          IndicesDeleteTemplate
	DiskUsage               IndicesDiskUsage
	Downsample              IndicesDownsample
	ExistsAlias             IndicesExistsAlias
	ExistsIndexTemplate     IndicesExistsIndexTemplate
	Exists                  IndicesExists
	ExistsTemplate          IndicesExistsTemplate
	ExplainDataLifecycle    IndicesExplainDataLifecycle
	FieldUsageStats         IndicesFieldUsageStats
	Flush                   IndicesFlush
	Forcemerge              IndicesForcemerge
	GetAlias                IndicesGetAlias
	GetDataLifecycle        IndicesGetDataLifecycle
	GetDataLifecycleStats   IndicesGetDataLifecycleStats
	GetDataStreamMappings   IndicesGetDataStreamMappings
	GetDataStreamOptions    IndicesGetDataStreamOptions
	GetDataStream           IndicesGetDataStream
	GetDataStreamSettings   IndicesGetDataStreamSettings
	GetFieldMapping         IndicesGetFieldMapping
	GetIndexTemplate        IndicesGetIndexTemplate
	GetMapping              IndicesGetMapping
	GetMigrateReindexStatus IndicesGetMigrateReindexStatus
	Get                     IndicesGet
	GetSettings             IndicesGetSettings
	GetTemplate             IndicesGetTemplate
	MigrateReindex          IndicesMigrateReindex
	MigrateToDataStream     IndicesMigrateToDataStream
	ModifyDataStream        IndicesModifyDataStream
	Open                    IndicesOpen
	PromoteDataStream       IndicesPromoteDataStream
	PutAlias                IndicesPutAlias
	PutDataLifecycle        IndicesPutDataLifecycle
	PutDataStreamMappings   IndicesPutDataStreamMappings
	PutDataStreamOptions    IndicesPutDataStreamOptions
	PutDataStreamSettings   IndicesPutDataStreamSettings
	PutIndexTemplate        IndicesPutIndexTemplate
	PutMapping              IndicesPutMapping
	PutSettings             IndicesPutSettings
	PutTemplate             IndicesPutTemplate
	Recovery                IndicesRecovery
	Refresh                 IndicesRefresh
	ReloadSearchAnalyzers   IndicesReloadSearchAnalyzers
	RemoveBlock             IndicesRemoveBlock
	ResolveCluster          IndicesResolveCluster
	ResolveIndex            IndicesResolveIndex
	Rollover                IndicesRollover
	Segments                IndicesSegments
	ShardStores             IndicesShardStores
	Shrink                  IndicesShrink
	SimulateIndexTemplate   IndicesSimulateIndexTemplate
	SimulateTemplate        IndicesSimulateTemplate
	Split                   IndicesSplit
	Stats                   IndicesStats
	UpdateAliases           IndicesUpdateAliases
	ValidateQuery           IndicesValidateQuery
}

type Ingest struct {
	DeleteGeoipDatabase      IngestDeleteGeoipDatabase
	DeleteIPLocationDatabase IngestDeleteIPLocationDatabase
	DeletePipeline           IngestDeletePipeline
	GeoIPStats               IngestGeoIPStats
	GetGeoipDatabase         IngestGetGeoipDatabase
	GetIPLocationDatabase    IngestGetIPLocationDatabase
	GetPipeline              IngestGetPipeline
	ProcessorGrok            IngestProcessorGrok
	PutGeoipDatabase         IngestPutGeoipDatabase
	PutIPLocationDatabase    IngestPutIPLocationDatabase
	PutPipeline              IngestPutPipeline
	Simulate                 IngestSimulate
}

type Nodes struct {
	ClearRepositoriesMeteringArchive NodesClearRepositoriesMeteringArchive
	GetRepositoriesMeteringInfo      NodesGetRepositoriesMeteringInfo
	HotThreads                       NodesHotThreads
	Info                             NodesInfo
	ReloadSecureSettings             NodesReloadSecureSettings
	Stats                            NodesStats
	Usage                            NodesUsage
}

type Remote struct {
}

type Snapshot struct {
	CleanupRepository         SnapshotCleanupRepository
	Clone                     SnapshotClone
	CreateRepository          SnapshotCreateRepository
	Create                    SnapshotCreate
	DeleteRepository          SnapshotDeleteRepository
	Delete                    SnapshotDelete
	GetRepository             SnapshotGetRepository
	Get                       SnapshotGet
	RepositoryAnalyze         SnapshotRepositoryAnalyze
	RepositoryVerifyIntegrity SnapshotRepositoryVerifyIntegrity
	Restore                   SnapshotRestore
	Status                    SnapshotStatus
	VerifyRepository          SnapshotVerifyRepository
}

type Tasks struct {
	Cancel TasksCancel
	Get    TasksGet
	List   TasksList
}

type AsyncSearch struct {
	Delete AsyncSearchDelete
	Get    AsyncSearchGet
	Status AsyncSearchStatus
	Submit AsyncSearchSubmit
}

type CCR struct {
	DeleteAutoFollowPattern CCRDeleteAutoFollowPattern
	FollowInfo              CCRFollowInfo
	Follow                  CCRFollow
	FollowStats             CCRFollowStats
	ForgetFollower          CCRForgetFollower
	GetAutoFollowPattern    CCRGetAutoFollowPattern
	PauseAutoFollowPattern  CCRPauseAutoFollowPattern
	PauseFollow             CCRPauseFollow
	PutAutoFollowPattern    CCRPutAutoFollowPattern
	ResumeAutoFollowPattern CCRResumeAutoFollowPattern
	ResumeFollow            CCRResumeFollow
	Stats                   CCRStats
	Unfollow                CCRUnfollow
}

type ILM struct {
	DeleteLifecycle    ILMDeleteLifecycle
	ExplainLifecycle   ILMExplainLifecycle
	GetLifecycle       ILMGetLifecycle
	GetStatus          ILMGetStatus
	MigrateToDataTiers ILMMigrateToDataTiers
	MoveToStep         ILMMoveToStep
	PutLifecycle       ILMPutLifecycle
	RemovePolicy       ILMRemovePolicy
	Retry              ILMRetry
	Start              ILMStart
	Stop               ILMStop
}

type License struct {
	Delete         LicenseDelete
	GetBasicStatus LicenseGetBasicStatus
	Get            LicenseGet
	GetTrialStatus LicenseGetTrialStatus
	Post           LicensePost
	PostStartBasic LicensePostStartBasic
	PostStartTrial LicensePostStartTrial
}

type Migration struct {
	Deprecations            MigrationDeprecations
	GetFeatureUpgradeStatus MigrationGetFeatureUpgradeStatus
	PostFeatureUpgrade      MigrationPostFeatureUpgrade
}

type ML struct {
	ClearTrainedModelDeploymentCache MLClearTrainedModelDeploymentCache
	CloseJob                         MLCloseJob
	DeleteCalendarEvent              MLDeleteCalendarEvent
	DeleteCalendarJob                MLDeleteCalendarJob
	DeleteCalendar                   MLDeleteCalendar
	DeleteDataFrameAnalytics         MLDeleteDataFrameAnalytics
	DeleteDatafeed                   MLDeleteDatafeed
	DeleteExpiredData                MLDeleteExpiredData
	DeleteFilter                     MLDeleteFilter
	DeleteForecast                   MLDeleteForecast
	DeleteJob                        MLDeleteJob
	DeleteModelSnapshot              MLDeleteModelSnapshot
	DeleteTrainedModelAlias          MLDeleteTrainedModelAlias
	DeleteTrainedModel               MLDeleteTrainedModel
	EstimateModelMemory              MLEstimateModelMemory
	EvaluateDataFrame                MLEvaluateDataFrame
	ExplainDataFrameAnalytics        MLExplainDataFrameAnalytics
	FlushJob                         MLFlushJob
	Forecast                         MLForecast
	GetBuckets                       MLGetBuckets
	GetCalendarEvents                MLGetCalendarEvents
	GetCalendars                     MLGetCalendars
	GetCategories                    MLGetCategories
	GetDataFrameAnalytics            MLGetDataFrameAnalytics
	GetDataFrameAnalyticsStats       MLGetDataFrameAnalyticsStats
	GetDatafeedStats                 MLGetDatafeedStats
	GetDatafeeds                     MLGetDatafeeds
	GetFilters                       MLGetFilters
	GetInfluencers                   MLGetInfluencers
	GetJobStats                      MLGetJobStats
	GetJobs                          MLGetJobs
	GetMemoryStats                   MLGetMemoryStats
	GetModelSnapshotUpgradeStats     MLGetModelSnapshotUpgradeStats
	GetModelSnapshots                MLGetModelSnapshots
	GetOverallBuckets                MLGetOverallBuckets
	GetRecords                       MLGetRecords
	GetTrainedModels                 MLGetTrainedModels
	GetTrainedModelsStats            MLGetTrainedModelsStats
	InferTrainedModel                MLInferTrainedModel
	Info                             MLInfo
	OpenJob                          MLOpenJob
	PostCalendarEvents               MLPostCalendarEvents
	PostData                         MLPostData
	PreviewDataFrameAnalytics        MLPreviewDataFrameAnalytics
	PreviewDatafeed                  MLPreviewDatafeed
	PutCalendarJob                   MLPutCalendarJob
	PutCalendar                      MLPutCalendar
	PutDataFrameAnalytics            MLPutDataFrameAnalytics
	PutDatafeed                      MLPutDatafeed
	PutFilter                        MLPutFilter
	PutJob                           MLPutJob
	PutTrainedModelAlias             MLPutTrainedModelAlias
	PutTrainedModelDefinitionPart    MLPutTrainedModelDefinitionPart
	PutTrainedModel                  MLPutTrainedModel
	PutTrainedModelVocabulary        MLPutTrainedModelVocabulary
	ResetJob                         MLResetJob
	RevertModelSnapshot              MLRevertModelSnapshot
	SetUpgradeMode                   MLSetUpgradeMode
	StartDataFrameAnalytics          MLStartDataFrameAnalytics
	StartDatafeed                    MLStartDatafeed
	StartTrainedModelDeployment      MLStartTrainedModelDeployment
	StopDataFrameAnalytics           MLStopDataFrameAnalytics
	StopDatafeed                     MLStopDatafeed
	StopTrainedModelDeployment       MLStopTrainedModelDeployment
	UpdateDataFrameAnalytics         MLUpdateDataFrameAnalytics
	UpdateDatafeed                   MLUpdateDatafeed
	UpdateFilter                     MLUpdateFilter
	UpdateJob                        MLUpdateJob
	UpdateModelSnapshot              MLUpdateModelSnapshot
	UpdateTrainedModelDeployment     MLUpdateTrainedModelDeployment
	UpgradeJobSnapshot               MLUpgradeJobSnapshot
	ValidateDetector                 MLValidateDetector
	Validate                         MLValidate
}

type Monitoring struct {
	Bulk MonitoringBulk
}

type Rollup struct {
	DeleteJob    RollupDeleteJob
	GetJobs      RollupGetJobs
	GetCaps      RollupGetRollupCaps
	GetIndexCaps RollupGetRollupIndexCaps
	PutJob       RollupPutJob
	Search       RollupRollupSearch
	StartJob     RollupStartJob
	StopJob      RollupStopJob
}

type Security struct {
	ActivateUserProfile         SecurityActivateUserProfile
	Authenticate                SecurityAuthenticate
	BulkDeleteRole              SecurityBulkDeleteRole
	BulkPutRole                 SecurityBulkPutRole
	BulkUpdateAPIKeys           SecurityBulkUpdateAPIKeys
	ChangePassword              SecurityChangePassword
	ClearAPIKeyCache            SecurityClearAPIKeyCache
	ClearCachedPrivileges       SecurityClearCachedPrivileges
	ClearCachedRealms           SecurityClearCachedRealms
	ClearCachedRoles            SecurityClearCachedRoles
	ClearCachedServiceTokens    SecurityClearCachedServiceTokens
	CloneAPIKey                 SecurityCloneAPIKey
	CreateAPIKey                SecurityCreateAPIKey
	CreateCrossClusterAPIKey    SecurityCreateCrossClusterAPIKey
	CreateServiceToken          SecurityCreateServiceToken
	DelegatePki                 SecurityDelegatePki
	DeletePrivileges            SecurityDeletePrivileges
	DeleteRoleMapping           SecurityDeleteRoleMapping
	DeleteRole                  SecurityDeleteRole
	DeleteServiceToken          SecurityDeleteServiceToken
	DeleteUser                  SecurityDeleteUser
	DisableUserProfile          SecurityDisableUserProfile
	DisableUser                 SecurityDisableUser
	EnableUserProfile           SecurityEnableUserProfile
	EnableUser                  SecurityEnableUser
	EnrollKibana                SecurityEnrollKibana
	EnrollNode                  SecurityEnrollNode
	GetAPIKey                   SecurityGetAPIKey
	GetBuiltinPrivileges        SecurityGetBuiltinPrivileges
	GetPrivileges               SecurityGetPrivileges
	GetRoleMapping              SecurityGetRoleMapping
	GetRole                     SecurityGetRole
	GetServiceAccounts          SecurityGetServiceAccounts
	GetServiceCredentials       SecurityGetServiceCredentials
	GetSettings                 SecurityGetSettings
	GetStats                    SecurityGetStats
	GetToken                    SecurityGetToken
	GetUserPrivileges           SecurityGetUserPrivileges
	GetUserProfile              SecurityGetUserProfile
	GetUser                     SecurityGetUser
	GrantAPIKey                 SecurityGrantAPIKey
	HasPrivileges               SecurityHasPrivileges
	HasPrivilegesUserProfile    SecurityHasPrivilegesUserProfile
	InvalidateAPIKey            SecurityInvalidateAPIKey
	InvalidateToken             SecurityInvalidateToken
	OidcAuthenticate            SecurityOidcAuthenticate
	OidcLogout                  SecurityOidcLogout
	OidcPrepareAuthentication   SecurityOidcPrepareAuthentication
	PutPrivileges               SecurityPutPrivileges
	PutRoleMapping              SecurityPutRoleMapping
	PutRole                     SecurityPutRole
	PutUser                     SecurityPutUser
	QueryAPIKeys                SecurityQueryAPIKeys
	QueryRole                   SecurityQueryRole
	QueryUser                   SecurityQueryUser
	SamlAuthenticate            SecuritySamlAuthenticate
	SamlCompleteLogout          SecuritySamlCompleteLogout
	SamlInvalidate              SecuritySamlInvalidate
	SamlLogout                  SecuritySamlLogout
	SamlPrepareAuthentication   SecuritySamlPrepareAuthentication
	SamlServiceProviderMetadata SecuritySamlServiceProviderMetadata
	SuggestUserProfiles         SecuritySuggestUserProfiles
	UpdateAPIKey                SecurityUpdateAPIKey
	UpdateCrossClusterAPIKey    SecurityUpdateCrossClusterAPIKey
	UpdateSettings              SecurityUpdateSettings
	UpdateUserProfileData       SecurityUpdateUserProfileData
}

type SQL struct {
	ClearCursor    SQLClearCursor
	DeleteAsync    SQLDeleteAsync
	GetAsync       SQLGetAsync
	GetAsyncStatus SQLGetAsyncStatus
	Query          SQLQuery
	Translate      SQLTranslate
}

type SSL struct {
	Certificates SSLCertificates
}

type Watcher struct {
	AckWatch        WatcherAckWatch
	ActivateWatch   WatcherActivateWatch
	DeactivateWatch WatcherDeactivateWatch
	DeleteWatch     WatcherDeleteWatch
	ExecuteWatch    WatcherExecuteWatch
	GetSettings     WatcherGetSettings
	GetWatch        WatcherGetWatch
	PutWatch        WatcherPutWatch
	QueryWatches    WatcherQueryWatches
	Start           WatcherStart
	Stats           WatcherStats
	Stop            WatcherStop
	UpdateSettings  WatcherUpdateSettings
}

type XPack struct {
	Info  XPackInfo
	Usage XPackUsage
}

func New(t Transport) *API { _ = "STUB: not implemented"; return nil }
