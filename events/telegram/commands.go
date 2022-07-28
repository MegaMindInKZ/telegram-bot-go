package telegram

import (
	"context"
)

func (p *Processor) doCmd(ctx context.Context, text string, chatID int, username string) error {
	if p.storage.IsManager(ctx, username) {
		return p.doCmdManager(ctx, text, chatID, username)
	}
	return p.doCmdClient(ctx, text, chatID, username)
}

// func (p *Processor) WorkWithNumber(ctx context.Context, chatID, number int, username string) (err error) {
// 	defer func() { err = e.WrapIfErr("Can't work with number", err) }()
// }
