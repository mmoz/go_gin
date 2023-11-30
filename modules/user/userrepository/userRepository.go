package userrepository

import (
	"database/sql"
	"errors"
	"log"
	"mmoz/crud/modules/user"
)

type (
	UserRepositoryService interface {
		GetUserAllUsers() ([]*user.UserProfileEnt, error)
		InsertPlayer(req *user.CreateUserReq) error
		GetUserByUsername(username string) (*user.UserProfileEnt, error)
	}

	userRepository struct {
		db *sql.DB
	}
)

func NewUserRepository(db *sql.DB) UserRepositoryService {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserAllUsers() ([]*user.UserProfileEnt, error) {
	stmt, err := r.db.Prepare("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*user.UserProfileEnt{}

	for rows.Next() {
		user := new(user.UserProfileEnt)
		err := rows.Scan(&user.Username, &user.Password, &user.Role, &user.RefreshToken, &user.IsTokenActive)
		if err != nil {
			log.Printf("Error: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) InsertPlayer(req *user.CreateUserReq) error {

	stmt, err := r.db.Prepare("INSERT INTO users (username, password, roles,refreshtoken,istokenactive) VALUES (?, ?, ?, ? ,?)")
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	_, err = stmt.Exec(req.Username, req.Password, req.Role, req.RefreshToken, 1)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return nil
}

func (r *userRepository) GetUserByUsername(username string) (*user.UserProfileEnt, error) {
	stmt, err := r.db.Prepare("SELECT * FROM users WHERE username = ?")
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	row := stmt.QueryRow(username)

	user := new(user.UserProfileEnt)
	err = row.Scan(&user.Username, &user.Password, &user.Role, &user.RefreshToken, &user.IsTokenActive)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, errors.New("Failed: User not found")
	}

	return user, nil
}
