package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	_interface "github.com/nightsilvertech/foo/service/interface"
)

func makeGetDetailFooEndpoint(usecase _interface.FooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.GetDetailFoo(ctx, request.(*pb.Select))
		return res, err
	}
}

func (e FooEndpoint) GetDetailFoo(ctx context.Context, req *pb.Select) (*pb.Foo, error) {
	res, err := e.GetDetailFooEndpoint(ctx, req)
	if err != nil {
		return res.(*pb.Foo), err
	}
	return res.(*pb.Foo), nil
}
