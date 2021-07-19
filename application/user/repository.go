package user

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/google/uuid"
)

type RepositoryUser interface {
	Login(user models.UserLogin) (*models.User, common.Err)
	CreateUser(newUser models.User) common.Err
	UpdateUser(newUser models.User) common.Err
	GetUserById(userId uuid.UUID) (*models.User, common.Err)
	GetUsers(userType string, start uint8, limit uint8) ([]models.User, common.Err)
	GetUsersAll(userType string) ([]models.User, common.Err)
}