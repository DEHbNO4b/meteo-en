package filesource

import "errors"

var (
	ErrEmptyDataSource   = errors.New("data source path is required")
	ErrInvalidDataString = errors.New("invalid data in source string")
	ErrEmptyData         = errors.New("empty data")
)
