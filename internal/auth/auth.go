package auth

import (
	"github.com/Abraxas-365/toolkit/pkg/lucia"
	"github.com/gofiber/fiber/v2"
)

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

// Return user id and error
func GetUserId(c *fiber.Ctx) (int, error) {
	session := lucia.GetSession(c)
	userId, err := session.UserIDToInt()
	if err != nil {
		return 0, err
	}

	return userId, nil
}
