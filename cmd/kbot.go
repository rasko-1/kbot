package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

func LogUserMessage(c telebot.Context) {
	user := c.Sender()
	text := c.Text()
	log.Printf("User %s (%d) wrote: %s", user.Username, user.ID, text)
}

var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "Telegram bot",
	Long:    `A simple telegram bot`,

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
			LogUserMessage(c)
			return c.Send("Hello! I'm a simple bot.\n\n1. /quote - get a random quote\n2. /weather <city> - get the weather for a city")
		})

		kbot.Handle("/quote", func(c telebot.Context) error {
			LogUserMessage(c)
			url := ("https://dummyjson.com/quotes/random")
			resp, err := http.Get(url)
			if err != nil {
				return c.Send("Could not fetch quote: " + err.Error())
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				return c.Send("Failed to get quote")
			}
			var data struct {
				Quote  string `json:"quote"`
				Author string `json:"author"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				return c.Send("Could not decode quote data: " + err.Error())
			}
			msg := fmt.Sprintf("Quote: %s\n\nAuthor: %s", data.Quote, data.Author)
			return c.Send(msg)
		})

		kbot.Handle("/weather", func(c telebot.Context) error {
			LogUserMessage(c)
			args := c.Args()
			if len(args) == 0 {
				return c.Send("Please provide a city name. Example: /weather Kyiv")
			}
			city := strings.Join(args, " ")
			apiKey := os.Getenv("WeatherApiKey")
			if apiKey == "" {
				return c.Send("WeatherApiKey environment variable is not set")
			}
			url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=yes", apiKey, city)
			resp, err := http.Get(url)
			if err != nil {
				return c.Send("Could not fetch weather data: " + err.Error())
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				return c.Send("Failed to get weather for this city:" + city)
			}

			var data struct {
				Location struct {
					Name      string `json:"name"`
					Country   string `json:"country"`
					Localtime string `json:"localtime"`
				} `json:"location"`
				Current struct {
					TempC     float64 `json:"temp_c"`
					Condition struct {
						Text string `json:"text"`
						Icon string `json:"icon"`
					} `json:"condition"`
				} `json:"current"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				return c.Send("Could not decode weather data: " + err.Error())
			}

			msg := fmt.Sprintf("Weather in %s, %s\nTemperature: %.1f°C - %s\nTime: %s\n%s\n", data.Location.Name, data.Location.Country, data.Current.TempC, data.Current.Condition.Text, data.Location.Localtime, data.Current.Condition.Icon)
			return c.Send(msg)
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}
