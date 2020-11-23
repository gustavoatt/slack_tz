package slack

// HookRequest encapsulates the type of request send by Slack through its WebHook API.
type HookRequest struct {
	Command     string `json:"command"`
	Text        string `json:"text"`
	ResponseURL string `json:"response_url"`
	TriggerID   string `json:"trigger_id"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	// Context about Slack Workspace
	TeamID       string `json:"team_id"`
	EnterpriseID string `json:"enterprise_id"`
	TeamDomain   string `json:"team_domain"`
	APIAppID     string `json:"api_app_id"`
}
