package sqlite

import (
	"context"
	"log"
	"telegram-bot/storage"
)

func (s Storage) QuestionByProjectIDAndOrder(_ context.Context, projectID int, order int) (storage.Question, error) {
	var question storage.Question
	err := s.Database.QueryRow("SELECT ID, [ORDER], QUESTION, ANSWER, PROJECTID FROM QUESTION WHERE [ORDER] = ? AND PROJECTID = ?", order, projectID).Scan(&question.ID, &question.Order, &question.Question, &question.Answer, &question.ProjectID)
	if err != nil {
		return storage.Question{}, err
	}
	return question, nil
}

func (s Storage) ListQuestions(_ context.Context, projectID int) []storage.Question {
	var questions []storage.Question
	rows, err := s.Database.Query("SELECT ID, [ORDER], QUESTION, ANSWER, PROJECTID FROM QUESTION where PROJECTID = ?", projectID)
	if err != nil {
		log.Print(err)
		return questions
	}
	for rows.Next() {
		var question storage.Question
		err := rows.Scan(&question.ID, &question.Order, &question.Question, &question.Answer, &question.ProjectID)
		if err != nil {
			return questions
		}
		questions = append(questions, question)
	}
	return questions
}
