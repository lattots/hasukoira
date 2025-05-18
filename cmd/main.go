package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/lattots/hasukoira/pkg/telegram"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./data/secrets.env")
	if err != nil {
		log.Fatalln("error loading environment variables: ", err)
	}

	fmt.Println("Creating bot...")
	b, err := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"), bot.WithDefaultHandler(defaultHandler))
	if err != nil {
		log.Fatalln("error creating bot:\n", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b.Start(ctx)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Printf("Moi t. %s\n", update.Message.From.Username)
	telegram.SendAudio(b, ctx, int(update.Message.Chat.ID))
}
