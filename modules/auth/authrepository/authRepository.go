package authrepository

import (
	"database/sql"
	"log"
	"mmoz/crud/modules/auth"
)

type (
	AuthRepositoryService interface {
		CheckCredential(cres *auth.Credential) (*auth.Credential, error)
	}

	authrepository struct {
		db *sql.DB
	}
)

func NewAuthRepository(db *sql.DB) AuthRepositoryService {
	return &authrepository{db: db}
}

func (r *authrepository) CheckCredential(cres *auth.Credential) (*auth.Credential, error) {

	stmt, err := r.db.Prepare("SELECT * FROM users WHERE username = ? and istokenactive = 1")
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	users := new(auth.Credential)

	err = stmt.QueryRow(cres.Username).Scan(&users.ID, &users.Username, &users.Password, &users.Role, &users.RefreshToken, &users.IsTokenActive)
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, err
	}

	return &auth.Credential{
		ID:            users.ID,
		Username:      users.Username,
		Password:      users.Password,
		Role:          users.Role,
		RefreshToken:  users.RefreshToken,
		IsTokenActive: users.IsTokenActive,
	}, nil

}
