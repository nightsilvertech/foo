package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	_interface "github.com/nightsilvertech/foo/service/interface"
)

func makeAddFooEndpoint(usecase _interface.FooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.AddFoo(ctx, request.(*pb.Foo))
		return res, err
	}
}

func (e FooEndpoint) AddFoo(ctx context.Context, req *pb.Foo) (*pb.Foo, error) {
	res, err := e.AddFooEndpoint(ctx, req)
	if err != nil {
		return &pb.Foo{}, err
	}
	return res.(*pb.Foo), nil
}

