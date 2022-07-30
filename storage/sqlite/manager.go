package sqlite

import (
	"context"
	"telegram-bot/lib/e"
	"telegram-bot/storage"
)

func (s Storage) ManagerByID(_ context.Context, managerID int) storage.Manager {
	var manager storage.Manager
	err := s.Database.QueryRow("SELECT ID, USERNAME, CHATID, ISBUSY, CURRENTCLIENTID, FIRSTNAME, LASTNAME FROM MANAGER WHERE ID = ?", managerID).Scan(&manager.ID, &manager.Username, &manager.ChatID, &manager.IsBusy, &manager.CurrentClientID, &manager.FirstName, &manager.LastName)
	if err != nil {
		return manager
	}
	return manager
}

func (s Storage) ManagerByUsername(_ context.Context, username string) storage.Manager {
	var manager storage.Manager
	err := s.Database.QueryRow("SELECT ID, USERNAME, CHATID, ISBUSY, CURRENTCLIENTID FROM MANAGER WHERE USERNAME = ?", username).Scan(&manager.ID, &manager.Username, &manager.ChatID, &manager.IsBusy, &manager.CurrentClientID)
	if err != nil {
		return manager
	}
	return manager
}

func (s Storage) SetIsBusyForManager(ctx context.Context, manager storage.Manager, user storage.User) error {
	query, err := s.Database.Prepare("UPDATE MANAGER SET CURRENTCLIENTID = ?, ISBUSY = TRUE WHERE ID = ?;")
	query.Exec(user.ID, manager.ID)
	defer query.Close()
	return e.WrapIfErr("error: cannot update manager", err)
}

func (s Storage) UnsetIsBusyForManager(ctx context.Context, manager storage.Manager) error {
	query, err := s.Database.Prepare("UPDATE MANAGER SET CURRENTCLIENTID = NULL, ISBUSY = FALSE WHERE ID = ?;")
	query.Exec(manager.ID)
	defer query.Close()
	return e.WrapIfErr("error: cannot update manager", err)
}

func (s Storage) IsManager(ctx context.Context, username string) bool {
	var exists bool
	err := s.Database.QueryRow("SELECT EXISTS (SELECT * FROM MANAGER WHERE USERNAME = ?)", username).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
