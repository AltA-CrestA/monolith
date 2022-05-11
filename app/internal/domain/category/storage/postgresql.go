package storage

import (
	"monolith/pkg/client/postgresql"
	"monolith/pkg/logging"
)

type storage struct {
	client postgresql.Client
	logger logging.Logger
}
