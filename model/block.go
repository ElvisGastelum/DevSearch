package model

type Block struct {
	Type    string   `json:"type"`
	Text    textInfo `json:"text"`
	BlockID string   `json:"block_id"`
}

type textInfo struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

