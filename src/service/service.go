package service

import (
	"blogs_api/src/dto"
	"blogs_api/src/repository"
	"context"
	"fmt"
	"time"

	errorutils "blogs_api/utils/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var BlogRepository = new(repository.Blog)
var Err = errorutils.NewErr()

type BlogService struct{}

func (bs *BlogService) CreateBlog(c context.Context, req dto.CreateBlog) (dto.BlogHTTPResponse, *errorutils.Error) {
	var blog repository.Blog

	blog.Title = req.Title
	blog.Content = req.Content
	blog.Author = req.Author
	blog.CreatedAt = time.Now().UTC()
	blog.UpdatedAt = time.Now().UTC()

	BlogID, err := BlogRepository.CreateBlog(c, blog)
	if err != nil {
		fmt.Printf("Error creating Blog: %+v", err.Error())
		return dto.BlogHTTPResponse{}, &Err.INTERNAL_ERR
	}

	response := dto.BlogResponse{
		ID: BlogID.(primitive.ObjectID).Hex(),
	}

	return dto.BlogHTTPResponse{
		Status:  201,
		Message: "Blog Created successfully",
		Data:    response,
	}, nil
}

func (bs *BlogService) GetBlogByID(c context.Context, id string) (dto.GetBlogHTTPResponse, *errorutils.Error) {

	ctx, cancelFunc := context.WithCancel(c)
	defer cancelFunc()

	blog_id, _ := primitive.ObjectIDFromHex(id)

	blog, err := BlogRepository.FindBlogByID(ctx, blog_id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dto.GetBlogHTTPResponse{}, &Err.INVALID_ERR.BLOG_ID
		}
		fmt.Printf("Error fetching Blog: %+v", err.Error())
		return dto.GetBlogHTTPResponse{}, &Err.INTERNAL_ERR
	}

	response := dto.GetBlogResponse{
		ID:      blog.ID.Hex(),
		Title:   blog.Title,
		Content: blog.Content,
		Author:  blog.Author,
	}

	return dto.GetBlogHTTPResponse{
		Status:  200,
		Message: "The details of the blog for requested Id",
		Data:    response,
	}, nil
}

func (bs *BlogService) GetBlogList(c context.Context) (dto.BlogListHTTPResponse, *errorutils.Error) {

	ctx, cancelFunc := context.WithCancel(c)
	defer cancelFunc()

	var blog_resp []dto.GetBlogResponse

	blogs, err := BlogRepository.FindBlogList(ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dto.BlogListHTTPResponse{}, nil
		}
		fmt.Printf("Error fetching blog: %+v", err.Error())
		return dto.BlogListHTTPResponse{}, &Err.INTERNAL_ERR
	}

	for _, blogData := range blogs {
		response := dto.GetBlogResponse{
			ID:      blogData.ID.Hex(),
			Title:   blogData.Title,
			Content: blogData.Content,
			Author:  blogData.Author,
		}
		blog_resp = append(blog_resp, response)
	}

	if blogs != nil {

		return dto.BlogListHTTPResponse{
			Status:  200,
			Message: "Blog List success",
			Data:    blog_resp,
		}, nil

	} else {

		blog_resp = make([]dto.GetBlogResponse, 0)

		return dto.BlogListHTTPResponse{
			Status:  200,
			Message: "Blog List success",
			Data:    blog_resp,
		}, nil
	}
}
