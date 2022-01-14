package data

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/nightsilvertech/foo/gvar"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	pb "github.com/nightsilvertech/foo/protoc/api/v1"
	_interface "github.com/nightsilvertech/foo/repository/interface"
	"github.com/nightsilvertech/utl/errwrap"
)

var mutex = &sync.RWMutex{}

type dataReadWrite struct {
	db *sql.DB
}

func (d *dataReadWrite) WriteFoo(ctx context.Context, req *pb.Foo) (res *pb.Foo, err error) {
	const funcName = `WriteFoo`
	ctx, span := gvar.Tracer.StartSpan(ctx, funcName)
	defer span.End()

	currentTime := time.Now()
	req.CreatedAt = currentTime.Unix()
	req.UpdatedAt = currentTime.Unix()
	stmt, err := d.db.Prepare(`
	INSERT INTO foos(id, name, description, created_at, updated_at) VALUES (?,?,?,?,?)
	`)
	if err != nil {
		return res, errwrap.Wrap(funcName, "db.Prepare", err)
	}
	result, err := stmt.ExecContext(
		ctx,
		req.Id,          // id
		req.Name,        // name
		req.Description, // description
		currentTime,     // created_at
		currentTime,     // updated_at
	)
	if err != nil {
		return res, errwrap.Wrap(funcName, "stmt.ExecContext", err)
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return res, errwrap.Wrap(funcName, "result.RowsAffected", err)
	}
	return req, nil
}

func (d *dataReadWrite) ModifyFoo(ctx context.Context, req *pb.Foo) (res *pb.Foo, err error) {
	const funcName = `ModifyFoo`
	ctx, span := gvar.Tracer.StartSpan(ctx, funcName)
	defer span.End()

	currentTime := time.Now()
	req.UpdatedAt = currentTime.Unix()
	stmt, err := d.db.Prepare(`
	UPDATE foos
	SET name = ?, description = ?
	WHERE id = ?
	`)
	if err != nil {
		return res, errwrap.Wrap(funcName, "db.Prepare", err)
	}
	result, err := stmt.ExecContext(
		ctx,
		req.Name,        // name
		req.Description, // description
		req.Id,          // id
	)
	if err != nil {
		return res, errwrap.Wrap(funcName, "stmt.ExecContext", err)
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return res, errwrap.Wrap(funcName, "result.RowsAffected", err)
	}
	return req, nil
}

func (d *dataReadWrite) RemoveFoo(ctx context.Context, req *pb.Select) (res *pb.Foo, err error) {
	const funcName = `RemoveFoo`
	ctx, span := gvar.Tracer.StartSpan(ctx, funcName)
	defer span.End()

	stmt, err := d.db.Prepare(`SELECT * FROM foos WHERE id = ?`)
	if err != nil {
		return res, errwrap.Wrap(funcName, "db.Prepare", err)
	}
	mutex.Lock()
	row := stmt.QueryRowContext(ctx, req.Id)
	mutex.Unlock()

	var foo pb.Foo
	var createdAt, updatedAt time.Time
	err = row.Scan(
		&foo.Id,          // id
		&foo.Name,        // name
		&foo.Description, // description
		&createdAt,       // created_at
		&updatedAt,       // updated_at
	)
	if err != nil {
		return res, errwrap.Wrap(funcName, "row.Scan", err)
	}
	foo.CreatedAt = createdAt.Unix()
	foo.UpdatedAt = updatedAt.Unix()

	stmt, err = d.db.Prepare(`DELETE FROM foos WHERE id = ?`)
	if err != nil {
		return res, errwrap.Wrap(funcName, "db.Prepare", err)
	}
	result, err := stmt.ExecContext(
		ctx,
		req.Id, // id
	)
	if err != nil {
		return res, errwrap.Wrap(funcName, "stmt.ExecContext", err)
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return res, errwrap.Wrap(funcName, "result.RowsAffected", err)
	}
	return &foo, nil
}

func (d *dataReadWrite) ReadDetailFoo(ctx context.Context, selects *pb.Select) (res *pb.Foo, err error) {
	const funcName = `ReadDetailFoo`
	ctx, span := gvar.Tracer.StartSpan(ctx, funcName)
	defer span.End()

	stmt, err := d.db.Prepare(`SELECT * FROM foos WHERE id = ?`)
	if err != nil {
		return res, errwrap.Wrap(funcName, "db.Prepare", err)
	}
	mutex.Lock()
	row := stmt.QueryRowContext(ctx, selects.Id)
	mutex.Unlock()

	var Foo pb.Foo
	var createdAt, updatedAt time.Time
	err = row.Scan(
		&Foo.Id,          // id
		&Foo.Name,        // name
		&Foo.Description, // description
		&createdAt,       // created_at
		&updatedAt,       // updated_at
	)
	if err != nil {
		return res, errwrap.Wrap(funcName, "row.Scan", err)
	}
	Foo.CreatedAt = createdAt.Unix()
	Foo.UpdatedAt = updatedAt.Unix()
	return &Foo, nil
}

func (d *dataReadWrite) ReadAllFoo(ctx context.Context, req *pb.Pagination) (res *pb.Foos, err error) {
	const funcName = `ReadAllFoo`
	ctx, span := gvar.Tracer.StartSpan(ctx, funcName)
	defer span.End()

	stmt, err := d.db.Prepare(`SELECT * FROM foos ORDER BY created_at DESC`)
	if err != nil {
		return res, errwrap.Wrap(funcName, "db.Prepare", err)
	}
	mutex.Lock()
	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return res, errwrap.Wrap(funcName, "stmt.QueryContext", err)
	}
	mutex.Unlock()
	defer row.Close()

	var foos pb.Foos
	var Foo pb.Foo
	var createdAt, updatedAt time.Time
	for row.Next() {
		err = row.Scan(
			&Foo.Id,          // id
			&Foo.Name,        // name
			&Foo.Description, // description
			&createdAt,       // created_at
			&updatedAt,       // updated_at
		)
		if err != nil {
			return res, errwrap.Wrap(funcName, "row.Scan", err)
		}
		foos.Foos = append(foos.Foos, &pb.Foo{
			Id:          Foo.Id,
			Name:        Foo.Name,
			Description: Foo.Description,
			CreatedAt:   createdAt.Unix(),
			UpdatedAt:   updatedAt.Unix(),
		})
	}
	return &foos, nil
}

func NewDataReadWriter(username, password, host, port, name string) (_interface.DRW, error) {
	const funcName = `NewDataReadWriter`

	databaseUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username,
		password,
		host,
		port,
		name,
	)
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		return nil, errwrap.Wrap(funcName, "sql.Open", err)
	}

	return &dataReadWrite{
		db: db,
	}, nil
}
