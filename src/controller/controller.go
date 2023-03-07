package controller

import (
	"blogs_api/src/dto"
	"blogs_api/src/service"
	"blogs_api/src/validator"
	errorsutils "blogs_api/utils/errors"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var blogService = new(service.BlogService)
var Err = errorsutils.NewErr()

type BlogController struct{}

func (blogController *BlogController) CreateBlog(ctx *gin.Context) {

	var r map[string]interface{}

	var req dto.CreateBlog

	if err := ctx.BindJSON(&r); err != nil {
		errorsutils.ErrorHandler(ctx, &Err.INTERNAL_ERR)
		return
	}

	j, er := json.Marshal(r)
	if er != nil {
		fmt.Printf(er.Error())
	}
	json.Unmarshal(j, &req)

	errs := validator.CreateBLogValidator(ctx, req)
	if errs != nil {
		errorsutils.ErrorHandler(ctx, errs)
		return
	}

	response, e := blogService.CreateBlog(ctx, req)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (blogController *BlogController) GetBlogByID(ctx *gin.Context) {

	//artical id
	id := ctx.Param("id")

	errs := validator.GetBlogValidator(ctx, id)
	if errs != nil {
		errorsutils.ErrorHandler(ctx, errs)
		return
	}

	response, e := blogService.GetBlogByID(ctx, id)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (blogController *BlogController) GetBlogList(ctx *gin.Context) {

	response, e := blogService.GetBlogList(ctx)
	if e != nil {
		errorsutils.ErrorHandler(ctx, e)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
