package main

import (
	"context"
	"fmt"

	oczipkin "contrib.go.opencensus.io/exporter/zipkin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nightsilvertech/foo/constant"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"

	"log"
	"net"
	"net/http"

	ep "github.com/nightsilvertech/foo/endpoint"
	"github.com/nightsilvertech/foo/gvar"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	"github.com/nightsilvertech/foo/repository"
	"github.com/nightsilvertech/foo/service"
	"github.com/nightsilvertech/foo/transport"
	"github.com/nightsilvertech/utl/console"
	"github.com/openzipkin/zipkin-go"
	"github.com/soheilhy/cmux"
	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func ServeGRPC(listener net.Listener, service pb.FooServiceServer, serverOptions []grpc.ServerOption) error {
	var grpcServer *grpc.Server
	if len(serverOptions) > 0 {
		grpcServer = grpc.NewServer(serverOptions...)
	} else {
		grpcServer = grpc.NewServer()
	}
	pb.RegisterFooServiceServer(grpcServer, service)
	return grpcServer.Serve(listener)
}

func ServeHTTP(listener net.Listener, service pb.FooServiceServer) error {
	mux := runtime.NewServeMux()
	err := pb.RegisterFooServiceHandlerServer(context.Background(), mux, service)
	if err != nil {
		return err
	}
	return http.Serve(listener, mux)
}

func MergeServer(service pb.FooServiceServer, serverOptions []grpc.ServerOption) {
	port := fmt.Sprintf(":%s", "1901")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(listener)
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings(
		"content-type", "application/grpc",
	))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return ServeGRPC(grpcListener, service, serverOptions) })
	g.Go(func() error { return ServeHTTP(httpListener, service) })
	g.Go(func() error { return m.Serve() })

	log.Fatal(g.Wait())
}

func main() {
	gvar.Logger = console.CreateStdGoKitLog(constant.ServiceName, false)

	reporter := httpreporter.NewReporter("http://localhost:9411/api/v2/spans")
	localEndpoint, _ := zipkin.NewEndpoint(constant.ServiceName, "http://localhost:0")
	exporter := oczipkin.NewExporter(reporter, localEndpoint)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)

	repositories, err := repository.NewRepository()
	if err != nil {
		panic(err)
	}

	services := service.NewService(*repositories)
	endpoints := ep.NewFooEndpoint(services)
	server := transport.NewFooServer(endpoints)
	MergeServer(server, nil)
}
