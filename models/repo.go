// Generated by jot version 0.0.0 on 2016-11-23 04:32:04.306334247 +0000 UTC

package models 

type Repo struct {
	Parent           *Repo                  `json:"parent"`
	ArchiveUrl       string                 `json:"archive_url"`
	ContentsUrl      string                 `json:"contents_url"`
	TreesUrl         string                 `json:"trees_url"`
	CreatedAt        string                 `json:"created_at"`
	SvnUrl           string                 `json:"svn_url"`
	HasIssues        bool                   `json:"has_issues"`
	HasWiki          bool                   `json:"has_wiki"`
	Permissions      map[string]interface{} `json:"permissions"`
	GitUrl           string                 `json:"git_url"`
	KeysUrl          string                 `json:"keys_url"`
	CollaboratorsUrl string                 `json:"collaborators_url"`
	GitTagsUrl       string                 `json:"git_tags_url"`
	MirrorUrl        string                 `json:"mirror_url"`
	SubscriptionUrl  string                 `json:"subscription_url"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	ContributorsUrl  string                 `json:"contributors_url"`
	DeploymentsUrl   string                 `json:"deployments_url"`
	ForksUrl         string                 `json:"forks_url"`
	IssueCommentUrl  string                 `json:"issue_comment_url"`
	IssueEventsUrl   string                 `json:"issue_events_url"`
	OpenIssuesCount  float64                `json:"open_issues_count"`
	Fork             bool                   `json:"fork"`
	BranchesUrl      string                 `json:"branches_url"`
	TagsUrl          string                 `json:"tags_url"`
	SshUrl           string                 `json:"ssh_url"`
	SubscribersUrl   string                 `json:"subscribers_url"`
	Source           *Repo                  `json:"source"`
	HtmlUrl          string                 `json:"html_url"`
	GitCommitsUrl    string                 `json:"git_commits_url"`
	StatusesUrl      string                 `json:"statuses_url"`
	CommentsUrl      string                 `json:"comments_url"`
	MilestonesUrl    string                 `json:"milestones_url"`
	UpdatedAt        string                 `json:"updated_at"`
	EventsUrl        string                 `json:"events_url"`
	Homepage         string                 `json:"homepage"`
	StargazersUrl    string                 `json:"stargazers_url"`
	ForksCount       float64                `json:"forks_count"`
	StargazersCount  float64                `json:"stargazers_count"`
	HasDownloads     bool                   `json:"has_downloads"`
	PushedAt         string                 `json:"pushed_at"`
	CommitsUrl       string                 `json:"commits_url"`
	GitRefsUrl       string                 `json:"git_refs_url"`
	ReleasesUrl      string                 `json:"releases_url"`
	WatchersCount    float64                `json:"watchers_count"`
	CompareUrl       string                 `json:"compare_url"`
	MergesUrl        string                 `json:"merges_url"`
	PullsUrl         string                 `json:"pulls_url"`
	HasPages         bool                   `json:"has_pages"`
	Organization     map[string]interface{} `json:"organization"`
	Owner            *User                  `json:"owner"`
	CloneUrl         string                 `json:"clone_url"`
	HooksUrl         string                 `json:"hooks_url"`
	AssigneesUrl     string                 `json:"assignees_url"`
	DownloadsUrl     string                 `json:"downloads_url"`
	NotificationsUrl string                 `json:"notifications_url"`
	DefaultBranch    string                 `json:"default_branch"`
	SubscribersCount float64                `json:"subscribers_count"`
	Id               float64                `json:"id"`
	FullName         string                 `json:"full_name"`
	TeamsUrl         string                 `json:"teams_url"`
	Private          bool                   `json:"private"`
	LanguagesUrl     string                 `json:"languages_url"`
	IssuesUrl        string                 `json:"issues_url"`
	LabelsUrl        string                 `json:"labels_url"`
	Size             float64                `json:"size"`
	Url              string                 `json:"url"`
	BlobsUrl         string                 `json:"blobs_url"`
}