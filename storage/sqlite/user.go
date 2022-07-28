package sqlite

import (
	"context"
	"database/sql"
	"telegram-bot/lib/e"
	"telegram-bot/storage"
)

func (s Storage) UserByID(_ context.Context, userID int) storage.User {
	var user storage.User
	s.Database.QueryRow("SELECT ID, USERNAME, PROJECTID, CHATID, ONCHAT, FIRSTNAME, LASTNAME FROM USER WHERE ID = ?", userID).Scan(&user.ID, &user.Username, &user.ProjectID, &user.ChatID, &user.OnChat, &user.FirstName, &user.LastName)
	return user
}

func (s Storage) UserByUsername(_ context.Context, username string) storage.User {
	var user storage.User
	s.Database.QueryRow("SELECT ID, USERNAME, PROJECTID, CHATID, ONCHAT, FIRSTNAME, LASTNAME FROM USER WHERE USERNAME = ?", username).Scan(&user.ID, &user.Username, &user.ProjectID, &user.ChatID, &user.OnChat, &user.FirstName, &user.LastName)
	return user
}

func (s Storage) SetProjectForUser(_ context.Context, user storage.User, project storage.Project) error {
	query, err := s.Database.Prepare("UPDATE USER SET PROJECTID = ? WHERE ID = ?")
	query.Exec(project.ID, user.ID)
	defer query.Close()
	return e.WrapIfErr("error: cannot update user", err)
}

func (s Storage) UnsetProjectForUser(_ context.Context, user storage.User) error {
	query, err := s.Database.Prepare("UPDATE USER SET PROJECTID = NULL WHERE ID = ?")
	query.Exec(user.ID)
	defer query.Close()
	return e.WrapIfErr("error: cannot update user", err)
}

func (s Storage) SetOnChatForUser(_ context.Context, user storage.User) error {
	query, err := s.Database.Prepare("UPDATE USER SET ONCHAT = TRUE WHERE ID = ?")
	query.Exec(user.ID)
	defer query.Close()
	return e.WrapIfErr("error: cannot update user", err)
}

func (s Storage) UnsetOnChatForUser(_ context.Context, user storage.User) error {
	query, err := s.Database.Prepare("UPDATE USER SET ONCHAT = FALSE WHERE ID = ?")
	query.Exec(user.ID)
	defer query.Close()
	return e.WrapIfErr("error: cannot update user", err)
}

func (s Storage) InsertUser(_ context.Context, username string, chatID int) error {
	exist, err := s.isExistUser(username)
	if err != nil || exist {
		return err
	}
	query, err := s.Database.Prepare("INSERT INTO USER(USERNAME, CHATID) VALUES (?, ?)")
	if err != nil {
		return e.Wrap("something went wrong", err)
	}
	query.Exec(username, chatID)
	defer query.Close()
	return nil
}

func (s Storage) isExistUser(username string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS (SELECT * FROM USER WHERE USERNAME = ?)"
	err := s.Database.QueryRow(query, username).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, err
}
