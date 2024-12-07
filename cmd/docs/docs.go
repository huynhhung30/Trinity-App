// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/coupon-list": {
            "get": {
                "description": "Get Coupon List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coupon"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "from_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "keyword",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "to_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "list coupon",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/create-coupon": {
            "post": {
                "description": "Create a new coupon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coupon"
                ],
                "summary": "Create a coupon",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Require Email and Password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CouponModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "406": {
                        "description": "Not Acceptable"
                    }
                }
            }
        },
        "/member-login": {
            "post": {
                "description": "Authenticates a user using email and password, returns a JWT token if successful",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Authenticate a user",
                "parameters": [
                    {
                        "description": "User login object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UsersLoginRequestModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/member-register": {
            "post": {
                "description": "Create a new user with the given input data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "Require Email and Password",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "406": {
                        "description": "Not Acceptable"
                    }
                }
            }
        },
        "/silver-register": {
            "post": {
                "description": "Require input Id of coupon",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Silver Register",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "POST",
                        "name": "id_coupon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.QueryCoupon"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CouponModel": {
            "type": "object",
            "properties": {
                "count_registrations": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "discount": {
                    "type": "string"
                },
                "due_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_valid": {
                    "type": "boolean"
                },
                "tittle": {
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
        "models.QueryCoupon": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.UserModel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "email_verified": {
                    "type": "boolean"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_silver": {
                    "type": "boolean"
                },
                "last_name": {
                    "type": "string"
                },
                "login_method": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "social_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_type": {
                    "type": "string"
                }
            }
        },
        "models.UsersLoginRequestModel": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "login_method": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "social_id": {
                    "type": "string"
                },
                "user_type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5001",
	BasePath:         "/api/v1/trinity",
	Schemes:          []string{},
	Title:            "Trinity App Tag Service API",
	Description:      "A tag service api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}