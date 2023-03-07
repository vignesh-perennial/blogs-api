# blogs-api
Backend API for a simple blog system.

To run,

  sh blog.sh

sample Curls and response:

POST:

          curl --location --request POST 'localhost:8080/articals' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "title": "blog for golang",
          "content": "Sample Content for blogs api",
          "author": "devignesh"
      }'

Response:

        {
        "status": 201,
        "message": "Blog Created successfully",
        "data": {
            "id": "6406d6642a5852e5990c2cb3"
        }
    }

Error:

            curl --location --request POST 'localhost:8080/articals' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "title": "",
            "content": "Sample Content for blogs api",
            "author": "devignesh"
        }'

Res:

        {
            "code": 400,
            "message": "invalid_request_error",
            "description": "The request was unacceptable, due to missing a required parameter or invalid parameter.",
            "errors": [
                {
                    "field": "title",
                    "message": "This field should be valid string not greater than 255 characters"
                }
            ]
        }

GET:

      curl --location --request GET 'localhost:8080/articals/6406d6642a5852e5990c2cb3'

Response:

        {
            "status": 200,
            "message": "The details of the blog for requested Id",
            "data": {
                "id": "6406d6642a5852e5990c2cb3",
                "title": "blog for golang",
                "content": "Sample Content for blogs api",
                "author": "devignesh"
            }
        }

Error Response:

          {
          "code": 400,
          "message": "invalid_request_error",
          "description": "The request was unacceptable, due to missing a required parameter or invalid parameter.",
          "errors": [
              {
                  "field": "id",
                  "message": "This field should be a valid blog id"
              }
          ]
      }


Err:

            curl --location --request POST 'localhost:8080/articals' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "title": "",
            "content": "Sample Content for blogs api",
            "author": ""
        }'

Res:

            {
            "code": 400,
            "message": "invalid_request_error",
            "description": "The request was unacceptable, due to missing a required parameter or invalid parameter.",
            "errors": [
                {
                    "field": "title",
                    "message": "This field should be valid string not greater than 255 characters"
                },
                {
                    "field": "author",
                    "message": "This field should be valid string not greater than 100 characters"
                }
            ]
        }


GET All:

    curl --location --request GET 'localhost:8080/articals'

Response:

            {
            "status": 200,
            "message": "Blog List success",
            "data": [
                {
                    "id": "6405af0349229973f3c53e50",
                    "title": "test blog",
                    "content": "test content",
                    "author": "vicky"
                },
                {
                    "id": "64061d1b4d606f041fd9feca",
                    "title": "test blog 2",
                    "content": "test content 2",
                    "author": "vicky2"
                },
                {
                    "id": "6406d6642a5852e5990c2cb3",
                    "title": "blog for golang",
                    "content": "Sample Content for blogs api",
                    "author": "devignesh"
                }
            ]
        }
