package db

import (
	"github.com/intrntsrfr/meido-api/entity"
	_ "github.com/lib/pq" //import postgres
)

type ApiDB interface {
	GetGuilds() []*entity.Guild
	GetGuild(guildID string) (*entity.Guild, error)
	GetGuildMemberRoles(guildID string) []*entity.UserRole
	GetGuildFilters(guildID string) []*entity.Filter
	GetGuildWarns(guildID string) []*entity.Warn
}
