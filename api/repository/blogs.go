package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"time"
)

type BlogsRepository struct {
	db infrastructure.Database
}

func NewBlogsRepository(db infrastructure.Database) BlogsRepository {
	return BlogsRepository{
		db: db,
	}
}

func (c BlogsRepository) CreateBlog(Blogs models.Blog) error {
	return c.db.DB.Create(&Blogs).Error
}

func (c BlogsRepository) GetAllBlogs(cursor string) ([]models.Blog, error) {
	var blogs []models.Blog
	queryBuilder := c.db.DB.Model(&models.Blog{}).Order("created_at desc").Find(&blogs).Limit(20)
	if cursor != "" {
		time, _ := time.Parse(time.RFC3339, cursor)
		queryBuilder = queryBuilder.Where("created_at < ?", time)
	}
	return blogs, queryBuilder.Error
}

func (c BlogsRepository) GetOneBlog(blogId int64) (Blog models.Blog, err error) {
	return Blog, c.db.DB.Model(&models.Blog{}).Where("id = ?", blogId).First(&Blog).Error
}

func (c BlogsRepository) UpdateBlogs(Blog models.Blog) error {
	return c.db.DB.Model(&models.Blog{}).
		Where("id = ?", Blog.ID).
		Where("deleted_at IS NOT NULL").
		Updates(&Blog).Error
}

func (c BlogsRepository) DeleteBlog(blogId int64) error {
	return c.db.DB.Where("id = ?", blogId).Delete(&models.Blog{}).Error
}
