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
	uuid "github.com/satori/go.uuid"
	"go.opencensus.io/trace"
)

type service struct {
	tracer trace.Tracer
	repo   _repo.Repository
}

func (s service) AddFoo(ctx context.Context, foo *pb.Foo) (res *pb.Foo, err error) {
	const funcName = `AddFoo`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	// console log initialization
	ctx, consoleLog := console.Log(ctx, gvar.Logger, funcName)

	// upper log info
	level.Info(consoleLog).Log(console.LogInfo, "upper", console.LogData, foo)

	// logics
	foo.Id = uuid.NewV4().String()
	res, err = s.repo.Data.WriteFoo(ctx, foo)
	if err != nil {
		// error log
		level.Error(consoleLog).Log(console.LogErr, err)
		// span set status when error
		span.SetStatus(trace.Status{Code: int32(trace.StatusCodeInternal), Message: err.Error()})
		return res, err
	}

	createdBar, err := s.repo.Micro.BarService.AddBar(ctx, &pbBar.Bar{
		Id:          foo.Id,
		Name:        "Bar",
		Description: "This is bar from foo calling GRPC",
	})
	if err != nil {
		// error log
		level.Error(consoleLog).Log(console.LogErr, err)
		// span set status when error
		span.SetStatus(trace.Status{Code: int32(trace.StatusCodeInternal), Message: err.Error()})
		return res, err
	}
	level.Info(consoleLog).Log(console.LogInfo, "bar created", console.LogData, createdBar)

	// downer log info
	level.Info(consoleLog).Log(console.LogInfo, "downer")

	return res, nil
}

func (s service) EditFoo(ctx context.Context, foo *pb.Foo) (res *pb.Foo, err error) {
	const funcName = `EditFoo`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	// console log initialization
	ctx, consoleLog := console.Log(ctx, gvar.Logger, funcName)

	// upper log info
	level.Info(consoleLog).Log(console.LogInfo, "upper", console.LogData, foo)

	// logics
	res, err = s.repo.Data.ModifyFoo(ctx, foo)
	if err != nil {
		// error log
		level.Error(consoleLog).Log(console.LogErr, err)
		// span set status when error
		span.SetStatus(trace.Status{Code: int32(trace.StatusCodeInternal), Message: err.Error()})
		return res, err
	}

	// downer log info
	level.Info(consoleLog).Log(console.LogInfo, "downer")

	return res, nil
}

func (s service) DeleteFoo(ctx context.Context, selects *pb.Select) (res *pb.Foo, err error) {
	const funcName = `DeleteFoo`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	// console log initialization
	ctx, consoleLog := console.Log(ctx, gvar.Logger, funcName)

	// upper log info
	level.Info(consoleLog).Log(console.LogInfo, "upper", console.LogData, selects)

	// logics
	res, err = s.repo.Data.RemoveFoo(ctx, selects)
	if err != nil {
		// error log
		level.Error(consoleLog).Log(console.LogErr, err)
		// span set status when error
		span.SetStatus(trace.Status{Code: int32(trace.StatusCodeInternal), Message: err.Error()})
		return res, err
	}

	deletedBar, err := s.repo.Micro.BarService.DeleteBar(ctx, &pbBar.Select{
		Id: selects.Id,
	})
	if err != nil {
		// error log
		level.Error(consoleLog).Log(console.LogErr, err)
		// span set status when error
		span.SetStatus(trace.Status{Code: int32(trace.StatusCodeInternal), Message: err.Error()})
		return res, err
	}
	level.Info(consoleLog).Log(console.LogInfo, "bar deleted", console.LogData, deletedBar)

	// downer log info
	level.Info(consoleLog).Log(console.LogInfo, "downer")

	return res, nil
}

func (s service) GetDetailFoo(ctx context.Context, selects *pb.Select) (res *pb.Foo, err error) {
	const funcName = `GetDetailFoo`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	// console log initialization
	ctx, consoleLog := console.Log(ctx, gvar.Logger, funcName)

	// upper log info
	level.Info(consoleLog).Log(console.LogInfo, "upper", console.LogData, selects)

	// logics
	res, err = s.repo.Data.ReadDetailFoo(ctx, selects)
	if err != nil {
		// error log
		level.Error(consoleLog).Log(console.LogErr, err)
		// span set status when error
		span.SetStatus(trace.Status{Code: int32(trace.StatusCodeInternal), Message: err.Error()})
		return res, err
	}

	// downer log info
	level.Info(consoleLog).Log(console.LogInfo, "downer")

	return res, nil
}

func (s service) GetAllFoo(ctx context.Context, pagination *pb.Pagination) (res *pb.Foos, err error) {
	const funcName = `GetAllFoo`
	_, span := s.tracer.StartSpan(ctx, funcName)
	defer span.End()

	// console log initialization
	ctx, consoleLog := console.Log(ctx, gvar.Logger, funcName)

	// upper log info
	level.Info(consoleLog).Log(console.LogInfo, "upper", console.LogData, pagination)

	// logics
	res, err = s.repo.Data.ReadAllFoo(ctx, pagination)
	if err != nil {
		// error log
		level.Error(consoleLog).Log(console.LogErr, err)
		// span set status when error
		span.SetStatus(trace.Status{Code: int32(trace.StatusCodeInternal), Message: err.Error()})
		return res, err
	}

	// downer log info
	level.Info(consoleLog).Log(console.LogInfo, "downer")

	return res, nil
}

func NewService(repo _repo.Repository, tracer trace.Tracer) _interface.FooService {
	return &service{
		tracer: tracer,
		repo:   repo,
	}
}
