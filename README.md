# kbot

## Overview
`kbot` is a Telegram bot written in Go that provides useful commands such as getting a random quote and checking the weather in a specified city. The project uses [Cobra](https://github.com/spf13/cobra) for CLI and [telebot](https://github.com/tucnak/telebot) for working with the Telegram API.

## Features
- Responds to the `/start` command
- Get a random quote with `/quote`
- Check the weather in a city with `/weather <city>`

## Prerequisites
- Go 1.18 or newer
- Telegram Bot Token (get it from [BotFather](https://t.me/botfather))
- API key for [weatherapi.com](https://www.weatherapi.com/)
- (Recommended) `.env` file for storing tokens

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/kbot.git
    cd kbot
    ```
2. Install dependencies:
    ```bash
    go mod tidy
    ```

## Configuration
Create a `.env` file in the project root and add your tokens:
```
TELE_TOKEN=your_telegram_token
WeatherApiKey=your_weatherapi_key
```

## Usage
Run the bot with:
```bash
go run main.go kbot
```
or build and run the binary:
```bash
go build -o kbot
./kbot kbot
```

## Bot Link
You can find the bot here: [kbot](https://t.me/josefSX6bot)