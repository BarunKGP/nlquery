package ports

import (
	"context"

	"github.com/BarunKGP/nlquery/internal/database"
)

type PersistentConn interface {
	database.DBTX
	Close(context.Context) error
}
