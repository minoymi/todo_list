package main

import (
	repo "todolist/repository"
	"todolist/router"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db := repo.ConnectToRepo()
	defer db.Close()
	router.StartHttpServer(db)
}
