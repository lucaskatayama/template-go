package todo

import (
	"context"
	"github.com/lucaskatayama/hexgo/pkg/log"
)

func New(port Port) Service {
	return Service{port}
}

type Service struct {
	port Port
}

func (s Service) List(ctx context.Context) ([]Todo, error) {
	log.With("a", 1).Debugf("hello %s", "john doe")
	return s.port.List(ctx)
}

func (s Service) Create(ctx context.Context) (Todo, error) {
	panic("implement me")
}

func (s Service) Update(ctx context.Context, id string, desc string) (Todo, error) {
	panic("implement me")
}
