package userusecase

import (
	"errors"
	"log"
	"mmoz/crud/modules"
	"mmoz/crud/modules/user"
	"mmoz/crud/modules/user/userrepository"
	"mmoz/crud/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecaseService interface {
		GetUserAllUsers() ([]*user.UserProfile, error)
		CreatePlayer(req *user.CreateUserReq) error
		GetUserByUsername(username string, token *modules.Token) (*user.UserProfile, error)
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
		log.Printf("Error getting user all users: %v", err)
		return nil, errors.New("Error getting user all users")
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
		log.Printf("Error hashing password: %v", err)
		return errors.New("Error hashing password")
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
		log.Printf("Error inserting player: %v", err)
		return errors.New("Error inserting player")
	}
	return nil
}

func (u *userUsecase) GetUserByUsername(username string, token *modules.Token) (*user.UserProfile, error) {

	if token.Role != "admin" && token.Username != username {
		return nil, errors.New("Cannot get other user's profile")
	}

	ent, err := u.userRepository.GetUserByUsername(username)
	if err != nil {
		log.Printf("Error getting user by username: %v", err)
		return nil, errors.New("Error getting user by username")
	}

	return &user.UserProfile{
		Username: ent.Username,
		Role:     ent.Role,
	}, nil
}
