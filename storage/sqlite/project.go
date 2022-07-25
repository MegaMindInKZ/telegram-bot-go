package sqlite

import (
	"context"
	"telegram-bot/lib/e"
	"telegram-bot/storage"
)

func (s Storage) Project(_ context.Context, projectID int) (storage.Project, error) {
	var project storage.Project
	err := s.database.QueryRow("SELECT ID, NAME, MANAGER WHERE ID = ?", projectID).Scan(&project.ID, &project.Name, &project.ManagerID)
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
