package telegram

import (
	"context"
	"telegram-bot/lib/e"
	"telegram-bot/storage"
)

const (
	EndChat = "/endchat"
)

func (p *Processor) doCmdManager(ctx context.Context, text string, chatID int, username string) error {
	manager := p.storage.ManagerByUsername(ctx, username)
	if text == EndChat {
		return p.EndChatForManager(ctx, chatID, manager)
	} else {
		return p.SendMessageToClient(ctx, text, chatID, manager)
	}
}

func (p *Processor) EndChatForManager(ctx context.Context, chatID int, manager storage.Manager) error {
	err := p.storage.UnsetIsBusyForManager(ctx, manager)
	if err != nil {
		return e.Wrap("error:cannot set IsBusy for manager", err)
	}
	err = p.storage.UnsetOnChatForUser(ctx, p.storage.UserByID(ctx, manager.CurrentClientID))
	if err != nil {
		return e.Wrap("error: cannot set OnChat for user", err)
	}
	return p.SendSuccessMessage(ctx, chatID, manager.Username)
}

func (p *Processor) SendMessageToClient(ctx context.Context, text string, chatID int, manager storage.Manager) error {
	client := p.storage.UserByID(ctx, manager.CurrentClientID)
	return p.tg.SendMessage(ctx, client.ChatID, text)
}
