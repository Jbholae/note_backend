package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type BlogServices struct {
	repository repository.BlogsRepository
}

func NewBlogService(
	repository repository.BlogsRepository,
) BlogServices {
	return BlogServices{
		repository: repository,
	}
}

func (c BlogServices) CreateBlog(Blogs models.Blog) error {
	return c.repository.CreateBlog(Blogs)
}

func (c BlogServices) GetAllBlogs(cursor string) ([]models.Blog, error) {
	return c.repository.GetAllBlogs(cursor)
}

func (c BlogServices) GetOneBlog(blogId int64) (Blog models.Blog, err error) {
	return c.repository.GetOneBlog(blogId)
}

func (c BlogServices) UpdateBlogs(Blog models.Blog) error {
	return c.repository.UpdateBlogs(Blog)
}

func(c BlogServices) DeleteBlog(blogId int64)error{
	return c.repository.DeleteBlog(blogId)
}
