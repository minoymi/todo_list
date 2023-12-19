package test

import (
	"log"
	"time"
	"todolist/domain"
	"todolist/repository"
)

func Ping(db *repository.Repo) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ping successful")
}

func QueryALL(db *repository.Repo) {
	var result domain.Todolist

	rows, err := db.Query("SELECT * FROM todolists")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&result.ID, &result.UserID, &result.Time, &result.Text, &result.Status)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result)
	}
}

func CreateTodoList(db *repository.Repo) {
	todo := domain.Todolist{
		ID:     4,
		UserID: 3,
		Time:   time.Date(2024, 10, 10, 10, 10, 10, 10, time.UTC),
		Text:   "testing my app",
		Status: "in process"}

	db.CreateTodoList(todo)
}

func DeleteTodolist(db *repository.Repo) {
	todo := domain.Todolist{
		ID:     4,
		UserID: 3,
		Time:   time.Date(2024, 10, 10, 10, 10, 10, 10, time.UTC),
		Text:   "testing my app",
		Status: "in process"}
	db.DeleteTodolist(todo.ID)
}

func UpdateTodolist(db *repository.Repo) {
	todo := domain.Todolist{
		ID:     4,
		UserID: 3,
		Time:   time.Date(9999, 10, 10, 10, 10, 10, 10, time.UTC),
		Text:   "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		Status: "HFU)#E$@HI)DFMK@P#$HF)U@$HJ"}
	db.UpdateTodolist(todo)
}
