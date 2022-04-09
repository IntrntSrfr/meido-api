package entity

import (
	"time"
)

type CommandLog struct {
	UID       int       `db:"uid" json:"uid"`
	Command   string    `db:"command" json:"command"`
	Args      string    `db:"args" json:"args"`
	UserID    string    `db:"user_id" json:"user_id"`
	GuildID   string    `db:"guild_id" json:"guild_id"`
	ChannelID string    `db:"channel_id" json:"channel_id"`
	MessageID string    `db:"message_id" json:"message_id"`
	SentAt    time.Time `db:"sent_at" json:"sent_at"`
}
