package microservice

import (
	barsvc "github.com/nightsilvertech/bar/service/interface"
	bardialer "github.com/nightsilvertech/bar/transport"
	"google.golang.org/grpc"
)

type Microservices struct {
	Connections []*grpc.ClientConn
	BarService  barsvc.BarService
}

func NewMicroservice() *Microservices {
	var connections []*grpc.ClientConn
	barSvc, barConn, err := bardialer.DialBarService("localhost:1900")
	if err != nil {
		return nil
	}
	connections = append(connections, barConn)
	return &Microservices{
		BarService: barSvc,
	}
}
