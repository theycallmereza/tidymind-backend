package database

import (
	"github.com/theycallmereza/tidymind-backend/graph/model"
)

func ToGraphQLUser(u *User) *model.User {
	if u == nil {
		return nil
	}
	tasks := make([]*model.Task, len(u.Tasks))
	for i, task := range u.Tasks {
		tasks[i] = ToGraphQLTask(&task)
	}
	return &model.User{
		ID:       u.ID.String(),
		Username: u.Username,
		Email:    u.Email,
		Tasks:    tasks,
	}
}

func ToGraphQLTask(t *Task) *model.Task {
	if t == nil {
		return nil
	}
	return &model.Task{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		User:        ToGraphQLUser(&t.User), // or nil if lazy-loaded
	}
}
