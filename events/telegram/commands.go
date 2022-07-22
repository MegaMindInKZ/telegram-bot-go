package telegram

import (
	"context"
	"log"
	"strings"
)

const (
	HelpCmd     = "/help"
	QuestionCmd = "/question"
	StartCmd    = "/start"
)

func doCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)
	log.Printf
}
