package models

// Admin - admin (user) model type
type Admin struct {
	// ID       string `json:"_id"`
	Email    string `json:"email"`
	GoogleID string `json:"googleid"`
}