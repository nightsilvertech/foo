package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nightsilvertech/foo/constant"
	ep "github.com/nightsilvertech/foo/endpoint"
	"github.com/nightsilvertech/foo/gvar"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	"github.com/nightsilvertech/foo/repository"
	"github.com/nightsilvertech/foo/service"
	"github.com/nightsilvertech/foo/transport"
	"github.com/nightsilvertech/utl/console"
	"github.com/nightsilvertech/utl/preparation"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
)

type Secure struct {
	ServerCertPath     string
	ServerKeyPath      string
	ServerNameOverride string
	Service            pb.FooServiceServer
}

func (secure Secure) ServeGRPC() error {
	level.Info(gvar.Logger).Log(console.LogInfo, "serving grpc server")
	address := fmt.Sprintf("%s:%s", constant.Host, constant.GrpcPort)
	serverCert, err := tls.LoadX509KeyPair(secure.ServerCertPath, secure.ServerKeyPath)
	if err != nil {
		return err
	}
	serverOpts := []grpc.ServerOption{grpc.Creds(credentials.NewServerTLSFromCert(&serverCert))}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(serverOpts...)
	pb.RegisterFooServiceServer(grpcServer, secure.Service)
	return grpcServer.Serve(listener)
}

func (secure Secure) ServeHTTP() error {
	level.Info(gvar.Logger).Log(console.LogInfo, "serving http server")
	httpAddress := fmt.Sprintf("%s:%s", constant.Host, constant.HttpPort)
	grpcAddress := fmt.Sprintf("%s:%s", constant.Host, constant.GrpcPort)
	clientCert, err := credentials.NewClientTLSFromFile(secure.ServerCertPath, secure.ServerNameOverride)
	if err != nil {
		return err
	}
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(clientCert)}

	mux := runtime.NewServeMux()
	pb.RegisterFooServiceServer(grpc.NewServer(), secure.Service)
	err = pb.RegisterFooServiceHandlerFromEndpoint(context.Background(), mux, grpcAddress, dialOptions)
	if err != nil {
		return err
	}
	return http.ListenAndServeTLS(httpAddress, secure.ServerCertPath, secure.ServerKeyPath, mux)
}

func Serve(service pb.FooServiceServer) {
	secure := Secure{
		ServerCertPath:     "C:\\Users\\Asus\\Desktop\\tls\\foo\\server.crt",
		ServerKeyPath:      "C:\\Users\\Asus\\Desktop\\tls\\foo\\server.key",
		ServerNameOverride: "0.0.0.0",
		Service:            service,
	}

	g := new(errgroup.Group)
	g.Go(func() error { return secure.ServeGRPC() })
	g.Go(func() error { return secure.ServeHTTP() })
	log.Fatal(g.Wait())
}

func main() {
	prepare := preparation.Data{
		//LoggingFilePath:            "C:\\Users\\Asus\\Desktop\\service.log",
		TracerUrl:                  "http://localhost:9411/api/v2/spans",
		CircuitBreakerTimeout:      constant.CircuitBreakerTimout,
		ServiceName:                constant.ServiceName,
		ZipkinEndpointPort:         constant.ZipkinHostPort,
		Debug:                      false,
		FractionProbabilitySampler: 1,
	}
	prepare.CircuitBreaker()
	gvar.Logger = prepare.Logger()
	gvar.Tracer = prepare.Tracer()

	repositories := repository.NewRepository()
	services := service.NewService(*repositories)
	endpoints := ep.NewFooEndpoint(services)
	server := transport.NewFooServer(endpoints)
	Serve(server)
}
