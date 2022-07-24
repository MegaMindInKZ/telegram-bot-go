package telegram

import (
	"context"
	"log"
	"strconv"
	"strings"
)

const (
	ProjectCmd  = "/project"
	HelpCmd     = "/help"
	QuestionCmd = "/question"
	StartCmd    = "/start"
)

func doCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)
	log.Printf("got new command '%s' from '%s'", text, username)
	if number, err := strconv.Atoi(text); err != nil {
		number = number + 1
	}
	switch text {
	case ProjectCmd:
		return ProjectUpdate(ctx, chatID, username)
	case HelpCmd:
		return SendHelp(ctx, chatID, username)
	case StartCmd:
		return StartChat(ctx, chatID, username)
	}
	return nil
}

func ProjectUpdate(ctx context.Context, chatID int, username string) (err error) {
	return nil
}

func SendHelp(ctx context.Context, chatID int, username string) (err error) {
	return nil
}

func StartChat(ctx context.Context, chatID int, username string) (err error) {
	return nil
}
