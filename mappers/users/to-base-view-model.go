package mappers

import (
	"github.com/wfl-junior/go-car-rental-api/models"
	UserViewModels "github.com/wfl-junior/go-car-rental-api/view-models/users"
)

func ToBaseViewModel(user models.User) UserViewModels.Base {
	return UserViewModels.Base{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
	}
}
