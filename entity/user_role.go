package entity

type UserRole struct {
	UID     int    `db:"uid" json:"uid"`
	GuildID string `db:"guild_id" json:"guild_id"`
	UserID  string `db:"user_id" json:"user_id"`
	RoleID  string `db:"role_id" json:"role_id"`
}
