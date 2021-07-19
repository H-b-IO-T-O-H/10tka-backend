package usecase

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/posts"
	"github.com/go-openapi/strfmt"
	"time"
)

type PostUseCase struct {
	repos posts.RepositoryPost
}

func NewPostUseCase(repos posts.RepositoryPost) posts.UseCase {
	return PostUseCase{
		repos: repos,
	}
}

func (p PostUseCase) CreatePost(post models.Post) (*models.Post, common.Err) {
	post.Created = strfmt.DateTime(time.Now())
	return p.repos.CreatePost(post)
}

func (p PostUseCase) GetCurrentPostId() (int, common.Err)  {
	return p.repos.GetCurrentPostId()
}

func (p PostUseCase) GetPostById(id int) (*models.Post, common.Err) {
	return p.repos.GetPostById(id)
}

func (p PostUseCase) GetPostsList(start int, limit int) ([]models.Post, common.Err) {
	return p.repos.GetPostsList(start, limit)
}

func (p PostUseCase) UpdatePost(post models.Post) (*models.Post, common.Err) {
	return p.repos.UpdatePost(post)
}

func (p PostUseCase) DeletePost(id int) common.Err {
	return p.repos.DeletePost(id)
}
