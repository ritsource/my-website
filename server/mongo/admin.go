package mongo

// Admin ...
type Admin struct {
	ID       string `json:"_id"`
	Email    string `json:"email"`
	GoogleID string `json:"googleid"`
}

// AdminServiceInterface ...
type AdminServiceInterface interface {
	CreateAdmin(a *Admin) error
	GetAdmin(a *Admin) (*Admin, error)
}
