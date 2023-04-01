package storage

import "errors"

var ErrConnectionRefused = errors.New("error connecting to grpc server")
