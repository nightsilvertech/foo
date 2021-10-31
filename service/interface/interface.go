package _interface

import (
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
)

type FooService interface {
	pb.FooServiceServer
}
