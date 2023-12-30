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
            "email": "fiber@swagger.io"
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
        "/auth": {
            "post": {
                "description": "Get access and refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get Auth tokens using credentials",
                "parameters": [
                    {
                        "description": "Auth user",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.FormAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Get access and refresh tokens",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get new tokens using refresh token",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/ingredient": {
            "get": {
                "description": "Get ingredients list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredient"
                ],
                "summary": "Get ingredients list",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "perPage",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domains.Ingredient"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new ingredient",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredient"
                ],
                "summary": "Create new ingredient",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.IngredientForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseAdd"
                        }
                    }
                }
            }
        },
        "/ingredient/{id}": {
            "get": {
                "description": "Retrieve ingredient by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredient"
                ],
                "summary": "Retrieve ingredient by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "request",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domains.Ingredient"
                        }
                    }
                }
            },
            "put": {
                "description": "Update ingredient",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredient"
                ],
                "summary": "Update ingredient",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "request",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.IngredientForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseStatus"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domains.Ingredient": {
            "type": "object",
            "properties": {
                "barcode": {
                    "type": "string"
                },
                "calories": {
                    "type": "integer"
                },
                "carbs": {
                    "type": "integer"
                },
                "fats": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "manufacturer": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "proteins": {
                    "type": "integer"
                }
            }
        },
        "forms.FormAuth": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "forms.IngredientForm": {
            "type": "object",
            "properties": {
                "barcode": {
                    "type": "string"
                },
                "calories": {
                    "type": "integer"
                },
                "carbs": {
                    "type": "integer"
                },
                "fats": {
                    "type": "integer"
                },
                "manufacturer": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "proteins": {
                    "type": "integer"
                }
            }
        },
        "responses.ResponseAdd": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
