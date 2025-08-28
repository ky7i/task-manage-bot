// Build this Go file.
// go build -o taskManageBot.exe .
//
// My Usage
// 1, Build this to taskManageBot.exe.
// 2, Create .env file to set discord-bot parameters.
// 3, Create below auto-hot-key file.
// ```
// +!d:: Run("taskManageBot.exe", <<yourWorkDirectory>>)
// ```

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
		log.Println(err)
		// Confirm a error message in 10 seconds
		// and return to close cmd App.
		time.Sleep(10 * time.Second)
		return
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println(err)
		time.Sleep(10 * time.Second)
		return
	}

	err = dg.Open()
	if err != nil {
		log.Println(err)
		time.Sleep(10 * time.Second)
		return
	}
	defer dg.Close()

	log.Println("Bot is running...")

	for {
		// In the Pomodoro Technique, work for 25m.
		time.Sleep(25 * time.Minute)
		dg.ChannelMessageSend(channelID, "進捗どうですか？")
		time.Sleep(5 * time.Minute)
	}
	// TODO: summarize and write progresses at Obsidian when given "Ctrl + c"
}
