package authusecase

import (
	"errors"
	"log"
	"mmoz/crud/modules/auth"
	"mmoz/crud/modules/auth/authrepository"
	"mmoz/crud/utils"

	"golang.org/x/crypto/bcrypt"
)

type (
	AuthUsecaseService interface {
		CheckLogin(cres *auth.CredentialReq) (*auth.CredentialRes, error)
	}

	authUsecase struct {
		authrepository authrepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authrepository authrepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{
		authrepository: authrepository,
	}
}

func (u *authUsecase) CheckLogin(cres *auth.CredentialReq) (*auth.CredentialRes, error) {

	result, password, err := u.authrepository.CheckCredential(cres)
	if err != nil {
		return nil, errors.New("failed to login")
	}
	//compare password
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(cres.Password)); err != nil {
		log.Printf("Error comparing password: %v", err)
		return nil, errors.New("Invalid username and password")
	}

	accessToken, err := utils.GenerateAccessToken(result.Username, result.Role)
	if err != nil {
		log.Printf("Error generating access token: %v", err)
		return nil, errors.New("Error generating access token")
	}

	return &auth.CredentialRes{
		Username:     result.Username,
		Role:         result.Role,
		RefreshToken: result.RefreshToken,
		AccessToken:  accessToken,
	}, nil
}
