package cache

import (
	"database/sql"

	_interface "github.com/nightsilvertech/foo/repository/interface"
)

type cacheReadWrite struct {
	db *sql.DB
}

func NewCacheReadWriter() _interface.CRW {
	return &cacheReadWrite{}
}
