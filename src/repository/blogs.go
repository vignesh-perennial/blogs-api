package repository

import (
	"blogs_api/utils/database"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var blogRepo = database.Db().Database("blog_db").Collection("blogs")

type Blog struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title,omitempty"`
	Content   string             `json:"content" bson:"content,omitempty"`
	Author    string             `json:"author" bson:"author,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	IsDeleted bool               `json:"is_deleted" bson:"is_deleted"`
}

func (b *Blog) CreateBlog(ctx context.Context, blog Blog) (interface{}, error) {

	r, err := blogRepo.InsertOne(ctx, blog)
	if err != nil {
		return "", err
	}

	return r.InsertedID, nil
}

func (b *Blog) FindBlogByID(ctx context.Context, id primitive.ObjectID) (Blog, error) {

	var blog Blog

	err := blogRepo.FindOne(ctx, bson.M{"_id": id}).Decode((&blog))
	if err != nil {
		return blog, err
	}

	return blog, nil
}

func (b *Blog) FindBlogByBson(ctx context.Context, filter bson.D) (Blog, error) {

	var blog Blog

	err := blogRepo.FindOne(ctx, filter).Decode((&blog))
	if err != nil {
		return blog, err
	}

	return blog, nil
}

func (b *Blog) FindBlogList(ctx context.Context) ([]Blog, error) {

	var blog []Blog

	options := options.Find()

	cur, err := blogRepo.Find(ctx, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var blogs Blog

		err := cur.Decode(&blogs)
		if err != nil {
			return nil, err
		}

		blog = append(blog, blogs)

	}

	return blog, nil
}
