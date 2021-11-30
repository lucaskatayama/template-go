package mysql

import (
	"context"
	"database/sql"
	"github.com/lucaskatayama/hexgo/internal/core/todo"
)

type Adapter struct {
	db *sql.DB
}

func New() Adapter {
	return Adapter{
		db: db,
	}
}

func (a Adapter) Create(ctx context.Context) (todo.Todo, error) {
	panic("implement me")
}

func (a Adapter) List(ctx context.Context) ([]todo.Todo, error) {
	return nil, nil
}
