package validator

import (
	"blogs_api/src/dto"
	"blogs_api/src/repository"
	errorutils "blogs_api/utils/errors"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Err = errorutils.NewErr()

func CreateBLogValidator(ctx context.Context, req dto.CreateBlog) interface{} {

	var invalidErrs []*errorutils.Error

	if len(req.Title) == 0 || len(req.Title) > 255 {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.TITLE)
		// return invalidErrs
	}

	if len(req.Author) == 0 || len(req.Author) > 100 {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.AUTHOR)
		// return invalidErrs
	}
	if len(req.Content) == 0 || len(req.Content) > 1000 {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.CONTENT)
		// return invalidErrs
	}

	if len(invalidErrs) != 0 {
		return invalidErrs
	}

	return nil
}

func GetBlogValidator(ctx context.Context, id string) interface{} {

	var invalidErrs []*errorutils.Error

	var BlogRepo = new(repository.Blog)

	blog_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.BLOG_ID)
		return invalidErrs
	}

	var filter bson.D
	filter = append(filter, primitive.E{Key: "_id", Value: blog_id})
	filter = append(filter, primitive.E{Key: "is_deleted", Value: false})

	blog, err := BlogRepo.FindBlogByBson(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			invalidErrs = append(invalidErrs, &Err.INVALID_ERR.BLOG_ID)
			return invalidErrs
		}

		return &Err.INTERNAL_ERR
	}

	if blog.ID == primitive.NilObjectID {
		invalidErrs = append(invalidErrs, &Err.INVALID_ERR.BLOG_ID)
		return invalidErrs
	}

	if len(invalidErrs) != 0 {
		return invalidErrs
	}

	return nil
}
