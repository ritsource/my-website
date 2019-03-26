package models

// Blog - Blog model type
type Blog struct {
	Title string `json:"title"`
	Description string `json:"description"`
	HTML string `json:"html"`
	Markdown string `json:"markdown"`
	ImageURL string `json:"image_url"`
	IsPublic string `json:"is_public"`
	IsDeleted string `json:"is_deleted"`
}