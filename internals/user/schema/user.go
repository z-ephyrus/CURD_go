package schema

import (
	"time"

	"github.com/uptrace/bun"
)

type Wser struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:",notnull"`
	Email     string    `bun:",unique,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}