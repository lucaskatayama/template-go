package memory

import (
	"context"
	"fmt"
	"time"

	"github.com/lucaskatayama/hexgo/internal/core/todo"
)

type Adapter struct {
}

func (a Adapter) Create(ctx context.Context) (todo.Todo, error) {
	return todo.Todo{
		Description: "test",
		At:          time.Now(),
	}, nil
}

func (a Adapter) List(ctx context.Context) ([]todo.Todo, error) {
	var list []todo.Todo
	for i := 0; i < 10; i++ {
		list = append(list, todo.Todo{
			Description: fmt.Sprintf("TODO %d", i),
			At:          time.Now(),
		})
	}
	return list, nil
}

func New() todo.Port {
	return Adapter{}
}
