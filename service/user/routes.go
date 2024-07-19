package user

import (
	"fmt"
	"net/http"

	"github.com/d11m08y03/algox/service/auth"
	"github.com/d11m08y03/algox/types"
	"github.com/d11m08y03/algox/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
  var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

  u, err := h.store.GetUserByEmail(payload.Email)
  if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wrong email or password"))
		return
  }

  if u.Password != payload.Password {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wrong email or password"))
		return
  }

  secret := []byte("YepSecret")
  token, err := auth.CreateJWT(secret, u.ID)
  if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
  }

  utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

  h.store.CreateUser(payload)
	utils.WriteJSON(w, http.StatusCreated, "User Created")
}
