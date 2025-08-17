package main

import (
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	err = dg.Open()
	if err != nil {
		panic(err)
	}
	defer dg.Close()

	log.Println("Bot is running...")

	for {
		// In the Pomodoro Technique, work for 25m.
		time.Sleep(25 * time.Minute)
		dg.ChannelMessageSend(channelID, "進捗どうですか？")
		time.Sleep(5 * time.Minute)
	}
}
