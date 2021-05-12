package repository

import "context"

type Queryer interface {
	Exec(ctx context.Context, query string, params map[string]interface{}) (interface{}, error)
}

type queryer struct {
	driver interface{}
}

func NewRepository(d interface{}) Queryer {
	return &queryer{
		driver: d,
	}
}

func (q *queryer) Exec(ctx context.Context, query string, params map[string]interface{}) (interface{}, error) {
	return nil, nil
}
