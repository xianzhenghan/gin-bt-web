// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCode = "code"

// Code mapped from table <code>
type Code struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID      int64     `gorm:"column:user_id;not null" json:"user_id"`
	Username    string    `gorm:"column:username;not null;comment:用户名" json:"username"` // 用户名
	Title       string    `gorm:"column:title;not null" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CodeContent string    `gorm:"column:code_content;not null" json:"code_content"`
	Language    string    `gorm:"column:language" json:"language"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	IsPublic    bool      `gorm:"column:is_public" json:"is_public"`
	Tags        string    `gorm:"column:tags" json:"tags"`
	Version     string    `gorm:"column:version" json:"version"`
	FileName    string    `gorm:"column:file_name" json:"file_name"`
	LineNumbers bool      `gorm:"column:line_numbers;default:1" json:"line_numbers"`
}

// TableName Code's table name
func (*Code) TableName() string {
	return TableNameCode
}
