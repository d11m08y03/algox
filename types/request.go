package types

import (
	"time"
)

type BloodRequestStore interface {
  CreateRequest(req BloodRequestPayload) error
  GetPendingRequests() ([]BloodRequest, error)
}

type BloodRequest struct {
	ID          int       `json:"id"`
	BloodType   string    `json:"bloodType"`
	RequesterID string    `json:"requesterID"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
}

type BloodRequestPayload struct {
	BloodType   string    `json:"bloodType"`
	RequesterID string    `json:"requesterID"`
}
