package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	_interface "github.com/nightsilvertech/foo/service/interface"
)

func makeGetAllFooEndpoint(usecase _interface.FooService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res, err := usecase.GetAllFoo(ctx, request.(*pb.Pagination))
		return res, err
	}
}

func (e FooEndpoint) GetAllFoo(ctx context.Context, req *pb.Pagination) (*pb.Foos, error) {
	res, err := e.GetAllFooEndpoint(ctx, req)
	if err != nil {
		return &pb.Foos{}, err
	}
	return res.(*pb.Foos), nil
}

