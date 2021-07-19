package repository

import (
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common/models"
	"github.com/H-b-IO-T-O-H/kts-backend/application/posts"
	"gorm.io/gorm"
	"net/http"
)

type pgStorage struct {
	db *gorm.DB
}

func (p pgStorage) CreatePost(post models.Post) (*models.Post, common.Err) {
	if err := p.db.Create(&post).Error; err != nil {
		return nil, common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	return &post, nil
}

func (p pgStorage) GetCurrentPostId() (int, common.Err) {
	var id int

	if err := p.db.Raw("select last_value from posts_id_seq;").Scan(&id).Error; err != nil {
		return 0, common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	return id, nil
}

func (p pgStorage) GetPostById(id int) (*models.Post, common.Err) {
	panic("implement me")
}

func (p pgStorage) GetPostsList(start int, limit int) ([]models.Post, common.Err) {
	panic("implement me")
}

func (p pgStorage) UpdatePost(post models.Post) (*models.Post, common.Err) {
	panic("implement me")
}

func (p pgStorage) DeletePost(id int) common.Err {
	panic("implement me")
}

func NewPgRepository(db *gorm.DB) posts.RepositoryPost {
	return &pgStorage{db: db}
}
