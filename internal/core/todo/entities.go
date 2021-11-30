package todo

import (
	"context"
	"time"
)

type Port interface{
	Create(ctx context.Context) (Todo, error)
	List(ctx context.Context) ([]Todo, error)
}

type Todo struct {
	Description string    `json:"description,omitempty"`
	At          time.Time `json:"at"`
}

