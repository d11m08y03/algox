package user

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/d11m08y03/algox/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.IsNGO,
		&user.IsHospital,
		&user.CreatedAt,
		&user.BloodType,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	query := `SELECT * FROM users WHERE id=?`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) CreateUser(user types.RegisterUserPayload) error {
	statement, err := s.db.Prepare(`
    INSERT INTO users (firstName, lastName, email, password, isHospital, isNGO, bloodType)
    VALUES (?, ?, ?, ?, ?, ?, ?)
  `)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.IsHospital,
		user.IsNGO,
		user.BloodType,
	)
	if err != nil {
    log.Println("Error creating new user:" + err.Error())
		return err
	}

	log.Println("User created successfully")
	return nil
}
