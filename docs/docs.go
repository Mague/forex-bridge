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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/operations": {
            "get": {
                "description": "get operations",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operations"
                ],
                "summary": "List operations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Operation"
                            }
                        }
                    }
                }
            }
        },
        "/operations/add": {
            "post": {
                "description": "add by json operation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Operations"
                ],
                "summary": "Add an operation",
                "parameters": [
                    {
                        "description": "Add operation",
                        "name": "operation",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payloads.Operation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Operation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "models.Operation": {
            "type": "object",
            "required": [
                "ORD_TYPE",
                "POS_NUM",
                "PRICE",
                "SL",
                "SYMBOL",
                "TP",
                "TR_ID"
            ],
            "properties": {
                "ORD_TYPE": {
                    "type": "integer"
                },
                "POS_NUM": {
                    "type": "number"
                },
                "PRICE": {
                    "type": "number"
                },
                "SL": {
                    "type": "number"
                },
                "SYMBOL": {
                    "type": "string"
                },
                "TP": {
                    "type": "number"
                },
                "TR_ID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "payloads.Operation": {
            "type": "array",
            "items": {
                "type": "object",
                "required": [
                    "ORD_TYPE",
                    "POS_NUM",
                    "PRICE",
                    "SL",
                    "SYMBOL",
                    "TP",
                    "TR_ID"
                ],
                "properties": {
                    "ORD_TYPE": {
                        "type": "integer"
                    },
                    "POS_NUM": {
                        "type": "number"
                    },
                    "PRICE": {
                        "type": "number"
                    },
                    "SL": {
                        "type": "number"
                    },
                    "SYMBOL": {
                        "type": "string"
                    },
                    "TP": {
                        "type": "number"
                    },
                    "TR_ID": {
                        "type": "integer"
                    }
                }
            }
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
