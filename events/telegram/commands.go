package telegram

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	ProjectCmd   = "/project"
	HelpCmd      = "/help"
	QuestionCmd  = "/question"
	StartCmd     = "/start"
	ListProjects = "/projects"
)

func (p *Processor) doCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)
	log.Printf("got new command '%s' from '%s'", text, username)
	if number, err := strconv.Atoi(text); err != nil {
		number = number + 1
		fmt.Print(number)
	}
	switch text {
	case ProjectCmd:
		return p.ProjectUpdate(ctx, chatID, username)
	case HelpCmd:
		return p.SendHelp(ctx, chatID, username)
	case StartCmd:
		return p.SendHello(ctx, chatID, username)
	case ListProjects:
		return p.SendProjects(ctx, chatID, username)
	default:
		return p.SendUnknownCommand(ctx, chatID, username)
	}
}

func (p *Processor) ProjectUpdate(ctx context.Context, chatID int, username string) (err error) {
	return nil
}

func (p *Processor) SendProjects(ctx context.Context, chatID int, username string) error {
	listProjects, err := p.storage.ListProjects(ctx)
	if err != nil {
		return p.tg.SendMessage(ctx, chatID, err.Error())
	}
	msg := ""
	for _, project := range listProjects {
		msg = msg + fmt.Sprintf("%d) %s", project.ID, project.Name)
	}
	return p.tg.SendMessage(ctx, chatID, msg)

}

func (p *Processor) SendHelp(ctx context.Context, chatID int, username string) error {
	return p.tg.SendMessage(ctx, chatID, msgHelp)
}

func (p *Processor) SendHello(ctx context.Context, chatID int, username string) (err error) {
	return p.tg.SendMessage(ctx, chatID, msgHello)
}

func (p *Processor) SendUnknownCommand(ctx context.Context, chatID int, username string) (err error) {
	return p.tg.SendMessage(ctx, chatID, msgUnknownCommand)
}

// func (p *Processor) WorkWithNumber(ctx context.Context, chatID, number int, username string) (err error) {
// 	defer func() { err = e.WrapIfErr("Can't work with number", err) }()
// }
