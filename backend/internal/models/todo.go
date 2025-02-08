package models

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `gorm:"size:255;not null" json:"title"` // タスクのタイトル
	Description string    `gorm:"type:text" json:"description"`   // タスクの詳細（オプション）
	Completed   bool      `gorm:"default:false" json:"completed"` // 完了フラグ
	DueDate     time.Time `json:"dueDate"`                        // 締め切り日時（オプション）
	CreatedAt   int64     `json:"createdAt"`
	UpdatedAt   int64     `json:"updatedAt"`
}
