package http

import (
	"context"
	"net/http"

	"github.com/CassiusThalles/fullcycle11_cartolafc/tree/main/Golang/internal/infra/db"
)

func ListPlayersHandler(ctx context.Context, queris db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players, err := queries.FindAllPlayers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(players)
	}
}