package gitlab

import "time"

type gitLabUserProject struct {
	ID                int           `json:"id"`
	Description       string        `json:"description"`
	Name              string        `json:"name"`
	NameWithNamespace string        `json:"name_with_namespace"`
	Path              string        `json:"path"`
	PathWithNamespace string        `json:"path_with_namespace"`
	CreatedAt         time.Time     `json:"created_at"`
	DefaultBranch     string        `json:"default_branch"`
	TagList           []interface{} `json:"tag_list"`
	Topics            []interface{} `json:"topics"`
	SSHURLToRepo      string        `json:"ssh_url_to_repo"`
	HTTPURLToRepo     string        `json:"http_url_to_repo"`
	WebURL            string        `json:"web_url"`
	ReadmeURL         string        `json:"readme_url"`
	AvatarURL         interface{}   `json:"avatar_url"`
	ForksCount        int           `json:"forks_count"`
	StarCount         int           `json:"star_count"`
	LastActivityAt    time.Time     `json:"last_activity_at"`
	Namespace         struct {
		ID        int         `json:"id"`
		Name      string      `json:"name"`
		Path      string      `json:"path"`
		Kind      string      `json:"kind"`
		FullPath  string      `json:"full_path"`
		ParentID  interface{} `json:"parent_id"`
		AvatarURL string      `json:"avatar_url"`
		WebURL    string      `json:"web_url"`
	} `json:"namespace"`
	ContainerRegistryImagePrefix string `json:"container_registry_image_prefix"`
	Links                        struct {
		Self          string `json:"self"`
		Issues        string `json:"issues"`
		MergeRequests string `json:"merge_requests"`
		RepoBranches  string `json:"repo_branches"`
		Labels        string `json:"labels"`
		Events        string `json:"events"`
		Members       string `json:"members"`
	} `json:"_links"`
	PackagesEnabled bool   `json:"packages_enabled"`
	EmptyRepo       bool   `json:"empty_repo"`
	Archived        bool   `json:"archived"`
	Visibility      string `json:"visibility"`
	Owner           struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"owner"`
	ResolveOutdatedDiffDiscussions bool `json:"resolve_outdated_diff_discussions"`
	ContainerExpirationPolicy      struct {
		Cadence       string      `json:"cadence"`
		Enabled       bool        `json:"enabled"`
		KeepN         int         `json:"keep_n"`
		OlderThan     string      `json:"older_than"`
		NameRegex     string      `json:"name_regex"`
		NameRegexKeep interface{} `json:"name_regex_keep"`
		NextRunAt     time.Time   `json:"next_run_at"`
	} `json:"container_expiration_policy"`
	IssuesEnabled                             bool          `json:"issues_enabled"`
	MergeRequestsEnabled                      bool          `json:"merge_requests_enabled"`
	WikiEnabled                               bool          `json:"wiki_enabled"`
	JobsEnabled                               bool          `json:"jobs_enabled"`
	SnippetsEnabled                           bool          `json:"snippets_enabled"`
	ContainerRegistryEnabled                  bool          `json:"container_registry_enabled"`
	ServiceDeskEnabled                        bool          `json:"service_desk_enabled"`
	ServiceDeskAddress                        string        `json:"service_desk_address"`
	CanCreateMergeRequestIn                   bool          `json:"can_create_merge_request_in"`
	IssuesAccessLevel                         string        `json:"issues_access_level"`
	RepositoryAccessLevel                     string        `json:"repository_access_level"`
	MergeRequestsAccessLevel                  string        `json:"merge_requests_access_level"`
	ForkingAccessLevel                        string        `json:"forking_access_level"`
	WikiAccessLevel                           string        `json:"wiki_access_level"`
	BuildsAccessLevel                         string        `json:"builds_access_level"`
	SnippetsAccessLevel                       string        `json:"snippets_access_level"`
	PagesAccessLevel                          string        `json:"pages_access_level"`
	OperationsAccessLevel                     string        `json:"operations_access_level"`
	AnalyticsAccessLevel                      string        `json:"analytics_access_level"`
	ContainerRegistryAccessLevel              string        `json:"container_registry_access_level"`
	EmailsDisabled                            interface{}   `json:"emails_disabled"`
	SharedRunnersEnabled                      bool          `json:"shared_runners_enabled"`
	LfsEnabled                                bool          `json:"lfs_enabled"`
	CreatorID                                 int           `json:"creator_id"`
	ImportStatus                              string        `json:"import_status"`
	OpenIssuesCount                           int           `json:"open_issues_count"`
	CiDefaultGitDepth                         int           `json:"ci_default_git_depth"`
	CiForwardDeploymentEnabled                bool          `json:"ci_forward_deployment_enabled"`
	CiJobTokenScopeEnabled                    bool          `json:"ci_job_token_scope_enabled"`
	PublicJobs                                bool          `json:"public_jobs"`
	BuildTimeout                              int           `json:"build_timeout"`
	AutoCancelPendingPipelines                string        `json:"auto_cancel_pending_pipelines"`
	BuildCoverageRegex                        interface{}   `json:"build_coverage_regex"`
	CiConfigPath                              string        `json:"ci_config_path"`
	SharedWithGroups                          []interface{} `json:"shared_with_groups"`
	OnlyAllowMergeIfPipelineSucceeds          bool          `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               interface{}   `json:"allow_merge_on_skipped_pipeline"`
	RestrictUserDefinedVariables              bool          `json:"restrict_user_defined_variables"`
	RequestAccessEnabled                      bool          `json:"request_access_enabled"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool          `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool          `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestLinkEnabled           bool          `json:"printing_merge_request_link_enabled"`
	MergeMethod                               string        `json:"merge_method"`
	SquashOption                              string        `json:"squash_option"`
	SuggestionCommitMessage                   interface{}   `json:"suggestion_commit_message"`
	AutoDevopsEnabled                         bool          `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string        `json:"auto_devops_deploy_strategy"`
	AutocloseReferencedIssues                 bool          `json:"autoclose_referenced_issues"`
	KeepLatestArtifact                        bool          `json:"keep_latest_artifact"`
	ExternalAuthorizationClassificationLabel  string        `json:"external_authorization_classification_label"`
	RequirementsEnabled                       bool          `json:"requirements_enabled"`
	SecurityAndComplianceEnabled              bool          `json:"security_and_compliance_enabled"`
	ComplianceFrameworks                      []interface{} `json:"compliance_frameworks"`
	Permissions                               struct {
		ProjectAccess struct {
			AccessLevel       int `json:"access_level"`
			NotificationLevel int `json:"notification_level"`
		} `json:"project_access"`
		GroupAccess interface{} `json:"group_access"`
	} `json:"permissions"`
}
