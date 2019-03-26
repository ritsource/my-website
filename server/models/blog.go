package models

// Blog - Blog model type
type Blog struct {
	Title string `json:"title"`
	Description string `json:"description"`
	HTML string `json:"html"`
	Markdown string `json:"markdown"`
	ImageURL string `json:"image_url"`
	IsPublic bool `json:"is_public"`
	IsDeleted bool `json:"is_deleted"`
}