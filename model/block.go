package model

import "encoding/json"

type SlashCommandResponse struct {
	Blocks [4]Block `json:"blocks,omitempty"`
}

type Block struct {
	Type      string    `json:"type,omitempty"`
	Text      TextInfo  `json:"text,omitempty"`
	Accessory Accessory `json:"accessory,omitempty"`
	BlockID   string    `json:"block_id,omitempty"`
	Elements  []Button  `josn:"elements"`
}

type TextInfo struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Accessory struct {
	Type  string        `json:"type,omitempty"`
	Text  AccessoryText `json:"text,omitempty"`
	Value string        `json:"value,omitempty"`
}

type AccessoryText struct {
	Type  string `json:"type,omitempty"`
	Text  string `json:"text,omitempty"`
	Emoji bool   `json:"emoji,omitempty"`
}

type Button struct {
	Type  string        `json:"type,omitempty"`
	Text  AccessoryText `json:"text,omitempty"`
	Style string        `json:"style,omitempty"`
	Value string        `json:"value,omitempty"`
}

// ToJSON return the actual value in []byte
func (scr *SlashCommandResponse) ToJSON() ([]byte, error) {
	body, err := json.Marshal(scr)
	if err != nil {
		return nil, err
	}
	return body, nil
}
