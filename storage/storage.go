package storage

import "context"

type Storage interface {
	UserByID(ctx context.Context, userID int) User            //done
	UserByUsername(ctx context.Context, username string) User //done
	InsertUser(ctx context.Context, username string) error    //done
	ProjectByID(ctx context.Context, projectID int) (Project, error)
	ManagerByID(ctx context.Context, managerID int) (Manager, error)                             //done
	ManagerByUsername(ctx context.Context, username string) (Manager, error)                     //done
	SetIsBusyForManager(ctx context.Context, manager Manager, user User) error                   //done
	UnsetIsBusyForManager(ctx context.Context, manager Manager) error                            //done
	IsManager(ctx context.Context, username string) bool                                         //done
	QuestionByProjectIDAndOrder(ctx context.Context, projectID int, order int) (Question, error) //done
	SetProjectForUser(ctx context.Context, user User, project Project) error                     //done
	UnsetProjectForUser(ctx context.Context, user User) error                                    //done
	SetOnChatForUser(ctx context.Context, user User) error                                       //done
	UnsetOnChatForUser(ctx context.Context, user User) error                                     //done
	ListQuestions(ctx context.Context, projectID int) []Question
	ListProjects(ctx context.Context) ([]Project, error) //done

}

type Manager struct {
	ID              int
	Username        string
	ChatID          int
	IsBusy          bool
	CurrentClientID int
	FirstName       string
	LastName        string
}

type User struct {
	ID        int
	Username  string
	ProjectID int
	OnChat    bool
	FirstName string
	LastName  string
}

type Question struct {
	ID        int
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
