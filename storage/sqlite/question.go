package sqlite

import (
	"context"
	"telegram-bot/storage"
)

func (s Storage) QuestionByProjectIDAndOrder(_ context.Context, projectID int, order int) (storage.Question, error) {
	var question storage.Question
	err := s.database.QueryRow("SELECT ID, ORDER, QUESTION, ANSWER, PROJECTID WHERE ORDER = ? AND PROJECTID = ?", order, projectID).Scan(&question.ID, &question.Order, &question.Question, &question.Answer, &question.ProjectID)
	if err != nil {
		return storage.Question{}, err
	}
	return question, nil
}

func (s Storage) ListQuestions(_ context.Context, projectID int) ([]storage.Question, error) {
	var questions []storage.Question
	rows, err := s.database.Query("SELECT id, order, question, answer, projectid where projectID = ?", projectID)
	if err != nil {
		return questions, err
	}
	for rows.Next() {
		var question storage.Question
		err := rows.Scan(question.ID, &question.Order, &question.Question, &question.Answer, &question.ProjectID)
		if err != nil {
			return questions, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}
