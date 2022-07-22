package storage

import "context"

type Storage interface {
	UpdateOrInsert(ctx context.Context, username string) error
	From(ctx context.Context, username string) (Project, error)
	ListProjects(ctx context.Context) []Project
	ListQuestion(ctx context.Context, project Project) []Question
	Answer(ctx context.Context, question Question) string
}

type Manager struct {
	username string
}

type User struct {
	username string
	project  Project
}

type Question struct {
	order    int
	question string
	answer   string
	project  Project
}

type Project struct {
	name    string
	manager Manager
}
