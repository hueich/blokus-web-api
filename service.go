package blokusService

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type BlokusService struct {
	router *mux.Router
}

func New(prefix string) *BlokusService {
	s := &BlokusService{}
	s.initRouter(prefix)
	return s
}

func (s *BlokusService) initRouter(prefix string) {
	r := mux.NewRouter()
	sr := r.PathPrefix(path.Join(prefix, "/games")).Subrouter()

	// Gets a webpage with a listing of games.
	sr.HandleFunc("", s.getGamesHandler).Methods("GET")
	// Creates a game.
	sr.HandleFunc("", s.newGameHandler).Methods("POST")

	g := sr.PathPrefix("/{gid:[0-9]+}").Subrouter()
	// Gets a webpage showing the specified game.
	g.HandleFunc("", s.getGameHandler).Methods("GET")
	// Gets various game state data.
	g.HandleFunc("/state", s.getGameStateHandler).Methods("GET")
	// Add a player.
	g.HandleFunc("/players", s.newPlayerHandler).Methods("POST")
	// Make a move in the game.
	g.HandleFunc("/moves", s.newMoveHandler).Methods("POST")

	s.router = r
}

func (s *BlokusService) Router() http.Handler {
	return s.router
}
