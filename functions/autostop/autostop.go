package autostop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// RequestBody defines a JSON that will be sent to Discord webhook.
type RequestBody struct {
	Content string `json:"content"`
}

func postDiscord() {
	hook := os.Getenv("DISCORD_HOOK")
	c := os.Getenv("AUTOSTOP_MESSAGE")

	rb := RequestBody{Content: c}
	fmt.Println(rb)

	json, err := json.Marshal(rb)
	fmt.Println(json)
	if err != nil {
		log.Fatalf("Failed to create Discord request: %v", err)
	}

	_, err = http.Post(hook, "application/json", bytes.NewBuffer(json))
	if err != nil {
		log.Fatalf("Failed to post on Discord: %v", err)
	}
}

func stopServer() {
	hook := os.Getenv("STOP_HOOK")

	rb := RequestBody{}
	json, err := json.Marshal(rb)
	if err != nil {
		log.Fatalf("Failed to create Stop request: %v", err)
	}

	_, err = http.Post(hook, "application/json", bytes.NewBuffer(json))
	if err != nil {
		log.Fatalf("Failed to call server-stopping webhook: %v", err)
	}
}

// Autostop stops a server and send a notification to Discord.
func Autostop(w http.ResponseWriter, r *http.Request) {
	postDiscord()
	stopServer()
}
