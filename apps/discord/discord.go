package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type messageTemplate struct {
	start string
	stop  string
}

func readMessageTemplate() *messageTemplate {
	start := os.Getenv("START_MESSAGE")
	stop := os.Getenv("STOP_MESSAGE")

	return &messageTemplate{start, stop}
}

type requestBody struct {
}

func newDiscordGo() (*discordgo.Session, error) {
	token := os.Getenv("DISCORD_TOKEN")
	str := fmt.Sprintf("Bot %v", token)

	return discordgo.New(str)
}

func startServer() error {
	hook := os.Getenv("START_HOOK")
	json, err := json.Marshal(requestBody{})

	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	_, err = http.Post(hook, "application/json", bytes.NewBuffer(json))
	return err
}

func stopServer() error {
	hook := os.Getenv("STOP_HOOK")
	json, err := json.Marshal(requestBody{})

	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	_, err = http.Post(hook, "application/json", bytes.NewBuffer(json))
	return err
}

func monitorMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	mt := readMessageTemplate()

	if strings.Contains(m.Content, ":minecraft_start:") {
		err := startServer()
		if err != nil {
			log.Fatalf("Starting server failed: %v", err)
		}
		s.ChannelMessageSend(m.ChannelID, mt.start)
	}
	if strings.Contains(m.Content, ":minecraft_stop:") {
		err := stopServer()
		if err != nil {
			log.Fatalf("Stopping server failed: %v", err)
		}
		s.ChannelMessageSend(m.ChannelID, mt.stop)
	}
}

func main() {
	dg, err := newDiscordGo()
	if err != nil {
		log.Fatalf("Failed to create Discord session: %v", err)
	}

	dg.AddHandler(monitorMessage)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		log.Fatalf("Failed to opening Discord connection: %v", err)
	}

	log.Println("Bot is now running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
