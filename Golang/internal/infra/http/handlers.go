package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	irepository "mygolangapp/internal/domain/repository"
	"mygolangapp/internal/infra/db"
	"mygolangapp/internal/infra/presenter"
	"github.com/go-chi/chi/v5"
)

func ListPlayersHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players, err := queries.FindAllPlayers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(players)
	}
}

func ListMyTeamPlayersHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamID := chi.URLParam(r, "teamID")
		players, err := queries.GetPlayersByMyTeamID(ctx, teamID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(players)
	}
}

func ListMatchesHandler(ctx context.Context, matchRepository irepository.MatchRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		matches, err := matchRepository.FindAll(ctx)
		var matchesPresenter presenter.Matches
		for _, match := range matches {
			matchesPresenter = append(matchesPresenter, presenter.NewMatchPresenter(match))
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(matchesPresenter)
	}
}

func ListMatchByIDHandler(ctx context.Context, matchRepository irepository.MatchRepositoryInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchID := chi.URLParam(r, "matchID")
		match, err := matchRepository.FindByID(ctx, matchID)
		matchPresenter := presenter.NewMatchPresenter(match)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(matchPresenter)
	}
}

func GetMyTeamBalanceHandler(ctx context.Context, queries db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamID := chi.URLParam(r, "teamID")
		fmt.Println(teamID)
		balance, err := queries.GetMyTeamBalance(ctx, teamID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resultJson := map[string]float64{"balance": balance}
		json.NewEncoder(w).Encode(resultJson)
	}
}