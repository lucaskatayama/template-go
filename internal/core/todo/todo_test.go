package todo_test

import (
	"context"
	"github.com/lucaskatayama/hexgo/internal/core/todo"
	"github.com/lucaskatayama/hexgo/internal/infra/mongodb"
	"testing"
)

func TestSimple(t *testing.T) {
	p := mongodb.New()

	srv := todo.New(p)


	srv.Create(context.Background())
}
