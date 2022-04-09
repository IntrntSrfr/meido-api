package entity

import (
	"time"
)

type Warn struct {
	UID         int        `db:"uid" json:"uid"`
	GuildID     string     `db:"guild_id" json:"guild_id"`
	UserID      string     `db:"user_id" json:"user_id"`
	Reason      string     `db:"reason" json:"reason"`
	GivenByID   string     `db:"given_by_id" json:"given_by_id"`
	GivenAt     time.Time  `db:"given_at" json:"given_at"`
	IsValid     bool       `db:"is_valid" json:"is_valid"`
	ClearedByID *string    `db:"cleared_by_id" json:"cleared_by_id"`
	ClearedAt   *time.Time `db:"cleared_at" json:"cleared_at"`
}
