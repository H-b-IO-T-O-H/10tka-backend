package repository

import (
	"fmt"
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
	var post models.Post

	if err := p.db.First(&post, id).Error; err != nil {
		msg := err.Error()
		if common.RecordNotFound(msg) {
			return nil, common.RespErr{Status: http.StatusNotFound, Message: common.NotFound}
		}
		return nil, common.RespErr{Status: http.StatusInternalServerError, Message: msg}
	}

	return &post, nil
}

func (p pgStorage) GetPostsList(start uint16, limit uint16) ([]models.Post, common.Err) {
	var postList []models.Post

	query := "select * from public.posts"
	if limit != 0 {
		query = fmt.Sprintf("%s offset(%d) limit(%d)", query, start, limit)
	}
	if err := p.db.Raw(query).Scan(&postList).Error; err != nil {
		return nil, common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	return postList, nil
}

func (p pgStorage) UpdatePost(post models.Post) (*models.Post, common.Err) {
	panic("implement me")
}

func (p pgStorage) DeletePost(id int) common.Err {
	var isExist bool

	if err := p.db.Raw("select exists(select 1 from public.posts where id = ?)", id).Scan(&isExist).Error; err != nil {
		return common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	if !isExist {
		return common.RespErr{Status: http.StatusNotFound, Message: common.NotFound}
	}
	if err := p.db.Table("public.posts").Delete(models.Post{PostId: id}).Error; err != nil {
		return common.RespErr{Status: http.StatusInternalServerError, Message: err.Error()}
	}
	return nil
}

func NewPgRepository(db *gorm.DB) posts.RepositoryPost {
	return &pgStorage{db: db}
}
