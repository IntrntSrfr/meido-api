package entity

type Guild struct {
	GuildID  string `db:"guild_id" json:"guild_id"`
	UseWarns bool   `db:"use_warns" json:"use_warns"`
	MaxWarns int    `db:"max_warns" json:"max_warns"`

	// described in days, 0 means infinite duration
	WarnDuration int `db:"warn_duration" json:"warn_duration"`
	//AutoRole     string `db:"autorole"`

	AutomodLogChannelID string `db:"automod_log_channel_id" json:"automod_log_channel_id"`
	FishingChannelID    string `db:"fishing_channel_id" json:"fishing_channel_id"`

	Warns       []*Warn     `json:"warns" db:"-"`
	MemberRoles []*UserRole `json:"member_roles" db:"-"`
	Filters     []*Filter   `json:"filters" db:"-"`
}
