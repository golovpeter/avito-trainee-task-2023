// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/segment/changeForUser": {
            "post": {
                "description": "delete and create user segments by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Change user segments",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/change_user_segments.ChangeUserSegmentsIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorOut"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorOut"
                        }
                    }
                }
            }
        },
        "/segment/create": {
            "post": {
                "description": "adding a new segment by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Create new segment",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/create_segment.CreateSegmentIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorOut"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorOut"
                        }
                    }
                }
            }
        },
        "/segment/delete": {
            "post": {
                "description": "deleting segment by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Delete segment",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/delete_segment.DeleteSegmentIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorOut"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorOut"
                        }
                    }
                }
            }
        },
        "/segments/user/{user_id}": {
            "get": {
                "description": "getting all user segments by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Get all user segments",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/get_user_segments.GetUserSegmentsOut"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorOut"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorOut"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "change_user_segments.ChangeUserSegmentsIn": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "add_segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "AVITO_VOICE_MESSAGE",
                        "AVITO_DICSOUNT_30"
                    ]
                },
                "delete_segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "AVITO_DICOUNT_50"
                    ]
                },
                "expired_at": {
                    "type": "string",
                    "example": "2023-08-27T15:40:00Z"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1000
                }
            }
        },
        "common.ErrorOut": {
            "type": "object",
            "properties": {
                "error_message": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "create_segment.CreateSegmentIn": {
            "type": "object",
            "required": [
                "segment_slug"
            ],
            "properties": {
                "percent_users": {
                    "type": "integer"
                },
                "segment_slug": {
                    "type": "string",
                    "example": "AVITO_VOICE_MESSAGE"
                }
            }
        },
        "delete_segment.DeleteSegmentIn": {
            "type": "object",
            "required": [
                "segment_slug"
            ],
            "properties": {
                "segment_slug": {
                    "type": "string",
                    "example": "AVITO_VOICE_MESSAGE"
                }
            }
        },
        "get_user_segments.GetUserSegmentsOut": {
            "type": "object",
            "properties": {
                "segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "AVITO_VOICE_MESSAGE",
                        " AVITO_DISCOUNT_30"
                    ]
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Dynamic User Segmentation api Swagger API",
	Description:      "API for Golang Dynamic User Segmentation api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}