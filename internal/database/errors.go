package database

import (
	"errors"

	"github.com/lib/pq"
)

const (
	UniqueViolation pq.ErrorCode = "23505"
)

func IsError(err error, code pq.ErrorCode) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == code {
		return true
	}

	return false
}
