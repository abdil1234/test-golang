package game

import (
	"github.com/abdil1234/test-golang/internal/app/commons"
	"github.com/abdil1234/test-golang/internal/app/domain/entities"
	"github.com/abdil1234/test-golang/internal/app/domain/repositories"
	"github.com/abdil1234/test-golang/internal/app/domain/repositories/models"
)

type IGameService interface {
	Insert(game models.Game) error
	Gets() []entities.GameEntity
}

type Service struct {
	opt  commons.Option
	repo *repositories.Repository
}

func NewGameService(opt commons.Option, repo *repositories.Repository) IGameService {
	return &Service{
		opt:  opt,
		repo: repo,
	}
}

func (svc *Service) Insert(game models.Game) error {
	return svc.repo.Game.Insert(game)
}

func (svc *Service) Gets() (games []entities.GameEntity) {
	models, _ := svc.repo.Game.Gets()

	for _, game := range models {
		model := entities.GameModel{Game: game}
		gameEntity := *model.ToGameEntity()
		games = append(games, gameEntity)
	}
	return games
}
