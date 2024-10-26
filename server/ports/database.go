package ports

import (
	"context"

	"github.com/BarunKGP/nlquery/core/database"
)

type PersistentConn interface {
	database.DBTX
	Close(context.Context) error
}
