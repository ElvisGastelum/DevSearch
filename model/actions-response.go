package model

import (
	"encoding/json"
)

// ActionResponse is the inteface to response from actions
type ActionResponse map[string]interface{}

// ToJSON return the actual value in []byte
func (ar *ActionResponse) ToJSON() ([]byte, error) {
	body, err := json.Marshal(ar)
	if err != nil {
		return nil, err
	}
	return body, nil
}
