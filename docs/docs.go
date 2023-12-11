// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v1/api/admin/profile/platform/add": {
            "post": {
                "description": "添加新的社交平台",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "create new social platform",
                "parameters": [
                    {
                        "description": "platforms",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddPlatformDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/admin/profile/platform/all": {
            "get": {
                "description": "查看所有社交平台",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "get all social platform",
                "responses": {}
            }
        },
        "/v1/api/admin/stat/count": {
            "get": {
                "description": "统计平台信息 今日新增贴子数 评论数 用户数  当月新增用户数 评论数 用户数",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "get stat data",
                "parameters": [
                    {
                        "enum": [
                            1,
                            2,
                            3,
                            4,
                            5,
                            6,
                            7,
                            8,
                            9,
                            10,
                            11,
                            12
                        ],
                        "type": "integer",
                        "name": "month",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/bio/{name}": {
            "get": {
                "description": "个人主页分享",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "profile share",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the user name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/comment/create": {
            "post": {
                "description": "创建评论",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "create comment",
                "parameters": [
                    {
                        "description": "comment content",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CommentDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/comment/del": {
            "delete": {
                "description": "删除评论",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "comment delete",
                "responses": {}
            }
        },
        "/v1/api/comment/operater": {
            "post": {
                "description": "操作评论(赞,踩,收藏,分享)",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "comment operator",
                "parameters": [
                    {
                        "description": "operator code",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CommentOperatorDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/comment/reply": {
            "post": {
                "description": "回复评论",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "reply comment",
                "responses": {}
            }
        },
        "/v1/api/comment/view": {
            "get": {
                "description": "查看评论",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comment"
                ],
                "summary": "get comment",
                "responses": {}
            }
        },
        "/v1/api/post/create": {
            "post": {
                "description": "发帖",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "create post",
                "responses": {}
            }
        },
        "/v1/api/post/del": {
            "delete": {
                "description": "删除帖子",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "operator post",
                "responses": {}
            }
        },
        "/v1/api/post/operater": {
            "post": {
                "description": "操作帖子 赞 踩 收藏",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "operator post",
                "responses": {}
            }
        },
        "/v1/api/post/search": {
            "get": {
                "description": "搜索帖子、文章",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "search post",
                "responses": {}
            }
        },
        "/v1/api/tag/follow": {
            "post": {
                "description": "关注话题",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "topic"
                ],
                "summary": "follow tag",
                "parameters": [
                    {
                        "description": "topic",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TagFollowDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/tag/hide": {
            "get": {
                "description": "隐藏话题标签",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "topic"
                ],
                "summary": "hide tag",
                "parameters": [
                    {
                        "type": "string",
                        "name": "tagid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/tag/list": {
            "get": {
                "description": "获取话题列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "topic"
                ],
                "summary": "get topic",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 20,
                        "minimum": 1,
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/tag/new": {
            "post": {
                "description": "创建新话题标签",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "topic"
                ],
                "summary": "create new topic",
                "parameters": [
                    {
                        "description": "topic",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TagDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/user/create": {
            "post": {
                "description": "创建新用户",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "admin account",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SignupDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "user login",
                "parameters": [
                    {
                        "description": "the user account",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/user/name/check": {
            "get": {
                "description": "检查昵称",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "check user name",
                "parameters": [
                    {
                        "minLength": 4,
                        "type": "string",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/api/user/profile": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "get user profile",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.AddPlatformDTO": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dto.CommentDTO": {
            "type": "object",
            "required": [
                "content",
                "post"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "minLength": 1
                },
                "post": {
                    "type": "integer"
                }
            }
        },
        "dto.CommentOperatorDTO": {
            "type": "object",
            "required": [
                "cid",
                "op"
            ],
            "properties": {
                "cid": {
                    "type": "integer"
                },
                "op": {
                    "type": "integer",
                    "enum": [
                        1,
                        2,
                        4,
                        8,
                        16
                    ]
                }
            }
        },
        "dto.LoginDTO": {
            "type": "object",
            "required": [
                "email",
                "key",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 10
                }
            }
        },
        "dto.SignupDTO": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 4
                },
                "password": {
                    "type": "string",
                    "minLength": 10
                }
            }
        },
        "dto.TagDTO": {
            "type": "object",
            "required": [
                "cover",
                "createdby",
                "description",
                "name"
            ],
            "properties": {
                "cover": {
                    "description": "封面 配图 图片连接地址",
                    "type": "string"
                },
                "createdby": {
                    "description": "创建者",
                    "type": "string"
                },
                "description": {
                    "description": "tag 描述",
                    "type": "string"
                },
                "name": {
                    "description": "tag 名称",
                    "type": "string"
                }
            }
        },
        "dto.TagFollowDTO": {
            "type": "object",
            "required": [
                "tagid"
            ],
            "properties": {
                "tagid": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Cabbage server API",
	Description:      "this is about novel swagger API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
