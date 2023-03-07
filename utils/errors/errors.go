package errors

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Err struct {
	INTERNAL_ERR  Error
	INVALID_ERR   Invalid_Error
	OBJ_NOT_FOUND Error
}

type Invalid_Error struct {
	BLOG_ID Error
	TITLE   Error
	AUTHOR  Error
	CONTENT Error
}

type InvalidErrorResponse struct {
	Code        int             `json:"code"`
	Message     string          `json:"message"`
	Description string          `json:"description"`
	Errors      []InvalidErrors `json:"errors"`
}

type CommonErrorResponse struct {
	Code        int                      `json:"code"`
	Message     string                   `json:"message"`
	Description string                   `json:"description"`
	Errors      []map[string]interface{} `json:"errors"`
}

type InvalidErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Code    int
	Type    string
	Field   string
	Message string
}

func ErrorHandler(ctx *gin.Context, err interface{}) {

	if e, ok := err.(*Error); ok {
		errResponse := CommonErrorResponse{
			Code:        e.Code,
			Message:     e.Type,
			Description: e.Message,
			Errors:      make([]map[string]interface{}, 0),
		}

		ctx.JSON(errResponse.Code, errResponse)
		var m map[string]interface{}
		json.NewDecoder(ctx.Request.Body).Decode(&m)
		return

	} else if er, ok := err.([]*Error); ok {
		errResponse := InvalidErrorResponse{}
		errResponse.Code = 400
		errResponse.Message = "invalid_request_error"
		errResponse.Description = "The request was unacceptable, due to missing a required parameter or invalid parameter."
		var invalid_fields []InvalidErrors

		for _, e := range er {
			invalid := InvalidErrors{
				Field:   e.Field,
				Message: e.Message,
			}
			invalid_fields = append(invalid_fields, invalid)
		}
		errResponse.Errors = append(errResponse.Errors, invalid_fields...)
		ctx.JSON(errResponse.Code, errResponse)
		var m map[string]interface{}
		json.NewDecoder(ctx.Request.Body).Decode(&m)

		return
	}

	errResponse := CommonErrorResponse{}
	errResponse.Code = 500
	errResponse.Message = "internal_server_error"
	errResponse.Description = "Internal server error."

	ctx.JSON(errResponse.Code, errResponse)
	var m map[string]interface{}
	json.NewDecoder(ctx.Request.Body).Decode(&m)

}

func NewErr() Err {
	e := Err{
		INTERNAL_ERR: Error{
			Code:    500,
			Type:    "internal_server_error",
			Message: "Internal server error",
		},

		INVALID_ERR: Invalid_Error{

			BLOG_ID: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "id",
				Message: "This field should be a valid blog id",
			},
			TITLE: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "title",
				Message: "This field should be valid string not greater than 255 characters",
			},
			AUTHOR: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "author",
				Message: "This field should be valid string not greater than 100 characters",
			},
			CONTENT: Error{
				Code:    400,
				Type:    "invalid_request_error",
				Field:   "content",
				Message: "This field should be valid string not greater than 1000 characters",
			},
		},

		OBJ_NOT_FOUND: Error{
			Code:    404,
			Type:    "object_not_found",
			Message: "The requested object does not exist or already deleted.",
		},
	}

	return e
}
