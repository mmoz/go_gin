package authrepository

import (
	"database/sql"
	"log"
	"mmoz/crud/modules/auth"
)

type (
	AuthRepositoryService interface {
		CheckCredential(cres *auth.CredentialReq) (*auth.Credential, string, error)
	}

	authrepository struct {
		db *sql.DB
	}
)

func NewAuthRepository(db *sql.DB) AuthRepositoryService {
	return &authrepository{db: db}
}

func (r *authrepository) CheckCredential(cres *auth.CredentialReq) (*auth.Credential, string, error) {

	stmt, err := r.db.Prepare("SELECT username,password,roles,refreshtoken FROM users WHERE username = ?")
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return nil, "", err
	}

	users := new(auth.CredentialRes)
	var password string

	err = stmt.QueryRow(cres.Username).Scan(&users.Username, &password, &users.Role, &users.RefreshToken)
	if err != nil {
		log.Printf("Error scanning row: %v", err)
		return nil, "", err
	}
	

	return &auth.Credential{
		Username:     users.Username,
		Role:         users.Role,
		RefreshToken: users.RefreshToken,
	}, password, nil

}
