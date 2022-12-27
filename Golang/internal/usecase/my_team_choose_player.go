package usecase

import (
	"context"

	"github.com/CassiusThalles/fullcycle11_cartolafc/tree/main/Golang/internal/domain/repository"
	"github.com/CassiusThalles/fullcycle11_cartolafc/tree/main/Golang/internal/domain/service"
	"github.com/CassiusThalles/fullcycle11_cartolafc/tree/main/Golang/pkg/uow"
)

type MyTeamChoosePlayerInput struct {
	ID string `json:"my_team_id"`
	PlayerID string `json:"player_id"`
}

type MyTeamChoosePlayerUseCase struct {
	Uow uow.UowInterface
}

func NewMyTeamChoosePlayerUseCase(uow uow.UowInterface) *MyTeamChoosePlayerUseCase {
	return &MyTeamChoosePlayerUseCase{
		Uow: uow
	}
}

func (u *MyTeamChoosePlayerUseCase) Execute(ctx context.Context, input MyTeamChoosePlayerInput) error {
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {
		playerRepo := u.getPlayerRepository(ctx)
		myTeamRepo := u.getMyTeamRepository(ctx)
		myTeam, err := myTeamRepo.FindByID(ctx, input.ID)
		if err != nil {
			return err
		}
		myPlayers, err := myTeamRepo.FindAllByIDs(ctx, myTeam.Players)
		if err != nil {
			return err
		}

		players, err := playerRepo.FindAllByIDs(ctx, input.PlayersID)
		if err != nil {
			return err
		}
		service.ChoosePlayers(myTeam, myPlayers, players)
		err = myTeamRepo.SavePlayers(ctx, myTeam)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (u *MyTeamChoosePlayerUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := u.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (u *MyTeamChoosePlayerUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := u.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}