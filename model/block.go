package model

import "encoding/json"

type SlashCommandResponse map[string]interface{}

// ToJSON return the actual value in []byte
func (scr *SlashCommandResponse) ToJSON() ([]byte, error) {
	body, err := json.Marshal(scr)
	if err != nil {
		return nil, err
	}
	return body, nil
}
