package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type DiscordConfig struct {
	Token   string
	Channel string
}

type requestBody struct {
}

func newDiscordGo() (*discordgo.Session, error) {
	token := os.Getenv("DISCORD_TOKEN")
	str := fmt.Sprintf("Bot %v", token)

	fmt.Println(str)

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
	if m.Content == "mc start" {
		err := startServer()
		if err != nil {
			log.Fatalf("Starting server failed: %v", err)
		}
		s.ChannelMessageSend(m.ChannelID, "たぶんサーバーが起動したよ！")
	}
	if m.Content == "mc stop" {
		err := stopServer()
		if err != nil {
			log.Fatalf("Stopping server failed: %v", err)
		}
		s.ChannelMessageSend(m.ChannelID, "たぶんサーバーが停止したよ！")
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
