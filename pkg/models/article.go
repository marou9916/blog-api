package models

type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}
