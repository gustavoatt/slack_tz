package slack

type ActionBlock struct {
	Type    string `json:"type"`
	BlockId string `json:"block_id"`
}
