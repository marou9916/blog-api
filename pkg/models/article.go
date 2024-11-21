package models

type Article struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}
