package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"telegram-bot/lib/e"
	"telegram-bot/storage"

	_ "github.com/mattn/go-sqlite3"
)

func (s Storage) UpdateOrInsertUser(_ context.Context, user storage.User) error {
	exists, err := s.isExistUser(user)
	if err != nil {
		e.Wrap("error checking if row exists", err)
	}
	if exists {
		return s.updateUser(user)
	}
	return s.insertUser(user)
}

func (s Storage) isExistUser(user storage.User) (bool, error) {
	var exists bool
	query := fmt.Sprint("SELECT EXISTS (SELECT * FROM USER WHERE USERNAME = ?)")
	err := s.database.QueryRow(query, user.Username).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, err
}

func (s Storage) insertUser(user storage.User) error {
	query, err := s.database.Prepare("INSERT INTO USER(USERNAME) VALUES (?)")
	if err != nil {
		return e.Wrap("something went wrong", err)
	}
	query.Exec(user.Username)
	defer query.Close()
	return nil
}

func (s Storage) updateUser(user storage.User) error {
	query, err := s.database.Prepare("UPDATE USER SET OnChat = ? WHERE USERNAME = ?")
	if err != nil {
		return e.Wrap("something went wrong", err)
	}
	query.Exec(user.OnChat, user.Username)
	defer query.Close()
	if user.Project != (storage.Project{}) {
		query, err := s.database.Prepare("UPDATE USER SET Project = ? WHERE USERNAME = ?")
		if err != nil {
			return e.Wrap("something went wrong", err)
		}
		query.Exec(user.Project.ID)
		defer query.Close()
	}
	return nil
}
