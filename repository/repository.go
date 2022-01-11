package repository

import (
	"github.com/nightsilvertech/foo/repository/data"
	_interface "github.com/nightsilvertech/foo/repository/interface"
	"github.com/nightsilvertech/foo/repository/microservice"
	"go.opencensus.io/trace"
)

type Repository struct {
	tracer trace.Tracer
	Data   _interface.DRW
	Cache  _interface.CRW
	Micro  microservice.Microservices
}

func NewRepository(tracer trace.Tracer) (repo *Repository) {
	dataReadWriter, _ := data.NewDataReadWriter("root", "root", "localhost", "3306", "foobar", tracer)
	return &Repository{
		tracer: tracer,
		Data:   dataReadWriter,
		Micro:  *microservice.NewMicroservice(),
	}
}
