// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Auth to get token",
                "produces": [
                    "application/json"
                ],
                "summary": "Provides a JSON web token",
                "operationId" : "Authenticate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/main"
                        }
                    }
                }
            }
        },
        "/orders":{
            "get": {
                "description": "Get all the orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders",
                    "list"
                ],
                "summary": "List existing orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Orders"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Create a new Orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders",
                    "create"
                ],
                "summary": "Create new orders",
                "parameters": [
                    {
                        "description": "Create orders",
                        "name": "customerName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    {
                        "description": "Create orders",
                        "name": "orderedAt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.InsertNewOrder"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.InsertNewOrder"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controller.InsertNewOrder"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get" : {
                "description": "Get order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders",
                    "list"
                ],
                "summary": "List orders by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order ID",
                        "name": "orderID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Orders"
                        }
                    }
                }
            },
            "put" : {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Update a single order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update Orders",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order ID",
                        "name": "orderID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type" : "string",
                        "description": "customer name",
                        "name": "customerName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateOrderByID"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateOrderByID"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateOrderByID"
                        }
                    }
                }
            },
            "delete" : {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Delete an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Delete Order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order ID",
                        "name": "orderID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.DeleteOrder"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions//controllers.DeleteOrder"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions//controllers.DeleteOrder"
                        }
                    }
                }
            }
        }
	},
	"definitions": {
        "main": {
            "type": "object",
            "properties": {
                "code": {
                    "type" : "integer"
                },
                "expire" : {
                    "type" : "date"
                },
                "token" : {
                    "type" : "string"
                }
            }
        },
        "models.Order": {
            "type" : "object",
            "properties": {
                "orderID" : {
                    "type" : "integer"
                },
                "customerName" : {
                    "type" : "string"
                },
                "orderedAt" : {
                    "type" : "string"
                },
                "items" : {
                    "type" : "array"
                }
            }
        },
        "controllers.InsertNewOrder": {
            "type": "object",
            "properties": {
                "order": {
                    "status" : {
                        "type" : "string"
                    },
                    "order" : {
                        "type" : "string"
                    }  
                }
            }
        },
        "controllers.UpdateOrder": {
            "type": "object",
            "properties": {
                "order": {
                    "status" : {
                        "type" : "string"
                    },
                    "order" : {
                        "type" : "string"
                    }  
                }
            }
        },
        "controllers.DeleteOrder":{
            "type": "object",
            "properties": {
                "order": {
                    "status" : {
                        "type" : "string"
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
