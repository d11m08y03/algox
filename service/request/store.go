package request

import (
	"database/sql"
	"log"

	"github.com/d11m08y03/algox/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateRequest(req types.BloodRequestPayload) error {
  statement, err := s.db.Prepare(`
    INSERT INTO requests (bloodType, requesterID, completed)
    VALUES (?, ?, ?)
  `)
  if err != nil {
    return err
  }
  defer statement.Close()

  _, err = statement.Exec(
    req.BloodType,
    req.RequesterID,
    false,
  )
  if err != nil {
    log.Println("Error creating new request")
    return err
  }

  log.Println("Request created successfully")
  return nil
}

func scanRowIntoBloodRequest(rows *sql.Rows) (types.BloodRequest, error) {
  var bloodRequest types.BloodRequest

  err := rows.Scan(
    &bloodRequest.ID,
    &bloodRequest.BloodType,
    &bloodRequest.RequesterID,
    &bloodRequest.Completed,
    &bloodRequest.CreatedAt,
  )
  if err != nil {
    return bloodRequest, err
  }

  return bloodRequest, nil
}

func (s *Store) GetPendingRequests() ([]types.BloodRequest, error) {
  query := `SELECT * FROM requests WHERE completed=false`

  rows, err := s.db.Query(query)
  if err != nil {
    return nil, err
  }

  var pendingRequests []types.BloodRequest
  for rows.Next() {
    r, err := scanRowIntoBloodRequest(rows)
    if err == nil {
      pendingRequests = append(pendingRequests, r)
    }
  }

	return pendingRequests, nil
}
