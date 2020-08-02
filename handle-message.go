package devsearchbot

import (
	"log"
	"net/http"
	"bytes"
	"fmt"

//	"github.com/slack-go/slack"
)

// handleMessage is function for handle the incomming messages
func handleMessage(URL, userName string) {
	var jsonStr = []byte(`{"text":"Hello there, ` + fmt.Sprintf("%s", userName) +`!"}`)
	post, err := http.NewRequest("POST",URL, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	post.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	executePost, er := client.Do(post)
	if er != nil {
		log.Fatal(er)
	}
	defer executePost.Body.Close()
}