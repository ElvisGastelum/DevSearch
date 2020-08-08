package app

import (
	"log"

	"github.com/elvisgastelum/devsearchbot/helpers"
	"github.com/elvisgastelum/devsearchbot/model"
)

// ButtonActions determines the response a certain button will give
func ButtonActions(action, url string) {

	jsonBytes, err := ActionResponse(action).ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	err = helpers.NewPostRequest(jsonBytes, url)
	if err != nil {
		log.Fatal(err)
	}
}

// ActionResponse return a response to actions
func ActionResponse(action string) *model.ActionResponse {
	actionResponse := model.ActionResponse{
		"replace_original": true,
		"delete_original":  true,
		"text":             nil,
		"response_type":    "ephemeral",
	}

	if action == "cancel" {
		return &actionResponse
	}

	actionResponse["text"] = action
	actionResponse["response_type"] = "in_channel"
	return &actionResponse
}
