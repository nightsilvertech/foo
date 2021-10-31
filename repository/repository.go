package repository

import (
	"github.com/nightsilvertech/foo/repository/data"
	_interface "github.com/nightsilvertech/foo/repository/interface"
)

type Repository struct {
	Data  _interface.DRW
	Cache _interface.CRW
}

func NewRepository() (repo *Repository,err error){
	dataReadWriter, err := data.NewDataReadWriter("root", "root", "localhost", "3306", "foobar")
	if err != nil {
		return repo, err
	}
	return &Repository{
		Data: dataReadWriter,
	}, nil
}
