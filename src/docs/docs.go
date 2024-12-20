// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mike Zinyoni",
            "email": "mzinyoni7"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1": {
            "get": {
                "description": "Gets Index",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Index"
                ],
                "summary": "Index",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login/{guid}": {
            "post": {
                "description": "Gets Credentials and Returns Auth Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "structs.LoginParameters GUID",
                        "name": "guid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.LoginResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/refresh": {
            "post": {
                "description": "Gets Credentials and Returns Auth Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Starts with Bearer ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.RefreshResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/secure": {
            "get": {
                "description": "Gets Credentials and Returns Auth Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Secure",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Starts with Bearer ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.SecureResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.LoginResponse": {
            "type": "object",
            "properties": {
                "refresh": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "structs.RefreshResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "structs.SecureResponse": {
            "type": "object",
            "properties": {
                "expiration": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer \" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go Auth API",
	Description:      "Go Auth API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
