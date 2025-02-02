// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/follow": {
            "post": {
                "description": "Follow a user by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "follow"
                ],
                "summary": "Follow a user",
                "parameters": [
                    {
                        "description": "Follow request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FollowRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/like": {
            "post": {
                "description": "Add a like to a specific post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "likes"
                ],
                "summary": "Like a post",
                "parameters": [
                    {
                        "description": "Like request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LikeRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a like from a specific post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "likes"
                ],
                "summary": "Unlike a post",
                "parameters": [
                    {
                        "description": "Unlike request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LikeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/like/{postID}": {
            "get": {
                "description": "Get the number of likes for a specific post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "likes"
                ],
                "summary": "Get likes count for a post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "postID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LikeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/photo": {
            "post": {
                "description": "Upload a new photo and associate it with a user",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Upload a photo",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Photo file",
                        "name": "photo",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/photo/{key}": {
            "get": {
                "description": "Retrieve a photo by its unique key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Get photo by key",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Photo key",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/post": {
            "post": {
                "description": "Create a new post with photo upload",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Create a new post",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Photo file",
                        "name": "photo",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Post caption",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/post/all": {
            "get": {
                "description": "Get a list of all posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get all posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Posts"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/post/user/{userId}": {
            "get": {
                "description": "Get all posts for a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get user's posts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Posts"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/post/{id}": {
            "get": {
                "description": "Get details of a specific post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get post by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Posts"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/post/{userId}": {
            "delete": {
                "description": "Delete a post by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Delete a post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/unfollow": {
            "delete": {
                "description": "Unfollow a user by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "follow"
                ],
                "summary": "Unfollow a user",
                "parameters": [
                    {
                        "description": "Unfollow request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FollowRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/user": {
            "put": {
                "description": "Update details of an existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user information",
                "parameters": [
                    {
                        "description": "Updated user information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Register a new user in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get details of a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user from the system by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customResponse.CustomStatus"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/user/{userID}/followers": {
            "get": {
                "description": "Retrieve a list of all followers for a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all followers of a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.GetAllFollowersResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        },
        "/user/{userID}/following": {
            "get": {
                "description": "Retrieve a list of all users followed by a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users followed by a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.GetAllFollowersResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/customResponse.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "customResponse.CustomStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "integer"
                }
            }
        },
        "customResponse.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.FollowRequest": {
            "type": "object",
            "properties": {
                "follower_id": {
                    "type": "integer"
                },
                "following_id": {
                    "type": "integer"
                }
            }
        },
        "models.GetAllFollowersResponse": {
            "type": "object",
            "properties": {
                "profile_pic": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.LikeRequest": {
            "type": "object",
            "properties": {
                "post_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.LikeResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "models.Posts": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "profile_pic": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Kirkagram API",
	Description:      "This is the API for Kirkagram application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
