package userusecase

import (
	"errors"
	"log"
	"mmoz/crud/modules/user"
	"mmoz/crud/modules/user/userrepository"
	"mmoz/crud/utils"

	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecaseService interface {
		GetUserAllUsers() ([]*user.UserProfile, error)
		CreatePlayer(req *user.CreateUserReq) error
		GetUserByUsername(username string) (*user.UserProfile, error)
	}
	userUsecase struct {
		userRepository userrepository.UserRepositoryService
	}
)

func NewUserUsecase(userRepository userrepository.UserRepositoryService) UserUsecaseService {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetUserAllUsers() ([]*user.UserProfile, error) {
	ents, err := u.userRepository.GetUserAllUsers()
	if err != nil {
		return nil, err
	}

	var users []*user.UserProfile
	for _, ent := range ents {
		user := &user.UserProfile{
			Username: ent.Username,
			Role:     ent.Role,
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *userUsecase) CreatePlayer(req *user.CreateUserReq) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	refreshToken, err := utils.GenerateRefreshToken(req.Username, req.Role)
	if err != nil {
		log.Printf("Error generating refresh token: %v", err)
		return errors.New("Error generating refresh token")
	}

	req.RefreshToken = refreshToken

	req.Password = string(hashedPassword)

	err = u.userRepository.InsertPlayer(req)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) GetUserByUsername(username string) (*user.UserProfile, error) {
	ent, err := u.userRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return &user.UserProfile{
		Username: ent.Username,
		Role:     ent.Role,
	}, nil
}
