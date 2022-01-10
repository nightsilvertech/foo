package microservice

import (
	barsvc "github.com/nightsilvertech/bar/service/interface"
	bardialer "github.com/nightsilvertech/bar/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Microservices struct {
	Connections []*grpc.ClientConn
	BarService  barsvc.BarService
}

func NewMicroservice() *Microservices {
	var connections []*grpc.ClientConn

	tlsCredentials, err := credentials.NewClientTLSFromFile(
		"C:\\Users\\Asus\\Desktop\\tls\\bar\\server.crt",
		"0.0.0.0",
	)
	if err != nil {
		return nil
	}
	barSvc, barConn, err := bardialer.DialBarService("0.0.0.0", "9081", tlsCredentials)
	if err != nil {
		return nil
	}

	connections = append(connections, barConn)
	return &Microservices{
		BarService: barSvc,
	}
}
