package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

// var (
// 	TelegramToken = os.Getenv("TelegramToken")
// )

var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "Telegram bot for conversation",
	Long:    `A simple telegram bot for conversation with a user.`,

	Run: func(cmd *cobra.Command, args []string) {
		TelegramToken := os.Getenv("TelegramToken")
		if TelegramToken == "" {
			log.Fatal("TelegramToken environment variable is not set")
		}

		fmt.Printf("kbot %s started\n", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TelegramToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Failed to create bot: %v", err)
		}

		kbot.Handle("/start", func(c telebot.Context) error {
			// return c.Send("Hello! I'm a simple bot. How can I help you?")
			user := c.Sender()
			text := c.Text()
			log.Printf("User %s (%d) wrote: %s", user.Username, user.ID, text)
			return c.Send("Hello! I'm a simple bot. How can I help you?" + text)
		})

		kbot.Handle("/quote", func(c telebot.Context) error {
			quotes := []string{
				"Не помиляється той, хто нічого не робить.",
				"Великі справи починаються з маленьких кроків.",
				"Успіх — це рух від невдачі до невдачі без втрати ентузіазму.",
				"Життя — це 10% того, що з тобою відбувається, і 90% того, як ти на це реагуєш.",
			}
			idn := time.Now().Unix() % int64(len(quotes))
			return c.Send(quotes[idn])
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}
