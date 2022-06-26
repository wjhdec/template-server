package model

import "github.com/wjhdec/template-server/pkg/dbgorm"

type Book struct {
	dbgorm.UUIDField
	Name           string `json:"name,omitempty"`
	Classification string `json:"classification,omitempty"`
	Describe       string `json:"describe,omitempty"`
}
