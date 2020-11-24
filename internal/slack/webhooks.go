package slack

// HookRequest encapsulates the type of request send by Slack through its WebHook API for slash commands.
// This is usually received as POST request with form data.
type HookRequest struct {
	Command     string `form:"command"`
	Text        string `form:"text"`
	ResponseURL string `form:"response_url"`
	TriggerID   string `form:"trigger_id"`
	UserID      string `form:"user_id"`
	UserName    string `form:"user_name"`
	// Context about Slack Workspace
	TeamID       string `form:"team_id"`
	EnterpriseID string `form:"enterprise_id"`
	TeamDomain   string `form:"team_domain"`
	APIAppID     string `form:"api_app_id"`
}
