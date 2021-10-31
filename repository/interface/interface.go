package _interface

import (
	"context"

	pb "github.com/nightsilvertech/foo/protoc/api/v1"
)

// DRW means data read writer this interface
// contains all data management function
type DRW interface {
	WriteFoo(ctx context.Context, req *pb.Foo) (res *pb.Foo, err error)
	ModifyFoo(ctx context.Context, req *pb.Foo) (res *pb.Foo, err error)
	RemoveFoo(ctx context.Context, req *pb.Select) (res *pb.Foo, err error)
	ReadDetailFoo(ctx context.Context, req *pb.Select) (res *pb.Foo, err error)
	ReadAllFoo(ctx context.Context, req *pb.Pagination) (res *pb.Foos, err error)
}

// CRW means cache read writer this interface
// contains all cache management function
type CRW interface {

}
