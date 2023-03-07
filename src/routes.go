package routes

import (
	blogs "blogs_api/src/controller"

	"github.com/gin-gonic/gin"
)

var blogController = new(blogs.BlogController)

func BlogRoutes(r *gin.Engine) {

	blogs := r.Group("/articals")
	{
		blogs.POST("", blogController.CreateBlog)
	}
	{
		blogs.GET("/:id", blogController.GetBlogByID)
	}
	{
		blogs.GET("", blogController.GetBlogList)
	}
}
