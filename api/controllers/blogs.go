package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogsController struct {
	blogServices services.BlogServices
	logger       infrastructure.Logger
}

func NewBlogController(
	blogServices services.BlogServices,
	logger infrastructure.Logger,
) BlogsController {
	return BlogsController{}
}

func (cc BlogsController) CreateBlog(c *gin.Context) {
	blog := models.Blog{}

	if err := c.ShouldBindJSON(&blog); err != nil {
		cc.logger.Zap.Error("Error [CreateBlog] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Faild to bind blog data")
		responses.HandleError(c, err)
		return
	}

	if err := cc.blogServices.CreateBlog(blog); err != nil {
		cc.logger.Zap.Error("Error [CreateBlog] [db CreateBlog] : ", err.Error())
		err := errors.InternalError.Wrap(err, " Failed to create blog")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Blog Created Successfully")
}

func (cc BlogsController) GetAllBlogs(c *gin.Context) {
	cursor := c.Param("cursor")
	blog, err := cc.blogServices.GetAllBlogs(cursor)

	if err != nil {
		cc.logger.Zap.Error("Error [GetAllBlog] [dB GetAllBlog] : ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Get All Blogs")
		responses.HandleError(c, err)
		return
	}
	responses.JSON(c, http.StatusOK, blog)
}

func (cc BlogsController) GetOneBlog(c *gin.Context) {
	blogId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	blog, err := cc.blogServices.GetOneBlog(blogId)

	if err != nil {
		cc.logger.Zap.Error("Error finding Blog : ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get Blog")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, blog)
}

func (cc BlogsController) UpdateBlogs(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	blog := models.Blog{}
	blog, err := cc.blogServices.GetOneBlog(id)

	if err != nil {
		cc.logger.Zap.Error("Error finding Blog", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get Blog")
		responses.HandleError(c, err)
		return
	}

	if err := c.ShouldBind(&blog); err != nil {
		cc.logger.Zap.Error("Error [UpdateBlog] (ShouldBindJson) : ", err)
		responses.ErrorJSON(c, http.StatusBadRequest, "Failed to update blog")
		return
	}

	blog.ID = id

	if err := cc.blogServices.UpdateBlogs(blog); err != nil {
		cc.logger.Zap.Error("Error updating blog", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to update Blog")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Updated")
}

func (cc BlogsController) DeleteBlog(c *gin.Context) {
	blogId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := cc.blogServices.DeleteBlogs(blogId)

	if err != nil {
		cc.logger.Zap.Error("Error Deleting Blog", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete Blog")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Deleted")
}
