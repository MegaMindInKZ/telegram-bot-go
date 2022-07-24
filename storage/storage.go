package storage

import "context"

type Storage interface {
	UpdateOrInsertUser(ctx context.Context, user User) error
	Project(ctx context.Context, projectID int) (Project, error)
	ListProjects(ctx context.Context) ([]Project, error)
	ListQuestion(ctx context.Context, project Project) ([]Question, error)
	Question(ctx context.Context, user User, order int) (Question, error)
	Manager(ctx context.Context, managerID int) (Manager, error)
	SetManagerAndUserBusy(ctx context.Context, managerID int, userID int) error
	UnsetManagerAndUserBusy(ctx context.Context, managerID int, userID int) error
}

type Manager struct {
	ID       int
	Username string
	IsBusy   bool
}

type User struct {
	Username  string
	ProjectID int
	OnChat    bool
}

type Question struct {
	Order     int
	Question  string
	Answer    string
	ProjectID int
}

type Project struct {
	ID        int
	Name      string
	ManagerID int
}
