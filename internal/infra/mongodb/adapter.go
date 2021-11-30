package mongodb

import (
	"context"
	"github.com/lucaskatayama/hexgo/internal/core/todo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Adapter struct {
	db *mongo.Database
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
