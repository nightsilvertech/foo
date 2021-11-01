package service

import (
	"context"
	"github.com/go-kit/kit/log/level"
	pbBar "github.com/nightsilvertech/bar/protoc/api/v1"
	"github.com/nightsilvertech/foo/gvar"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	_repo "github.com/nightsilvertech/foo/repository"
	_interface "github.com/nightsilvertech/foo/service/interface"
	"github.com/nightsilvertech/utl/console"
	"github.com/nightsilvertech/utl/pass"
	uuid "github.com/satori/go.uuid"
)

type service struct {
	repo _repo.Repository
}

func (s service) AddFoo(ctx context.Context, foo *pb.Foo) (res *pb.Foo, err error) {
	const funcName = `AddFoo`

	ctx, consoleLog := console.Log(ctx, gvar.Logger, funcName)

	level.Info(consoleLog).Log(console.LogInfo, "upper", console.LogData, foo)

	foo.Id = uuid.NewV4().String()

	createdBar, err := s.repo.Micro.BarService.AddBar(pass.Create(ctx, "THISISIDDDD"), &pbBar.Bar{
		Name:        "Bar 1",
		Description: "This is bar 1",
	})
	if err != nil {
		level.Error(consoleLog).Log(console.LogErr, err)
		return res, err
	}

	level.Info(consoleLog).Log(console.LogData, createdBar)

	level.Info(consoleLog).Log(console.LogInfo, "downer")

	return s.repo.Data.WriteFoo(ctx, foo)
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
	return s.repo.Data.ReadAllFoo(ctx, pagination)
}

func NewService(repo _repo.Repository) _interface.FooService {
	return &service{
		repo: repo,
	}
}
