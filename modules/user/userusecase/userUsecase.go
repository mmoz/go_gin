package userusecase

import (
	"errors"
	"log"
	"mmoz/crud/modules"
	"mmoz/crud/modules/user"
	"mmoz/crud/modules/user/userrepository"
	"mmoz/crud/utils"

	"github.com/google/uuid"
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

	isUnique, err := u.userRepository.IsUniquePlayer(req.Username)
	if err != nil {
		log.Printf("Error checking unique username: %v", err)
		return errors.New("Error checking unique username")
	}

	if !isUnique {
		return errors.New("Username already exists")
	}

	newUUID := uuid.New()

	req.ID = newUUID.String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return errors.New("Error hashing password")
	}

	refreshToken, err := utils.GenerateRefreshToken(req.ID, req.Username, req.Role)
	if err != nil {
		log.Printf("Error generating refresh token: %v", err)
		return errors.New("Error generating refresh token")
	}

	req.RefreshToken = refreshToken

	req.Password = string(hashedPassword)

	user := &user.UserProfileEnt{
		ID:            req.ID,
		Username:      req.Username,
		Password:      req.Password,
		Role:          req.Role,
		RefreshToken:  req.RefreshToken,
		IsTokenActive: 1,
	}

	err = u.userRepository.InsertPlayer(user)
	if err != nil {
		log.Printf("Error inserting player: %v", err)
		return errors.New("Error inserting player")
	}
	return nil
}

func (u *userUsecase) GetUserByUsername(username string, token *modules.Token) (*user.UserProfile, error) {

	ent, err := u.userRepository.GetUserByUsername(username)
	if err != nil {
		log.Printf("Error getting user by username: %v", err)
		return nil, errors.New("Error getting user by username")
	}

	users := new(user.UserProfile)
	if token.Role == "admin" {
		users = &user.UserProfile{
			Username: ent.Username,
			Role:     ent.Role,
		}
		return users, nil
	}

	return &user.UserProfile{
		Username: ent.Username,
	}, nil
}
