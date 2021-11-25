package config

import "errors"

var ( // biz error
	ErrRowsAffectedInvalid = errors.New("db update rows affected invalid")
)

var ( // sys error
	ErrSysPanic = errors.New("system panic")
)
