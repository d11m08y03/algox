package ai

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/d11m08y03/algox/types"
	"github.com/d11m08y03/algox/utils"
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/getBloodDemand", h.handleGetBloodDemand).Methods("POST")
}

func (h *Handler) handleGetBloodDemand(w http.ResponseWriter, r *http.Request) {
  var payload types.BloodDemandRequestPayload

  if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
  }

  jsonData, err := json.Marshal(payload)
  if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
  }

  url := ""
  response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
  if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
  }
  defer response.Body.Close()

  var responseData types.BloodDemandResultPayload
  if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
  }

  utils.WriteJSON(w, http.StatusOK, responseData)
}
