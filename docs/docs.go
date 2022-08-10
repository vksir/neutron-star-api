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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/calculator/ipv4": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calculator"
                ],
                "summary": "解析 ipv4",
                "parameters": [
                    {
                        "type": "string",
                        "name": "ip_addr",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/calculator.IPv4"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Err"
                        }
                    }
                }
            }
        },
        "/calculator/ipv6": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calculator"
                ],
                "summary": "解析 ipv6",
                "parameters": [
                    {
                        "type": "string",
                        "name": "ip_addr",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/calculator.IPv6"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Err"
                        }
                    }
                }
            }
        },
        "/pixiv": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pixiv"
                ],
                "summary": "获取 pixiv 图片 url",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "num",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pixiv.ImageInDB"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Err"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Err"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "calculator.IPv4": {
            "type": "object",
            "properties": {
                "cidr": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                },
                "end": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "mask_bits": {
                    "type": "integer"
                },
                "start": {
                    "type": "string"
                },
                "subnet_mask": {
                    "type": "string"
                },
                "wildcard_mask": {
                    "type": "string"
                }
            }
        },
        "calculator.IPv6": {
            "type": "object",
            "properties": {
                "long": {
                    "$ref": "#/definitions/calculator.SubIPv6"
                },
                "short": {
                    "$ref": "#/definitions/calculator.SubIPv6"
                }
            }
        },
        "calculator.SubIPv6": {
            "type": "object",
            "properties": {
                "cidr": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                },
                "end": {
                    "type": "string"
                },
                "ip": {
                    "type": "string"
                },
                "mask_bits": {
                    "type": "integer"
                },
                "start": {
                    "type": "string"
                }
            }
        },
        "model.Err": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                }
            }
        },
        "pixiv.ImageInDB": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_r18": {
                    "type": "boolean"
                },
                "origin_url": {
                    "type": "string"
                },
                "relative_path": {
                    "type": "string"
                },
                "tags": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Neutron Star API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
