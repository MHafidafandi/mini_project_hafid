package usecase

import (
	"miniproject/constant"
	"miniproject/helper"
	"miniproject/internal/dto/request"
	"miniproject/internal/models"
	"miniproject/internal/repository"
	"time"

	"github.com/google/uuid"
)

type UserUsecase interface {
	RegisterUser(userDTO request.UserRequest) error
	LoginUser(email string, password string) (string, error)
	FindUserById(userId string) (*models.User, error)
	UpdateUser(id string, userDTO request.UserUpdate) error
	DeleteUser(userId string) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(
	userRepo repository.UserRepository,

) UserUsecase {
	return userUsecase{
		userRepository: userRepo,
	}
}

func (u userUsecase) RegisterUser(userDTO request.UserRequest) error {
	user, _ := u.userRepository.FindByEmail(userDTO.Email)

	if user != nil {
		return constant.ErrDataAlreadyExist
	}

	hashedPassword, _ := helper.HashPassword(userDTO.Password)

	userUC := models.User{
		ID:        uuid.NewString(),
		Name:      userDTO.Name,
		Phone:     userDTO.Phone,
		Address:   userDTO.Address,
		Role:      userDTO.Role,
		Email:     userDTO.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.userRepository.Create(userUC)

	if err != nil {
		return err
	}

	return nil
}

func (u userUsecase) LoginUser(email string, password string) (string, error) {

	user, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return "", err
	}

	ok := helper.ComparePassword(user.Password, password)

	if !ok {
		return "", constant.ErrRecordNotFound
	}

	token, _ := helper.CreateToken(user.ID, user.Role)

	return token, nil
}

func (u userUsecase) FindUserById(userId string) (*models.User, error) {
	user, err := u.userRepository.FindById(userId)

	if err != nil {
		return nil, constant.ErrRecordNotFound
	}

	return user, nil
}

func (u userUsecase) UpdateUser(id string, userDTO request.UserUpdate) error {
	var err error

	_, err = u.FindUserById(id)

	if err != nil {
		return err
	}

	userUC := models.User{
		Name:      userDTO.Name,
		Phone:     userDTO.Phone,
		Address:   userDTO.Address,
		Role:      userDTO.Role,
		UpdatedAt: time.Now(),
	}

	err = u.userRepository.Update(id, userUC)

	if err != nil {
		return err
	}

	return nil

}
func (u userUsecase) DeleteUser(userId string) error {
	var err error

	_, err = u.FindUserById(userId)

	if err != nil {
		return err
	}
	err = u.userRepository.Delete(userId)

	if err != nil {
		return err
	}

	return nil

}
