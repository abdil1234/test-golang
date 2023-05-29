package repositories

import (
	"github.com/abdil1234/test-golang/internal/app/commons"
	"github.com/abdil1234/test-golang/internal/app/domain/repositories/models"
)

type IGameRepository interface {
	Insert(game models.Game) error
	Gets() (games []models.Game, err error)
}

type GameRepository struct {
	opt commons.Option
}

func NewGameRepository(opt commons.Option) IGameRepository {
	return &GameRepository{
		opt: opt,
	}
}

func (r *GameRepository) Insert(game models.Game) error {
	return r.opt.Database.Create(game).Error
}

func (r *GameRepository) Gets() (games []models.Game, err error) {
	err = r.opt.Database.Select("*").Find(&games).Error

	if err != nil {
		return
	}
	return
}
