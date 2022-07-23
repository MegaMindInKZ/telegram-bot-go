package storage

import "context"

type Storage interface {
	UpdateOrInsertUser(ctx context.Context, user User) error
	From(ctx context.Context, username string) (Project, error)
	ListProjects(ctx context.Context) []Project
	ListQuestion(ctx context.Context, project Project) []Question
	Answer(ctx context.Context, question Question) string
	IsManagerBusy(ctx context.Context, manager Manager) (bool, err error)
	SetManagerAndUserBusy(ctx context.Context, manager Manager, user User) error
	UnsetManagerAndUserBusy(ctx context.Context, manager Manager, user User) error
}

type Manager struct {
	ID       int
	Username string
	IsBusy   bool
}

type User struct {
	Username string
	Project  Project
	OnChat   bool
}

type Question struct {
	Order    int
	Question string
	Answer   string
	Project  Project
}

type Project struct {
	ID      int
	Name    string
	Manager Manager
}
