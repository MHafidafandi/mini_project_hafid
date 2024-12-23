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
        "/foods": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve all available food items from the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Get all food items",
                "responses": {
                    "200": {
                        "description": "List of food items",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-array_models_Food"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new food entry with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Create a new food item",
                "parameters": [
                    {
                        "description": "Food details",
                        "name": "foodDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.FoodRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Food created successfully",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        },
        "/foods/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a specific food item by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Get a food item by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Food ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Food item found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-models_Food"
                        }
                    },
                    "404": {
                        "description": "Food item not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a specific food item with new details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Update an existing food item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Food ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated food details",
                        "name": "foodDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.FoodUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Food updated successfully",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "404": {
                        "description": "Food item not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a specific food item by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Delete a food item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Food ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Food deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "404": {
                        "description": "Food item not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        },
        "/orders": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new order with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "Order details",
                        "name": "orderDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Order created successfully",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        },
        "/orders/users/{user_id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve all orders for a specific user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get all orders by user ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of orders",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-array_models_Order"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a specific order by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get a specific order by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-models_Order"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "User login with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful, token returned",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-map_string_string"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "409": {
                        "description": "Wrong email or password",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Register a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "userDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "409": {
                        "description": "Email already exists",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve a specific user by their ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-models_User"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update user information by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User update details",
                        "name": "userDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated successfully",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a user by their ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        },
        "/webhook/midtrans": {
            "post": {
                "description": "Receives payment status updates from Midtrans and updates the transaction status accordingly",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "webhooks"
                ],
                "summary": "Handle Midtrans payment notification",
                "parameters": [
                    {
                        "description": "Midtrans notification payload",
                        "name": "notificationPayloads",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Transaction status updated successfully",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "404": {
                        "description": "Order not found",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse-any"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "map_string_string": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "models.Food": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expiry_date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "order_items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OrderItem"
                    }
                },
                "payment": {
                    "$ref": "#/definitions/models.Payment"
                },
                "payment_id": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.OrderItem": {
            "type": "object",
            "properties": {
                "food": {
                    "$ref": "#/definitions/models.Food"
                },
                "food_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "order": {
                    "$ref": "#/definitions/models.Order"
                },
                "order_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "models.Payment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "orders": {
                    "$ref": "#/definitions/models.Order"
                },
                "payment_link": {
                    "type": "string"
                },
                "payment_status": {
                    "type": "string"
                },
                "payment_type": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Order"
                    }
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "request.CreateOrderRequest": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "order_items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.OrderItem"
                    }
                },
                "payment_type": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "number"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "request.FoodRequest": {
            "type": "object",
            "required": [
                "expiry_date",
                "location",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "expiry_date": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "request.FoodUpdate": {
            "type": "object",
            "required": [
                "location",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.OrderItem": {
            "type": "object",
            "required": [
                "food_id",
                "quantity"
            ],
            "properties": {
                "food_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "request.UserRequest": {
            "type": "object",
            "required": [
                "address",
                "email",
                "name",
                "password",
                "phone",
                "role"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "request.UserUpdate": {
            "type": "object",
            "required": [
                "address",
                "name",
                "phone",
                "role"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "response.BaseResponse-any": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "response.BaseResponse-array_models_Food": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Food"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "response.BaseResponse-array_models_Order": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Order"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "response.BaseResponse-map_string_string": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/map_string_string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "response.BaseResponse-models_Food": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.Food"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "response.BaseResponse-models_Order": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.Order"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "response.BaseResponse-models_User": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.User"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "greenenvironment",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Mini Project EcoBite",
	Description:      "This is a sample server Swagger server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
