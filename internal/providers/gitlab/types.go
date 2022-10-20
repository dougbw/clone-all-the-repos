package gitlab

import "time"

type GitLabUser []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type GitLabGroup []struct {
	ID                             int         `json:"id"`
	WebURL                         string      `json:"web_url"`
	Name                           string      `json:"name"`
	Path                           string      `json:"path"`
	Description                    string      `json:"description"`
	Visibility                     string      `json:"visibility"`
	ShareWithGroupLock             bool        `json:"share_with_group_lock"`
	RequireTwoFactorAuthentication bool        `json:"require_two_factor_authentication"`
	TwoFactorGracePeriod           int         `json:"two_factor_grace_period"`
	ProjectCreationLevel           string      `json:"project_creation_level"`
	AutoDevopsEnabled              interface{} `json:"auto_devops_enabled"`
	SubgroupCreationLevel          string      `json:"subgroup_creation_level"`
	EmailsDisabled                 interface{} `json:"emails_disabled"`
	MentionsDisabled               interface{} `json:"mentions_disabled"`
	LfsEnabled                     bool        `json:"lfs_enabled"`
	DefaultBranchProtection        int         `json:"default_branch_protection"`
	AvatarURL                      interface{} `json:"avatar_url"`
	RequestAccessEnabled           bool        `json:"request_access_enabled"`
	FullName                       string      `json:"full_name"`
	FullPath                       string      `json:"full_path"`
	CreatedAt                      time.Time   `json:"created_at"`
	ParentID                       interface{} `json:"parent_id"`
	LdapCn                         interface{} `json:"ldap_cn"`
	LdapAccess                     interface{} `json:"ldap_access"`
}
