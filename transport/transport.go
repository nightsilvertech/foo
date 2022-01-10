package transport

import (
	"context"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	ep "github.com/nightsilvertech/foo/endpoint"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	"github.com/nightsilvertech/utl/console"
	"github.com/nightsilvertech/utl/jsonwebtoken"
)

type grpcFooServer struct {
	pb.FooServiceServer
	addFoo       grpctransport.Handler
	editFoo      grpctransport.Handler
	deleteFoo    grpctransport.Handler
	getAllFoo    grpctransport.Handler
	getDetailFoo grpctransport.Handler
}

func decodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func encodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func (g *grpcFooServer) AddFoo(ctx context.Context, Foo *pb.Foo) (*pb.Foo, error) {
	_, res, err := g.addFoo.ServeGRPC(ctx, Foo)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Foo), nil
}

func (g *grpcFooServer) EditFoo(ctx context.Context, Foo *pb.Foo) (*pb.Foo, error) {
	_, res, err := g.editFoo.ServeGRPC(ctx, Foo)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Foo), nil
}

func (g *grpcFooServer) DeleteFoo(ctx context.Context, selects *pb.Select) (*pb.Foo, error) {
	_, res, err := g.deleteFoo.ServeGRPC(ctx, selects)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Foo), nil
}

func (g *grpcFooServer) GetAllFoo(ctx context.Context, pagination *pb.Pagination) (*pb.Foos, error) {
	_, res, err := g.getAllFoo.ServeGRPC(ctx, pagination)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Foos), nil
}

func (g *grpcFooServer) GetDetailFoo(ctx context.Context, selects *pb.Select) (*pb.Foo, error) {
	_, res, err := g.getDetailFoo.ServeGRPC(ctx, selects)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Foo), nil
}

func NewFooServer(endpoints ep.FooEndpoint) pb.FooServiceServer {
	options := []grpctransport.ServerOption{
		kitoc.GRPCServerTrace(),
		grpctransport.ServerBefore(
			console.RequestIDMetadataToContext(),
			jsonwebtoken.BearerTokenMetadataToContext(),
		),
	}
	return &grpcFooServer{
		addFoo: grpctransport.NewServer(
			endpoints.AddFooEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		editFoo: grpctransport.NewServer(
			endpoints.EditFooEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		deleteFoo: grpctransport.NewServer(
			endpoints.DeleteFooEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		getAllFoo: grpctransport.NewServer(
			endpoints.GetAllFooEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		getDetailFoo: grpctransport.NewServer(
			endpoints.GetDetailFooEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
	}
}
