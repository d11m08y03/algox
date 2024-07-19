package request

import (
	"encoding/json"
	"net/http"

	"github.com/d11m08y03/algox/types"
	"github.com/d11m08y03/algox/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.BloodRequestStore
}

func NewHandler(store types.BloodRequestStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
  router.HandleFunc("/registerRequest", h.handleRegister).Methods("POST")
  router.HandleFunc("/getPendingRequests", h.handleGetPendingRequests).Methods("GET")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
  var payload types.BloodRequestPayload
  if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
  }

  h.store.CreateRequest(payload)
}

func (h *Handler) handleGetPendingRequests(w http.ResponseWriter, r *http.Request) {
  pendingRequest, err := h.store.GetPendingRequests()
  if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
    return
  }

  jsonData, err := json.Marshal(pendingRequest)
  if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
    return
  }

  utils.WriteJSON(w, http.StatusOK, jsonData)
}
