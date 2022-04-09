package entity

import "github.com/google/uuid"

type Authentication struct {
	UserID string `db:"user_id" json:"userID"`
}

type AuthService interface {
	Get(userID string)
}

type AuthRepo interface {
	Get(userID string)
}

type AquariumService interface {
	Insert(aq *Aquarium) (*Aquarium, error)
	GetByID(userID string) (*Aquarium, error)
	Update(aq *Aquarium) (*Aquarium, error)
	//DeleteByID(uid uuid.UUID) error
}
type AquariumRepository interface {
	Insert(aq *Aquarium) (*Aquarium, error)
	GetByID(userID string) (*Aquarium, error)
	Update(aq *Aquarium) (*Aquarium, error)
	//DeleteByID(uid uuid.UUID) error
}

type CommandLogService interface {
	Insert(cl *CommandLog) (*CommandLog, error)
}
type CommandLogRepository interface {
	Insert(cl *CommandLog) (*CommandLog, error)
}

type FilterService interface {
	Insert(f *Filter) (*Filter, error)
	GetByID(uid uuid.UUID) (*Filter, error)
	Delete(uid uuid.UUID) error
}
type FilterRepository interface {
	Insert(f *Filter) (*Filter, error)
	GetByID(uid uuid.UUID) (*Filter, error)
	Delete(uid uuid.UUID) error
}

type UserRoleService interface {
	Insert(u *UserRole) (*UserRole, error)
	GetByIDS(GuildID, userID string) (*UserRole, error)
	Update(u *UserRole) (*UserRole, error)
	Delete(uid uuid.UUID) error
}
type UserRoleRepository interface {
	Insert(u *UserRole) (*UserRole, error)
	GetByIDS(GuildID, userID string) (*UserRole, error)
	Update(u *UserRole) (*UserRole, error)
	Delete(uid uuid.UUID) error
}

type WarnService interface {
}
type WarnRepository interface {
}
