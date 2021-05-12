package factory

import (
	"log"

	"go-boiler-plate/repository"
)

type Factory interface {
	QueryExecer() repository.Queryer
}

type factory struct {
	uri string
}

func NewFactory(uri string) Factory {
	return &factory{
		uri: uri,
	}
}

func (f *factory) driver() (interface{}, error) {
	return nil, nil
}

func (f *factory) QueryExecer() repository.Queryer {
	d, err := f.driver()
	if err != nil {
		log.Fatalf("Unable to establish connection: %s", err)
	}

	return repository.NewRepository(d)
}
