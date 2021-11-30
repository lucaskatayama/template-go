package pb

import (
	"context"

	"github.com/lucaskatayama/hexgo/internal/app"
)

type TodoServerImpl struct {
	UnimplementedTodoServer
}

func (s TodoServerImpl) ListTodo(ctx context.Context, request *TodoRequest) (*TodoResponse, error) {
	todos, err := app.Todo.List(ctx)

	var ts []*TodoItem
	for _, i := range todos {
		ts = append(ts, &TodoItem{Description: i.Description})
	}

	if err != nil {
		return nil, err
	}
	return &TodoResponse{
		Todos: ts,
	}, nil
}
