package sqlite

import (
	"context"
	"telegram-bot/lib/e"
	"telegram-bot/storage"
)

func (s Storage) ManagerByID(_ context.Context, managerID int) (storage.Manager, error) {
	var manager storage.Manager
	err := s.database.QueryRow("SELECT ID, USERNAME, ISBUSY, CURRENTCLIENTID, FIRSTNAME, LASTNAME WHERE ID = ?", managerID).Scan(&manager.ID, &manager.Username, &manager.IsBusy, &manager.CurrentClientID, &manager.FirstName, &manager.LastName)
	if err != nil {
		return manager, e.Wrap("Manager with this id doesn't exist", err)
	}
	return manager, nil
}

func (s Storage) ManagerByUsername(_ context.Context, username string) (storage.Manager, error) {
	var manager storage.Manager
	err := s.database.QueryRow("SELECT ID, USERNAME, ISBUSY, CURRENTCLIENTID, FIRSTNAME, LASTNAME WHERE USERNAME = ?", username).Scan(&manager.ID, &manager.Username, &manager.IsBusy, &manager.CurrentClientID, &manager.FirstName, &manager.LastName)
	if err != nil {
		return manager, e.Wrap("Manager with this id doesn't exist", err)
	}
	return manager, nil
}

func (s Storage) SetIsBusyForManager(ctx context.Context, manager storage.Manager, user storage.User) error {
	query, err := s.database.Prepare("UPDATE MANAGER SET CURRENTCLIENTID = ?, ISBUSY = TRUE WHERE ID = ?;")
	defer query.Close()
	query.Exec(user.ID, manager.ID)
	return e.WrapIfErr("error: cannot update manager", err)
}

func (s Storage) UnsetIsBusyForManager(ctx context.Context, manager storage.Manager) error {
	query, err := s.database.Prepare("UPDATE MANAGER SET CURRENTCLIENTID = NULL, ISBUSY = FALSE WHERE ID = ?;")
	defer query.Close()
	query.Exec(manager.ID)
	return e.WrapIfErr("error: cannot update manager", err)
}
