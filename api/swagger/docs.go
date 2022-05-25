// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2021-10-23 17:59:22.812821 +0800 CST m=+0.027672501

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
        "/account/api/basic/source/v1": {
            "get": {
                "description": "获取账户基础创建方式",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户基础信息"
                ],
                "summary": "获取账户基础创建方式",
                "parameters": [
                    {
                        "type": "string",
                        "description": "认证信息 eg:xxxx-xxxx-xxxx-xxx",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "response": {
                    "200": {
                        "description": "修改成功",
                        "schema": {
                            "$ref": "#/definitions/contant.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "参数解析错误",
                        "schema": {
                            "$ref": "#/definitions/contant.FailedRes"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/contant.FailedRes"
                        }
                    }
                }
            }
        },
        "/account/api/unauthorized/basic/source/v1": {
            "get": {
                "description": "获取账户基础创建方式-无auth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户基础信息"
                ],
                "summary": "获取账户基础创建方式-无auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "公司corpId",
                        "name": "corpId",
                        "in": "query",
                        "required": true
                    }
                ],
                "response": {
                    "200": {
                        "description": "修改成功",
                        "schema": {
                            "$ref": "#/definitions/contant.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "参数解析错误",
                        "schema": {
                            "$ref": "#/definitions/contant.FailedRes"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/contant.FailedRes"
                        }
                    }
                }
            }
        },
        "/account/api/unauthorized/get_register_code/v1": {
            "post": {
                "description": "获取注册码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户注册"
                ],
                "summary": "获取注册码",
                "parameters": [
                    {
                        "description": "创建参数",
                        "name": "create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.GetRegisterCodeReq"
                        }
                    }
                ],
                "response": {
                    "200": {
                        "description": "修改成功",
                        "schema": {
                            "$ref": "#/definitions/contant.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "参数解析错误",
                        "schema": {
                            "$ref": "#/definitions/contant.FailedRes"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/contant.FailedRes"
                        }
                    }
                }
            }
        },
        "/account/api/unauthorized/get_register_info/v1": {
            "post": {
                "description": "查询注册状态信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "账户注册"
                ],
                "summary": "查询注册状态信息",
                "parameters": [
                    {
                        "description": "创建参数",
                        "name": "create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/router.GetRegisterInfoReq"
                        }
                    }
                ],
                "response": {
                    "200": {
                        "description": "修改成功",
                        "schema": {
                            "$ref": "#/definitions/contant.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "参数解析错误",
                        "schema": {
                            "$ref": "#/definitions/contant.FailedRes"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/contant.FailedRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "contant.FailedRes": {
            "type": "object",
            "required": [
                "data",
                "msg",
                "success"
            ],
            "properties": {
                "data": {
                    "description": "返回的数据",
                    "type": "object"
                },
                "msg": {
                    "description": "请求结果的message",
                    "type": "string",
                    "example": "req_data_val_error"
                },
                "success": {
                    "description": "请求结果，失败:false，成功:true",
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "contant.SuccessRes": {
            "type": "object",
            "required": [
                "data",
                "msg",
                "success"
            ],
            "properties": {
                "data": {
                    "description": "返回的数据",
                    "type": "object"
                },
                "msg": {
                    "description": "请求结果的message",
                    "type": "string",
                    "example": "ok"
                },
                "success": {
                    "description": "请求结果，失败:false，成功:true",
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "router.GetRegisterCodeReq": {
            "type": "object",
            "required": [
                "template_id"
            ],
            "properties": {
                "template_id": {
                    "description": "推广包的id 由业务方自主传入",
                    "type": "string"
                }
            }
        },
        "router.GetRegisterInfoReq": {
            "type": "object",
            "required": [
                "registerCode"
            ],
            "properties": {
                "registerCode": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "mk-dev.dustess.com",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "示例服务(mid-oauth2-srv)",
	Description: "示例服务",
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