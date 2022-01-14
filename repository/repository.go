package repository

import (
	"github.com/nightsilvertech/foo/repository/data"
	_interface "github.com/nightsilvertech/foo/repository/interface"
	"github.com/nightsilvertech/foo/repository/microservice"
)

type Repository struct {
	Data   _interface.DRW
	Cache  _interface.CRW
	Micro  microservice.Microservices
}

func NewRepository() (repo *Repository) {
	dataReadWriter, _ := data.NewDataReadWriter("root", "root", "localhost", "3306", "foobar")
	return &Repository{
		Data:   dataReadWriter,
		Micro:  *microservice.NewMicroservice(),
	}
}
