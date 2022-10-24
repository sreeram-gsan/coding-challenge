// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://sreeramganesan.com",
        "contact": {
            "name": "Sreeram Ganesan",
            "url": "https://sreeramganesan.com",
            "email": "srga8641@colorado.edu"
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
        "/items": {
            "get": {
                "description": "Responds with the list of all items as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get items array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.item"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Takes a item JSON and store in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Store a new item",
                "parameters": [
                    {
                        "description": "item JSON",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.item"
                        }
                    }
                }
            }
        },
        "/items/{id}": {
            "get": {
                "description": "Returns the item whose id value matches the provided id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get a single item by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "search item by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.item"
                        }
                    }
                }
            },
            "patch": {
                "description": "Takes a item JSON and updates its value in DB.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Update an existing item",
                "parameters": [
                    {
                        "description": "item JSON",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.item"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "update item by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.item"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.item": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                },
                "unit_price": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/inventory/v1",
	Schemes:          []string{},
	Title:            "Inventory API",
	Description:      "An inventory management API with Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
