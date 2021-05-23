package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/axeldeveloper/go-gqlgen-todos/api/dal"
	"github.com/axeldeveloper/go-gqlgen-todos/graph/generated"
	"github.com/axeldeveloper/go-gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:        fmt.Sprintf("T%d", rand.Int()),
		Text:      input.Text,
		Done:      true,
		UserID:    input.UserID,
		CreatedAt: time.Now().UTC(),
	}

	db, err := dal.Connect()
	//statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS todo (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT NOT NULL, done BOOLEAN DEFAULT(TRUE), created_at DATETIME NULL,  user_id INTEGER  NOT NULL )")
	//statement.Exec()

	//stm, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, email TEXT NOT NULL, created_at DATETIME NULL)")
	//stm.Exec()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sql := "INSERT INTO todo (text, done, user_id, created_at) VALUES(?, ?, ?, ?)"

	//sql2 := "INSERT INTO users (name, email) VALUES(?, ?)"
	///_, _ = dal.LogAndExec(db, sql2, "Axel Alexander", "axel_nomore@hotmail.com")

	//rows, err := dal.LogAndQuery(db, sql, todo.Text, todo.Done, todo.UserID, todo.CreatedAt)
	rows, err := dal.LogAndExec(db, sql, todo.Text, todo.Done, todo.UserID, todo.CreatedAt)
	if err != nil {
		fmt.Println("Entrei no next")
		return todo, err
	}

	todo.ID = strconv.FormatInt(rows, 10)

	/*
		if err != nil || !rows.Next() {
			fmt.Println("Entrei no next")
			return todo, err
		}
		defer rows.Close()

		if err := rows.Scan(&todo.ID); err != nil {
			fmt.Println("entrei no scam")
			fmt.Errorf(err.Error())
			return &model.Todo{}, err
		}
	*/

	db.Close()
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {

	fmt.Println("Entrei no  Todos -> ")

	var todos []*model.Todo
	db, err := dal.Connect()
	if err != nil {
		panic(" não conectado")
	}

	rows, err := dal.LogAndQuery(db, "SELECT id, text, user_id, created_at FROM todo")

	if err != nil {
		fmt.Println("query failed", err)
	}

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Text, &todo.UserID, &todo.CreatedAt); err != nil {
			fmt.Println("Next failed", err)
			return nil, err
		}

		fmt.Println("Entrei no  Todos -> append ", todo.ID)
		todos = append(todos, &todo)
	}
	defer rows.Close()
	db.Close()
	return todos, nil
	//return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	//panic(fmt.Errorf("not implemented"))

	fmt.Println("Entrei no  User -> UserID", obj.UserID)
	db, err := dal.Connect()
	if err != nil {
		panic(" não conectado")
	}

	rows, _ := dal.LogAndQuery(db, "SELECT id, name FROM users  where id = ?", obj.UserID)
	defer rows.Close()
	//
	if !rows.Next() {
		return &model.User{}, nil
	}
	var user model.User
	if err := rows.Scan(&user.ID, &user.Name); err != nil {
		return &model.User{}, err
	}
	fmt.Println("User ID ", &user.ID)
	fmt.Println("User Name ", &user.Name)
	return &user, nil
	//return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
