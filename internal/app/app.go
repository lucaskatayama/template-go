package app

import (
	"github.com/lucaskatayama/hexgo/internal/core/todo"
	"github.com/lucaskatayama/hexgo/internal/infra/memory"
)

var (
	Todo todo.Service
)

func Config() {

	//mongodb.Connect(ctx)
	// mysql.Connect(ctx)

	Todo = todo.New(memory.New())
	//Todo = todo.New(mongodb.New())
	// Todo = todo.New(mysql.New())

}
