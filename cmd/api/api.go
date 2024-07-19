package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/d11m08y03/algox/service/ai"
	"github.com/d11m08y03/algox/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
  addr string
  db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
  return &APIServer{
    addr: addr,
    db: db,
  }
}

func (s *APIServer) Run() error {
  router := mux.NewRouter()

  log.Println("Listening on", s.addr)

  userStore := user.NewStore(s.db)
  userHandler := user.NewHandler(userStore)
  userHandler.RegisterRoutes(router)

  aiHanlder := ai.NewHandler()
  aiHanlder.RegisterRoutes(router)

  return http.ListenAndServe(s.addr, router)
}
