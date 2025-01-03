package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `gorm:"size:255;not null" json:"title"` // タスクのタイトル
	Description string    `gorm:"type:text" json:"description"`   // タスクの詳細（オプション）
	Completed   bool      `gorm:"default:false" json:"completed"` // 完了フラグ
	DueDate     time.Time `json:"due_date"`                       // 締め切り日時（オプション）
}
