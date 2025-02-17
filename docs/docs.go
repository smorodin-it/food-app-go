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
                            "$ref": "#/definitions/forms.AuthForm"
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "$ref": "#/definitions/responses.ResponseApi-array_domains_Ingredient"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseAdd"
                        }
                    }
                }
            }
        },
        "/ingredient/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "$ref": "#/definitions/responses.ResponseApi-domains_Ingredient"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseStatus"
                        }
                    }
                }
            }
        },
        "/inventory": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get inventory list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory"
                ],
                "summary": "Get inventory list",
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
                            "$ref": "#/definitions/responses.ResponseApi-array_domains_Inventory"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new inventory",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory"
                ],
                "summary": "Create new inventory",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.InventoryForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseAdd"
                        }
                    }
                }
            }
        },
        "/inventory/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve inventory by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory"
                ],
                "summary": "Retrieve inventory by id",
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
                            "$ref": "#/definitions/responses.ResponseApi-domains_Inventory"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update inventory",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory"
                ],
                "summary": "Update inventory",
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
                            "$ref": "#/definitions/forms.InventoryForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseStatus"
                        }
                    }
                }
            }
        },
        "/meal": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get meals list by auth user id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Get meals list by auth user id",
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
                            "$ref": "#/definitions/responses.ResponseApi-array_domains_Meal"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new meal",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Create new meal",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.MealForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseAdd"
                        }
                    }
                }
            }
        },
        "/meal/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get meals list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Get meals list",
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
                            "$ref": "#/definitions/responses.ResponseApi-array_domains_Meal"
                        }
                    }
                }
            }
        },
        "/meal/ingredient": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add ingredient to meal",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Add ingredient to meal",
                "parameters": [
                    {
                        "description": "body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.MealIngredientAddForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseAdd"
                        }
                    }
                }
            }
        },
        "/meal/ingredient/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update weight of ingredient in meal",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Update weight of ingredient in meal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ingredient in meal id",
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
                            "$ref": "#/definitions/forms.MealIngredientUpdateForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseAdd"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete ingredient from meal",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Delete ingredient from meal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ingredient in meal id",
                        "name": "request",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseStatus"
                        }
                    }
                }
            }
        },
        "/meal/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve meal by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Retrieve meal by id",
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
                            "$ref": "#/definitions/responses.ResponseApi-domains_Meal"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update meal",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Update meal",
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
                            "$ref": "#/definitions/forms.MealForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseStatus"
                        }
                    }
                }
            }
        },
        "/meal/{id}/ingredient": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get ingredients list in meal",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meal"
                ],
                "summary": "Get ingredients list in meal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "meal id",
                        "name": "request",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-array_responses_MealIngredientResp"
                        }
                    }
                }
            }
        },
        "/measurement": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get measurements list by user id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Measurement"
                ],
                "summary": "Get measurements list by user id",
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
                            "$ref": "#/definitions/responses.ResponseApi-array_domains_Measurement"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new measurement",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Measurement"
                ],
                "summary": "Create new measurement",
                "parameters": [
                    {
                        "type": "string",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "weight",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseAdd"
                        }
                    }
                }
            }
        },
        "/measurement/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get measurement by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Measurement"
                ],
                "summary": "Get measurement by id",
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
                            "$ref": "#/definitions/responses.ResponseApi-domains_Measurement"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update measurement",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Measurement"
                ],
                "summary": "Update measurement",
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
                            "$ref": "#/definitions/forms.MeasurementUpdateForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseStatus"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete measurement",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Measurement"
                ],
                "summary": "Delete measurement",
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
                            "$ref": "#/definitions/responses.ResponseApi-responses_ResponseStatus"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domains.Ingredient": {
            "type": "object",
            "required": [
                "id",
                "manufacturer",
                "name"
            ],
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
        "domains.Inventory": {
            "type": "object",
            "required": [
                "id",
                "name",
                "weight"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "domains.Meal": {
            "type": "object",
            "required": [
                "id",
                "name",
                "totalWeight"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "totalWeight": {
                    "type": "integer"
                }
            }
        },
        "domains.Measurement": {
            "type": "object",
            "required": [
                "id",
                "weight"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "forms.AuthForm": {
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
        "forms.IngredientForm": {
            "type": "object",
            "required": [
                "manufacturer",
                "name"
            ],
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
        "forms.InventoryForm": {
            "type": "object",
            "required": [
                "name",
                "weight"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "forms.MealForm": {
            "type": "object",
            "required": [
                "name",
                "totalWeight"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "totalWeight": {
                    "type": "integer"
                }
            }
        },
        "forms.MealIngredientAddForm": {
            "type": "object",
            "required": [
                "ingredientId",
                "mealId",
                "totalWeight"
            ],
            "properties": {
                "ingredientId": {
                    "type": "string"
                },
                "mealId": {
                    "type": "string"
                },
                "totalWeight": {
                    "type": "integer"
                }
            }
        },
        "forms.MealIngredientUpdateForm": {
            "type": "object",
            "required": [
                "totalWeight"
            ],
            "properties": {
                "totalWeight": {
                    "type": "integer"
                }
            }
        },
        "forms.MeasurementUpdateForm": {
            "type": "object",
            "required": [
                "weight"
            ],
            "properties": {
                "weight": {
                    "type": "integer"
                }
            }
        },
        "responses.MealIngredientResp": {
            "type": "object",
            "required": [
                "id",
                "name",
                "weight"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "weight": {
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
        "responses.ResponseApi-array_domains_Ingredient": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domains.Ingredient"
                    }
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-array_domains_Inventory": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domains.Inventory"
                    }
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-array_domains_Meal": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domains.Meal"
                    }
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-array_domains_Measurement": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domains.Measurement"
                    }
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-array_responses_MealIngredientResp": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.MealIngredientResp"
                    }
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-domains_Ingredient": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/domains.Ingredient"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-domains_Inventory": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/domains.Inventory"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-domains_Meal": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/domains.Meal"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-domains_Measurement": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/domains.Measurement"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-responses_ResponseAdd": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/responses.ResponseAdd"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseApi-responses_ResponseStatus": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/responses.ResponseStatus"
                },
                "error": {
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
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "food-app backend",
	Description:      "swagger for food-app api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
