package telegram

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"telegram-bot/storage"
)

const (
	ProjectCmd    = "/disableproject"
	HelpCmd       = "/help"
	QuestionCmd   = "/question"
	StartCmd      = "/start"
	ListProjects  = "/projects"
	ListQuestions = "/questions"
)

func (p *Processor) doCmdClient(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)
	log.Printf("got new command '%s' from '%s' %d", text, username, chatID)
	if number, _ := strconv.Atoi(text); number != 0 {
		return p.workWithNumbers(ctx, chatID, number, username)
	}
	switch text {
	case ProjectCmd:
		return p.ProjectUpdate(ctx, chatID, username)
	case HelpCmd:
		return p.SendHelp(ctx, chatID, username)
	case StartCmd:
		p.storage.InsertUser(ctx, username)
		return p.SendHello(ctx, chatID, username)
	case ListProjects:
		return p.SendProjects(ctx, chatID, username)
	case ListQuestions:
		return p.SendQuestions(ctx, chatID, username)
	default:
		return p.SendUnknownCommand(ctx, chatID, username)
	}
}
func (p *Processor) workWithNumbers(ctx context.Context, chatID int, number int, username string) error {
	user := p.storage.UserByUsername(ctx, username)
	if user.ProjectID == 0 {
		return p.setProjectForUser(ctx, chatID, user, number)
	}
	return p.sendQuestionWithAnswer(ctx, chatID, number, username)
}

func (p *Processor) sendQuestionWithAnswer(ctx context.Context, chatID, number int, username string) error {
	question, err := p.storage.QuestionByProjectIDAndOrder(ctx, p.storage.UserByUsername(ctx, username).ProjectID, number)
	if err != nil {
		log.Print(question.ID)
		return p.SendUnknownCommand(ctx, chatID, username)
	}
	text := question.Question + "\n" + question.Answer
	return p.tg.SendMessage(ctx, chatID, text)
}

func (p *Processor) setProjectForUser(ctx context.Context, chatID int, user storage.User, projectID int) error {
	project, err := p.storage.ProjectByID(ctx, projectID)
	if err != nil {
		log.Print(err)
		return p.SendUnknownCommand(ctx, chatID, user.Username)
	}
	p.SendSuccessMessage(ctx, chatID, user.Username)
	return p.storage.SetProjectForUser(ctx, user, project)
}

func (p *Processor) ProjectUpdate(ctx context.Context, chatID int, username string) (err error) {
	p.SendSuccessMessage(ctx, chatID, username)
	return p.storage.UnsetProjectForUser(ctx, p.storage.UserByUsername(ctx, username))
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

func (p *Processor) SendQuestions(ctx context.Context, chatID int, username string) error {
	user := p.storage.UserByUsername(ctx, username)
	questions := p.storage.ListQuestions(ctx, user.ProjectID)
	msg := ""
	for _, quesiton := range questions {
		msg = msg + fmt.Sprintf("%d) %s\n", quesiton.Order, quesiton.Question)
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

func (p *Processor) SendSuccessMessage(ctx context.Context, chatID int, username string) error {
	return p.tg.SendMessage(ctx, chatID, msgSuccess)
}
