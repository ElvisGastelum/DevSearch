package slackbot_test

import (
	"testing"
	"github.com/elvisgastelum/dev-search/slackbot"
)

func TestHelloWorld(t *testing.T){
	name := "Elvis"
	expected := "Hello, " + name + "!"
	result := slackbot.HelloName("Elvis")

	if result != expected {
		t.Fatalf("Expected, %s, but you got %s", expected, result)
	}
}
