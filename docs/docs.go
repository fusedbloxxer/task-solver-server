// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "url": "https://github.com/fusedbloxxer"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/config": {
            "get": {
                "description": "Gets the app settings for the environment the server is running in.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Get the full configuration file for the server.",
                "responses": {
                    "200": {
                        "description": "The configuration file is returned.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Fetch all the stored task results from the server. They are unordered and unfiltered.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Retrieve all stored task results",
                "responses": {
                    "200": {
                        "description": "All task results are returned as an array",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TaskResult"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch the tasks",
                        "schema": {
                            "$ref": "#/definitions/model.BadRequestError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes the task results from the server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Delete all stored task results",
                "responses": {
                    "200": {
                        "description": "All tasks are deleted"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.BadRequestError"
                        }
                    }
                }
            }
        },
        "/tasks/:taskId": {
            "get": {
                "description": "Fetch the saved result from the server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Get a saved task result by its document id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Used to identify the task",
                        "name": "taskId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The task result is returned",
                        "schema": {
                            "$ref": "#/definitions/model.TaskResult"
                        }
                    },
                    "400": {
                        "description": "The task does not exist",
                        "schema": {
                            "$ref": "#/definitions/model.BadRequestError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a saved task result using the id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Delete a task result using its id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Used to identify the task",
                        "name": "taskId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The task is removed"
                    },
                    "400": {
                        "description": "The taskId does not exist",
                        "schema": {
                            "$ref": "#/definitions/model.BadRequestError"
                        }
                    }
                }
            }
        },
        "/tasks/indexes": {
            "get": {
                "description": "Fetch from the server the possible problem types implemented.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Retrieve the possible problem types or indexes",
                "responses": {
                    "200": {
                        "description": "The problem indexes are returned as an array. It is unordered.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to fetch the problem indexes.",
                        "schema": {
                            "$ref": "#/definitions/model.BadRequestError"
                        }
                    }
                }
            }
        },
        "/tasks/solve": {
            "post": {
                "description": "Solve a task by using the context and the index of the problem. Save the results.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "task"
                ],
                "summary": "Solve a task and save the result",
                "parameters": [
                    {
                        "description": "The task to be solved. Its index must be obtained from /tasks/indexes.",
                        "name": "Task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The task result containing an id for the saved value and the answer",
                        "schema": {
                            "$ref": "#/definitions/model.TaskResult"
                        }
                    },
                    "400": {
                        "description": "The task model is invalid",
                        "schema": {
                            "$ref": "#/definitions/model.BadRequestError"
                        }
                    }
                }
            }
        },
        "/test": {
            "get": {
                "description": "Tests if the API is working. A \"Hello, World!\" message should always be returned.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test"
                ],
                "summary": "Test that the API is responding",
                "responses": {
                    "200": {
                        "description": "The message \"Hello, World!\" is returned",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BadRequestError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "required": [
                "context",
                "index"
            ],
            "properties": {
                "context": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                },
                "index": {
                    "type": "integer"
                }
            }
        },
        "model.TaskResult": {
            "type": "object",
            "required": [
                "answer",
                "id",
                "task"
            ],
            "properties": {
                "answer": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "task": {
                    "$ref": "#/definitions/model.Task"
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
	Host:        "127.0.0.1:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Tasks API",
	Description: "This API can be used to solve tasks and save the results to firebase",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
