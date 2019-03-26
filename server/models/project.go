package models

// Project - Project model type
type Project struct {
	Title string `json:"title"`
	Description string `json:"description"`
	HTML string `json:"html"`
	Markdown string `json:"markdown"`
	Link string `json:"link"`
	ImageURL string `json:"image_url"`
	IsPublic string `json:"is_public"`
	IsDeleted string `json:"is_deleted"`
}