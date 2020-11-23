package slack

// HookRequest encapsulates the type of request send by Slack through its WebHook API.
type HookRequest struct {
	Command     string `json:"command"`
	Text        string `json:"text"`
	ResponseURL string `json:"response_url"`
	TriggerId   string `json:"trigger_id"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	// Context about Slack Workspace
	TeamId       string `json:"team_id"`
	EnterpriseId string `json:"enterprise_id"`
	TeamDomain   string `json:"team_domain"`
	ApiAppId     string `json:"api_app_id"`
}
