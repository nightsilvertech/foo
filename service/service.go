package service

import (
	"context"

	"github.com/nightsilvertech/foo/gvar"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	_repo "github.com/nightsilvertech/foo/repository"
	_interface "github.com/nightsilvertech/foo/service/interface"
	"github.com/nightsilvertech/foo/util"
	uuid "github.com/satori/go.uuid"
)

type service struct {
	repo _repo.Repository
}

func (s service) AddFoo(ctx context.Context, Foo *pb.Foo) (*pb.Foo, error) {
	const funcName = `AddFoo`
	Foo.Id = uuid.NewV4().String()
	return s.repo.Data.WriteFoo(ctx, Foo)
}

func (s service) EditFoo(ctx context.Context, Foo *pb.Foo) (*pb.Foo, error) {
	const funcName = `EditFoo`
	return s.repo.Data.ModifyFoo(ctx, Foo)
}

func (s service) DeleteFoo(ctx context.Context, selects *pb.Select) (*pb.Foo, error) {
	const funcName = `DeleteFoo`
	return s.repo.Data.RemoveFoo(ctx, selects)
}

func (s service) GetDetailFoo(ctx context.Context, selects *pb.Select) (*pb.Foo, error) {
	const funcName = `GetDetailFoo`
	return s.repo.Data.ReadDetailFoo(ctx, selects)
}

func (s service) GetAllFoo(ctx context.Context, pagination *pb.Pagination) (*pb.Foos, error) {
	const funcName = `GetAllFoo`
	console := util.ConsoleLog(gvar.Logger)
	console.Log("test", "helo")
	return s.repo.Data.ReadAllFoo(ctx, pagination)
}

func NewService(repo _repo.Repository) _interface.FooService {
	return &service{
		repo: repo,
	}
}
