package slack

type ActionBlock struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id"`
}
