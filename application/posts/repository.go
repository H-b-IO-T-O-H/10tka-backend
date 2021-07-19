package posts

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
)

type RepositoryPost interface {
	CreatePost(post models.Post) (*models.Post, common.Err)
	GetCurrentPostId() (int, common.Err)
	GetPostById(id int) (*models.Post, common.Err)
	GetPostsList(start uint16, limit uint16) ([]models.Post, common.Err)
	UpdatePost(post models.Post) (*models.Post, common.Err)
	DeletePost(id int) common.Err
}
