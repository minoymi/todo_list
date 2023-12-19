package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"todolist/domain"
	"todolist/repository"
)

var repo *repository.Repo

type handler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		http.Error(w, err.Message, err.Code)
		log.Println(err.Error, err.Message, err.Code)
	}
}

func StartHttpServer(r *repository.Repo) {
	repo = r

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.Handle("/todolist", handler(todolistController))

	log.Fatal(s.ListenAndServe())
}

func todolistController(w http.ResponseWriter, r *http.Request) *appError {
	switch r.Method {
	case http.MethodGet:
		return GET_todolist(w, r)
	case http.MethodPost:
		return POST_todolist(w, r)
	case http.MethodPut:
		return PUT_todolist(w, r)
	case http.MethodDelete:
		return DELETE_todolist(w, r)
	default:
		return &appError{nil, "invalid http method", 400}
	}
}

func GET_todolist(w http.ResponseWriter, r *http.Request) *appError {
	if len(r.Header["Userid"]) == 0 {
		return &appError{nil, "no userid", 400}
	}

	Userid, err := strconv.ParseInt(r.Header["Userid"][0], 10, 64)
	if err != nil {
		return &appError{err, "invalid userid", 400}
	}

	todolists, err := domain.ReadAllForGivenUser(Userid, repo)
	if err != nil {
		return &appError{err, "internal server error", 500}
	}

	response, err := json.Marshal(todolists)
	if err != nil {
		return &appError{err, "internal server error", 500}
	}

	_, err = w.Write(response)
	if err != nil {
		return &appError{err, "internal server error", 500}
	}
	return nil
}

func POST_todolist(w http.ResponseWriter, r *http.Request) *appError {
	todo := domain.Todolist{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		return &appError{err, "bad request", 400}
	}

	err = domain.Create(todo, repo)
	if err != nil {
		return &appError{err, "bad request", 400}
	}

	return nil
}

func PUT_todolist(w http.ResponseWriter, r *http.Request) *appError {
	todo := domain.Todolist{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		return &appError{err, "bad request", 400}
	}
	err = domain.Update(todo, repo)
	if err != nil {
		return &appError{err, "internal server error", 500}
	}

	return nil
}

func DELETE_todolist(w http.ResponseWriter, r *http.Request) *appError {
	if len(r.Header["Id"]) == 0 {
		return &appError{nil, "no todolist id", 400}
	}

	id, err := strconv.ParseInt(r.Header["Id"][0], 10, 64)
	if err != nil {
		return &appError{err, "invalid todolist id", 400}
	}

	err = domain.Delete(id, repo)
	if err != nil {
		return &appError{err, "internal server error", 500}
	}
	return nil
}
