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
        "/room-types": {
            "get": {
                "description": "List room types",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room-types"
                ],
                "summary": "List room types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/rooms.RoomTypesResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create room type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room-types"
                ],
                "summary": "Create room type",
                "parameters": [
                    {
                        "description": "RoomTypeCreate",
                        "name": "RoomType",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rooms.RoomTypeCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rooms.RoomTypeCreate"
                        }
                    }
                }
            }
        },
        "/room-types/{id}": {
            "put": {
                "description": "Update room type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room-types"
                ],
                "summary": "Update room type",
                "parameters": [
                    {
                        "type": "string",
                        "description": "RoomType ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "RoomTypeUpdate",
                        "name": "RoomType",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rooms.RoomTypeUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rooms.RoomTypeUpdate"
                        }
                    }
                }
            }
        },
        "/rooms": {
            "post": {
                "description": "Create room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Create room",
                "parameters": [
                    {
                        "description": "RoomCreate",
                        "name": "Room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rooms.RoomCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rooms.RoomCreate"
                        }
                    }
                }
            }
        },
        "/rooms/{id}": {
            "put": {
                "description": "Update room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Update room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Room ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "RoomUpdate",
                        "name": "Room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rooms.RoomUpdate"
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
        "/users": {
            "get": {
                "description": "Get a list of users with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "List users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default is 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page (default is 10)",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Cursor for pagination",
                        "name": "cursor",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/users.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "UserCreate",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.UserCreate"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Image": {
            "type": "object",
            "properties": {
                "cloud_name": {
                    "type": "string"
                },
                "extension": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "rooms.RoomCreate": {
            "type": "object",
            "required": [
                "name",
                "room_type_id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.Image"
                    }
                },
                "name": {
                    "type": "string"
                },
                "room_type_id": {
                    "type": "string"
                }
            }
        },
        "rooms.RoomTypeCreate": {
            "type": "object",
            "required": [
                "bed_count",
                "charges_for_cancellation",
                "food_option",
                "name",
                "price"
            ],
            "properties": {
                "bed_count": {
                    "type": "integer"
                },
                "charges_for_cancellation": {
                    "type": "number"
                },
                "food_option": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "rooms.RoomTypeUpdate": {
            "type": "object",
            "properties": {
                "bed_count": {
                    "type": "integer"
                },
                "charges_for_cancellation": {
                    "type": "number"
                },
                "food_option": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "rooms.RoomTypesResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "rooms.RoomUpdate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.Image"
                    }
                },
                "name": {
                    "type": "string"
                },
                "room_type_id": {
                    "type": "string"
                }
            }
        },
        "users.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "users.UserCreate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
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
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
