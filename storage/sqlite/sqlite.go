package sqlite

import (
	"context"
	"database/sql"
	"log"
	"telegram-bot/lib/e"
	"telegram-bot/storage"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./telegram-bot.db"

type Storage struct {
	database *sql.DB
}

func New() *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (s Storage) Project(_ context.Context, projectID int) (storage.Project, error) {
	var project storage.Project
	err := s.database.QueryRow("SELECT ID, NAME, MANAGER WHERE ID = ?", projectID).Scan(&project.ID, &project.ManagerID)
	if err != nil {
		return storage.Project{}, err
	}
	return project, nil
}

func (s Storage) ListProjects(ctx context.Context) ([]storage.Project, error) {
	var projects []storage.Project
	rows, err := s.database.Query("SELECT ID, NAME, MANAGERID FROM PROJECT")
	if err != nil {
		return projects, err
	}
	for rows.Next() {
		var project storage.Project
		err := rows.Scan(&project.ID, &project.Name, &project.ManagerID)
		if err != nil {
			return projects, e.Wrap("error: scanning error of row", err)
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (s Storage) ListQuestion(_ context.Context, project storage.Project) ([]storage.Question, error) {
	var questions []storage.Question
	rows, err := s.database.Query("SELECT order, question, answer, projectid where projectID = ?", project.ID)
	if err != nil {
		return questions, err
	}
	for rows.Next() {
		var question storage.Question
		err := rows.Scan(&question.Order, &question.Question, &question.Answer, &question.ProjectID)
		if err != nil {
			return questions, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (s Storage) Question(_ context.Context, user storage.User, order int) (storage.Question, error) {
	var question storage.Question
	err := s.database.QueryRow("SELECT ORDER, QUESTION, ANSWER, PROJECTID WHERE ORDER = ? AND PROJECTID = ?", order, user.ProjectID).Scan(&question.Order, &question.Question, &question.Answer, &question.ProjectID)
	if err != nil {
		return storage.Question{}, err
	}
	return question, nil
}

func (s Storage) Manager(_ context.Context, managerID int) (storage.Manager, error) {
	var manager storage.Manager
	err := s.database.QueryRow("SELECT ID, USERNAME, ISBUSY FROM MANAGER WHERE ID = ?", managerID).Scan(&manager.ID, manager.Username, manager.IsBusy)
	if err != nil {
		return storage.Manager{}, err
	}
	return manager, nil
}

func (s Storage) SetManagerAndUserBusy(_ context.Context, managerID int, userID int) error {
	query1, err := s.database.Prepare("UPDATE MANAGER SET ISBUSY = TRUE WHERE ID = ?")
	if err != nil {
		return err
	}
	defer query1.Close()
	query1.Exec(managerID)
	query2, err := s.database.Prepare("UPDATE USER SET ONCHAT = TRUE WHERE ID = ?")
	if err != nil {
		return err
	}
	query2.Exec(userID)
	defer query2.Close()
	return nil
}

func (s Storage) UnsetManagerAndUserBusy(_ context.Context, managerID int, userID int) error {
	query1, err := s.database.Prepare("UPDATE MANAGER SET ISBUSY = FALSE WHERE ID = ?")
	if err != nil {
		return err
	}
	defer query1.Close()
	query1.Exec(managerID)
	query2, err := s.database.Prepare("UPDATE USER SET ONCHAT = FALSE WHERE ID = ?")
	if err != nil {
		return err
	}
	query2.Exec(userID)
	defer query2.Close()
	return nil
}
