package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt"
	"questionGame/entity"
	"questionGame/pkg/phonenumber"
	"time"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(PhoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type Service struct {
	signKey string
	repo    Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User
}

func New(repo Repository, signKey string) Service {
	return Service{repo: repo, signKey: signKey}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO: we should verify phone number by verification code

	// validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("invalid phone number")
	}

	// check uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error while checking if phone number exists: %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number %s is already used", req.PhoneNumber)
		}
	}

	// validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name should be greater than 3 characters")
	}

	// TODO: check password with regex pattern
	// validate password
	if len(req.PhoneNumber) < 8 {
		return RegisterResponse{}, fmt.Errorf("phone_number should be greater than 8 characters")
	}

	// TODO: replace md5 with bcrypt

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    getMD5Hash(req.Password),
	}

	// create new user in storage
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error while registering user: %w", err)
	}

	// return created user
	return RegisterResponse{
		User: createdUser,
	}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO: it would be better to use two separate methods for existence check and getUserByPhoneNumber

	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error while getting user by phone number: %w", err)
	}

	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password is wrong")
	}

	if user.Password != getMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password is wrong")
	}

	// jwt token
	token, err := createToken(user.ID, s.signKey)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error while creating token: %w", err)
	}

	return LoginResponse{AccessToken: token}, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}

// all request inputs for interactor/service could be sanitized

func (s Service) GetProfile(req ProfileRequest) (ProfileResponse, error) {
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// I have not expect the repository call return "record not found" error,
		// because I assume the interactor input is sanitized
		// TODO: we can use Rich Error
		return ProfileResponse{}, fmt.Errorf("unexpected error while getting user by ID: %w", err)
	}

	// return data
	return ProfileResponse{Name: user.Name}, nil
}

type Claims struct {
	StandardClaims jwt.StandardClaims
	UserID         uint
}

func (c Claims) Valid() error {
	return nil
}

func createToken(userID uint, signKey string) (string, error) {
	// TODO: replace with rsa 256
	// create a signer for rsa 256
	//t := jwt.New(jwt.GetSigningMethod("HS256"))

	// set claim
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
			ExpiresAt: int64(time.Hour * 24 * 7),
		},
		UserID: userID,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(signKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
