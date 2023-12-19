package repository

import (
	"database/sql"
	"log"
	"todolist/domain"
)

var databaseCredentials string = "user=postgres password=1 host=localhost port=5432 database=todo sslmode=disable"

type Repo struct {
	*sql.DB
}

func ConnectToRepo() *Repo {
	db, err := sql.Open("pgx", databaseCredentials)
	if err != nil {
		log.Fatal(err)
	}
	return &Repo{db}
}

func (db *Repo) ReadAllTodolistsForUser(userid int64) ([]domain.Todolist, error) {
	rows, err := db.Query("SELECT * FROM todolists WHERE userID = $1;", userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRowsIntoSlice(rows)
}

func (db *Repo) CreateTodoList(data domain.Todolist) error {
	_, err := db.Exec("INSERT INTO todolists (userid, timedate, tasks, misc) VALUES ($1, $2, $3, $4)", data.UserID, data.Time, data.Text, data.Status)
	if err != nil {
		return err
	}
	return nil
}

func (db *Repo) DeleteTodolist(ID int64) error {
	_, err := db.Exec("DELETE FROM todolists WHERE ID = $1", ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *Repo) UpdateTodolist(data domain.Todolist) error {
	_, err := db.Exec("UPDATE todolists SET timedate = $1, tasks = $2, misc = $3 WHERE id = $4", data.Time, data.Text, data.Status, data.ID)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoSlice(rows *sql.Rows) ([]domain.Todolist, error) {
	var slc []domain.Todolist

	for rows.Next() {
		var result domain.Todolist
		err := rows.Scan(&result.ID, &result.UserID, &result.Time, &result.Text, &result.Status)
		if err != nil {
			return nil, err
		}
		slc = append(slc, result)
	}
	return slc, nil
}
