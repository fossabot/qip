package repository

import (
	"github.com/laminne/notepod/pkg/models/domain"
	"github.com/laminne/notepod/pkg/utils/id"
)

type UserRepository interface {
	Create(u domain.User) (domain.User, error)
	Update(u domain.User) (domain.User, error)
	FindByID(id id.SnowFlakeID) (*domain.User, error)
	FindByUserName(n string) (*domain.User, error)
	FindByHost(h string) ([]domain.User, error)
}
