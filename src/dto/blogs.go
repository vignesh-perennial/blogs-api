package dto

type CreateBlog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type BlogHTTPResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    BlogResponse `json:"data"`
}

type BlogResponse struct {
	ID string `json:"id"`
}

type GetBlogHTTPResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    GetBlogResponse `json:"data"`
}

type BlogListHTTPResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []GetBlogResponse `json:"data"`
}

type GetBlogResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
