package helpers

import (
	"bytes"
	"net/http"
)

// NewPostRequest create a post request from json []bytes
func NewPostRequest(json []byte, url string) error {
	post, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}

	post.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	executePost, err := client.Do(post)
	if err != nil {
		return err
	}
	defer executePost.Body.Close()

	return nil
}
