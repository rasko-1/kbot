# kbot

## Overview
`kbot` — це Telegram-бот, написаний на Go, який надає корисні команди, такі як отримання випадкової цитати та перегляд погоди у вказаному місті. Проєкт використовує бібліотеки [Cobra](https://github.com/spf13/cobra) для CLI та [telebot](https://github.com/tucnak/telebot) для роботи з Telegram API.

## Features
- Відповідь на команду `/start`
- Отримання випадкової цитати через `/quote`
- Перегляд погоди у місті через `/weather <місто>`

## Prerequisites
- Go 1.18 або новіше
- Telegram Bot Token (отримати у [BotFather](https://t.me/botfather))
- API ключ для [weatherapi.com](https://www.weatherapi.com/)
- (Рекомендовано) файл `.env` для зберігання токенів

## Installation
1. Клонуйте репозиторій:
    ```bash
    git clone https://github.com/yourusername/kbot.git
    cd kbot
    ```
2. Встановіть залежності:
    ```bash
    go mod tidy
    ```

## Configuration
Створіть файл `.env` у корені проєкту та додайте ваші токени:
```
TELE_TOKEN=your_telegram_token
WeatherApiKey=your_weatherapi_key
```

## Usage
Запустіть бота командою:
```bash
go run main.go kbot
```
або зібраний бінарник:
```bash
go build -o kbot
./kbot kbot
```

## Bot Link
Ви можете знайти бота за посиланням: [kbot](https://t.me/josefSX6bot)