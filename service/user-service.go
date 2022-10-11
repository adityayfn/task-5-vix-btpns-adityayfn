package service

import (
	"log"

	"github.com/adityayfn/task-5-vix-btpns-adityayfn/app"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/models"
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user models.UserUpdateModel) app.User
	Profile(userID string) app.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user models.UserUpdateModel) app.User {
	userToUpdate := app.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) app.User {
	return service.userRepository.ProfileUser(userID)
}
