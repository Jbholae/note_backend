package controllers

import (
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/responses"
	"net/http"

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
