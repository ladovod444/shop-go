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
            "name": "Dmitrii",
            "email": "ladovod@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/product": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a product",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves product based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Updates product based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "summary": "Delete a product by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/products": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "paginating results - ?page=1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "product search - ?title=Some title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "product search - ?description=Some descr",
                        "name": "description",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Product": {
            "type": "object",
            "properties": {
                "CurrentPrice": {
                    "description": "CurrentPrice float32 ` + "`" + `json:\"CurrentPrice\" example:\"10.01\" sql:\"type:decimal(10,2);\"` + "`" + `\nRegularPrice float32 ` + "`" + `json:\"RegularPrice\" example:\"10.01\" sql:\"type:decimal(10,2);\"` + "`" + `\nCreatedAt    time.Time ` + "`" + `json:\"CreatedAt\" example:\"2006-02-01T15:04:05Z\" gorm:\"default:current_timestamp\"` + "`" + `\nUpdatedAt    time.Time ` + "`" + `json:\"UpdatedAt\" example:\"2006-02-01T15:04:05Z\" gorm:\"default:current_timestamp\"` + "`" + `",
                    "type": "number",
                    "example": 10.01
                },
                "Description": {
                    "type": "string",
                    "example": "Product description"
                },
                "Image": {
                    "description": "CreatedAt    time.Time ` + "`" + `gorm:\"autoCreateTime:true\" json:\"createdAt\"` + "`" + `\nUpdatedAt    time.Time ` + "`" + `gorm:\"autoUpdateTime:true\" json:\"updatedAt\"` + "`" + `",
                    "type": "string",
                    "example": "Product Image"
                },
                "RegularPrice": {
                    "type": "number",
                    "example": 10.01
                },
                "Sku": {
                    "type": "string",
                    "example": "Product sku"
                },
                "Title": {
                    "type": "string",
                    "example": "Product title"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Products API",
	Description:      "Swagger API for Songs library API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
