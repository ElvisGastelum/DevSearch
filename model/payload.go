package model

import "encoding/json"

type Payload struct {
	Actions []actions `json:"actions"`
	Token string `json:"token"`
	ResponseURL string `json:"response_url"`
	Channel channel `json:"channel"`
	User user `json:"user"`
	Team team `json:"team"`
	ActionTS string `json:"action_ts"`
	MessageTS string `json:"message_ts"`
	AttachmentID string `json:"attachment_id"`
}

type actions struct {
	Value string `json:"value"`
	Name string `json:"name"`
}

type channel struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type user struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type team struct {
	ID string `json:"id"`
	Domain string `json:"domain"`
}

func (p *Payload) UnmarshallJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &p)
	if err != nil {
		return err
	}
	return nil
}