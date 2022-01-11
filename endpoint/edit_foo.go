package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	_interface "github.com/nightsilvertech/foo/service/interface"
)

func makeEditFooEndpoint(usecase _interface.FooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.EditFoo(ctx, request.(*pb.Foo))
		return res, err
	}
}

func (e FooEndpoint) EditFoo(ctx context.Context, req *pb.Foo) (*pb.Foo, error) {
	res, err := e.EditFooEndpoint(ctx, req)
	if err != nil {
		return &pb.Foo{}, err
	}
	return res.(*pb.Foo), nil
}

