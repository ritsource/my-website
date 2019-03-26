package models

// Files - Files model type
type Files struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Type string `json:"type"`
	FileURL string `json:"file_url"`
	IsPublic string `json:"is_public"`
	IsDeleted string `json:"is_deleted"`
}