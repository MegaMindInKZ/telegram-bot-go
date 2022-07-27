package sqlite

import (
	"context"
	"telegram-bot/lib/e"
	"telegram-bot/storage"
)

func (s Storage) UserByID(_ context.Context, userID int) storage.User {
	var user storage.User
	s.Database.QueryRow("SELECT ID, USERNAME, PROJECTID, ONCHAT, FIRSTNAME, LASTNAME WHERE ID = ?", userID).Scan(&user.ID, &user.Username, &user.ProjectID, &user.OnChat, &user.FirstName, &user.LastName)
	return user
}

func (s Storage) UserByUsername(_ context.Context, username string) storage.User {
	var user storage.User
	s.Database.QueryRow("SELECT ID, USERNAME, PROJECTID, ONCHAT, FIRSTNAME, LASTNAME WHERE USERNAME = ?", username).Scan(&user.ID, &user.Username, &user.ProjectID, &user.OnChat, &user.FirstName, &user.LastName)
	return user
}

func (s Storage) SetProjectForUser(_ context.Context, user storage.User, project storage.Project) error {
	query, err := s.Database.Prepare("UPDATE USER SET PROJECTID = ? WHERE ID = ?")
	defer query.Close()
	query.Exec(project.ID, user.ID)
	return e.WrapIfErr("error: cannot update user", err)
}

func (s Storage) UnsetProjectForUser(_ context.Context, user storage.User) error {
	query, err := s.Database.Prepare("UPDATE USER SET PROJECTID = NULL WHERE ID = ?")
	defer query.Close()
	query.Exec(user.ID)
	return e.WrapIfErr("error: cannot update user", err)
}

func (s Storage) SetOnChatForUser(_ context.Context, user storage.User) error {
	query, err := s.Database.Prepare("UPDATE USER SET ONCHAT = TRUE WHERE ID = ?")
	defer query.Close()
	query.Exec(user.ID)
	return e.WrapIfErr("error: cannot update user", err)
}

func (s Storage) UnsetOnChatForUser(_ context.Context, user storage.User) error {
	query, err := s.Database.Prepare("UPDATE USER SET ONCHAT = FALSE WHERE ID = ?")
	defer query.Close()
	query.Exec(user.ID)
	return e.WrapIfErr("error: cannot update user", err)
}
