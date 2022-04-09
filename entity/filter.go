package entity

type Filter struct {
	UID     int    `db:"uid" json:"uid"`
	GuildID string `db:"guild_id" json:"guild_id"`
	Phrase  string `db:"phrase" json:"phrase"`
}
