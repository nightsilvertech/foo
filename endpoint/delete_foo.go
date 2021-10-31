package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	_interface "github.com/nightsilvertech/foo/service/interface"
)

func makeDeleteFooEndpoint(usecase _interface.FooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.DeleteFoo(ctx, request.(*pb.Select))
		return res, err
	}
}

func (e FooEndpoint) DeleteFoo(ctx context.Context, req *pb.Select) (*pb.Foo, error) {
	res, err := e.DeleteFooEndpoint(ctx, req)
	if err != nil {
		return res.(*pb.Foo), err
	}
	return res.(*pb.Foo), nil
}
