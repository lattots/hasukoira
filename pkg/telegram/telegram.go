package telegram

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SendAudio(bot *bot.Bot, ctx context.Context, chatID int) {
	params, err := getSendVoiceParams(chatID)
	if err != nil {
		log.Printf("error getting send voice params: %v\n", err)
		return
	}

	_, err = bot.SendVoice(ctx, params)
	if err != nil {
		log.Printf("error sending voice: %v\n", err)
		return
	}
}

func getSendVoiceParams(chatID int) (*bot.SendVoiceParams, error) {
	voiceFile, err := os.Open("./data/moi.mp3")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	reader := bufio.NewReader(voiceFile)

	voice := &models.InputFileUpload{
		Filename: "moi",
		Data:     reader,
	}

	params := &bot.SendVoiceParams{
		Voice:  voice,
		ChatID: chatID,
	}
	return params, nil
}
