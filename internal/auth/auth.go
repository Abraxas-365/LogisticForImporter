package auth

type AuthUser struct {
	ID         string
	Email      string
	Name       string
	ProviderID string
	Provider   string
}

func (au *AuthUser) GetID() string {
	return au.ID
}
