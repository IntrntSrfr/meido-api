package db

import (
	"github.com/intrntsrfr/meido-api/entity"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type PsqlDB struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewPsqlDB(db *sqlx.DB, log *zap.Logger) *PsqlDB {
	return &PsqlDB{db: db, log: log}
}

func (p *PsqlDB) GetGuilds() []*entity.Guild {
	guilds := make([]*entity.Guild, 0)
	err := p.db.Select(&guilds, "SELECT * FROM guild")
	if err != nil {
		p.log.Error("error", zap.Error(err))
		return guilds
	}
	return guilds
}

func (p *PsqlDB) GetGuild(guildID string) (*entity.Guild, error) {
	guild := &entity.Guild{}
	err := p.db.Get(guild, "SELECT * FROM guild WHERE guild_id=$1 LIMIT 1", guildID)
	if err != nil {
		p.log.Error("error", zap.Error(err))
		return nil, err
	}
	/*
		// these are not wanted due to the potential size
		guild.Warns = p.GetGuildWarns(guildID)
		guild.Filters = p.GetGuildFilters(guildID)
	*/
	guild.MemberRoles = p.GetGuildMemberRoles(guildID)

	return guild, nil
}

func (p *PsqlDB) GetGuildMemberRoles(guildID string) []*entity.UserRole {
	roles := make([]*entity.UserRole, 0)
	err := p.db.Select(&roles, "SELECT * FROM user_role WHERE guild_id=$1", guildID)
	if err != nil {
		p.log.Error("error", zap.Error(err))
		return roles
	}
	return roles
}

func (p *PsqlDB) GetGuildFilters(guildID string) []*entity.Filter {
	roles := make([]*entity.Filter, 0)
	err := p.db.Select(&roles, "SELECT * FROM filter WHERE guild_id=$1", guildID)
	if err != nil {
		p.log.Error("error", zap.Error(err))
		return roles
	}
	return roles
}

func (p *PsqlDB) GetGuildWarns(guildID string) []*entity.Warn {
	roles := make([]*entity.Warn, 0)
	err := p.db.Select(&roles, "SELECT * FROM warn WHERE guild_id=$1", guildID)
	if err != nil {
		p.log.Error("error", zap.Error(err))
		return roles
	}
	return roles
}
