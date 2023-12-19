package domain

import (
	"fmt"
	"time"
)

type Todolist struct {
	ID     int64
	UserID int64
	Time   time.Time
	Text   string
	Status string
}

type repository interface {
	CreateTodoList(Todolist) error
	ReadAllTodolistsForUser(int64) ([]Todolist, error)
	UpdateTodolist(Todolist) error
	DeleteTodolist(int64) error
}

const maxTextLength = 1023
const maxStatusLength = 255

func Create(todo Todolist, repo repository) error {
	if err := validateInput(todo); err != nil {
		return err
	}

	return repo.CreateTodoList(todo)
}

func ReadAllForGivenUser(UserID int64, repo repository) ([]Todolist, error) {
	return repo.ReadAllTodolistsForUser(UserID)
}

func Update(todo Todolist, repo repository) error {
	if err := validateInput(todo); err != nil {
		return err
	}

	return repo.UpdateTodolist(todo)
}

func Delete(ID int64, repo repository) error {
	return repo.DeleteTodolist(ID)
}

func validateInput(todo Todolist) error {
	if len(todo.Text) > maxTextLength {
		return fmt.Errorf("text too long, max is %d characters", maxTextLength)
	}
	if len(todo.Status) > maxStatusLength {
		return fmt.Errorf("status too long, max is %d characters", maxStatusLength)
	}
	return nil
}
