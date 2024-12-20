package userapi

import (
	"github.com/Abraxas-365/cabo/internal/auth"
	"github.com/Abraxas-365/cabo/internal/user"
	"github.com/Abraxas-365/cabo/internal/user/usersrv"
	"github.com/Abraxas-365/toolkit/pkg/errors"
	"github.com/Abraxas-365/toolkit/pkg/lucia"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for user endpoints
func SetupRoutes(app *fiber.App, service *usersrv.Service, authMiddleware *lucia.AuthMiddleware[*auth.AuthUser]) {
	// Create a new user
	app.Post("/users", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {

		newUserInput := struct {
			DocumentType   user.DocumentType `json:"document_type"`       // Document type (DNI, CE, Passport)
			DocumentNumber string            `json:"document_number"`     // The user's document number
			FirstName      string            `json:"first_name"`          // First name
			LastName       string            `json:"last_name,omitempty"` // Last name (optional)
			Phone          string            `json:"phone,omitempty"`     // Contact phone number
			Email          string            `json:"email,omitempty"`     // Contact email
		}{}

		if err := c.BodyParser(&newUserInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}

		newUser := user.User{
			DocumentType:   newUserInput.DocumentType,
			DocumentNumber: newUserInput.DocumentNumber,
			FirstName:      newUserInput.FirstName,
			LastName:       newUserInput.LastName,
			Phone:          newUserInput.Phone,
			Email:          newUserInput.Email,
		}

		if err := newUser.Validate(); err != nil {
			return err
		}

		if err := service.CreateUser(c.Context(), &newUser); err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(newUser)
	})

	// Get a user by ID
	app.Get("/users/:id", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}

		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}
		if sessionUserId != id {
			return errors.ErrForbidden("User not allowed to access this resource")
		}

		user, err := service.GetUser(c.Context(), id)
		if err != nil {
			if errors.IsNotFound(err) {
				return errors.ErrNotFound("User not found")
			}
			return err
		}

		return c.JSON(user)
	})

	// Update an existing user
	app.Put("/users/:id", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}
		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		if sessionUserId != id {
			return errors.ErrForbidden("User not allowed to access this resource")
		}
		updateUserInput := struct {
			DocumentType   user.DocumentType `json:"document_type"`       // Document type (DNI, CE, Passport)
			DocumentNumber string            `json:"document_number"`     // The user's document number
			FirstName      string            `json:"first_name"`          // First name
			LastName       string            `json:"last_name,omitempty"` // Last name (optional)
			Phone          string            `json:"phone,omitempty"`     // Contact phone number
			Email          string            `json:"email,omitempty"`     // Contact email
		}{}

		if err := c.BodyParser(&updateUserInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}

		updatedUser := user.User{
			ID:             id,
			DocumentType:   updateUserInput.DocumentType,
			DocumentNumber: updateUserInput.DocumentNumber,
			FirstName:      updateUserInput.FirstName,
			LastName:       updateUserInput.LastName,
			Phone:          updateUserInput.Phone,
			Email:          updateUserInput.Email,
		}

		if err := updatedUser.Validate(); err != nil {
			return err
		}

		uu, err := service.UpdateUser(c.Context(), &updatedUser)
		if err != nil {
			return err
		}

		return c.JSON(uu)
	})

	// Add a new direction for a user
	app.Post("/users/:id/direction", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}

		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}
		if sessionUserId != id {
			return errors.ErrForbidden("User not allowed to access this resource")
		}

		newDirectionInput := new(user.Direction)
		if err := c.BodyParser(newDirectionInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}
		newDirectionInput.UserID = id

		if err := newDirectionInput.Validate(); err != nil {
			return err
		}

		direction, err := service.SaveDirection(c.Context(), newDirectionInput)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(direction)
	})

	// Add a new direction for a user
	app.Post("/users/:id/direction", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}

		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}
		if sessionUserId != id {
			return errors.ErrForbidden("User not allowed to access this resource")
		}

		// Define the input structure for creating a new direction
		newDirectionInput := struct {
			Country      string `json:"country" validate:"required"`       // Country is required
			State        string `json:"state" validate:"required"`         // State is required
			City         string `json:"city" validate:"required"`          // City is required
			AddressLine1 string `json:"address_line1" validate:"required"` // Address Line 1 is required
			AddressLine2 string `json:"address_line2,omitempty"`           // Address Line 2 (optional)
			PostalCode   string `json:"postal_code" validate:"required"`   // Postal code is required
		}{}

		if err := c.BodyParser(&newDirectionInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}

		// Map input to the Direction model
		newDirection := user.Direction{
			UserID:       id,
			Country:      newDirectionInput.Country,
			State:        newDirectionInput.State,
			City:         newDirectionInput.City,
			AddressLine1: newDirectionInput.AddressLine1,
			AddressLine2: newDirectionInput.AddressLine2,
			PostalCode:   newDirectionInput.PostalCode,
		}

		// Validate the new direction
		if err := newDirection.Validate(); err != nil {
			return err
		}

		// Save the direction
		direction, err := service.SaveDirection(c.Context(), &newDirection)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(direction)
	})

	// Update a user's direction
	app.Put("/users/:id/direction/:directionId", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}

		directionId, err := c.ParamsInt("directionId")
		if err != nil {
			return errors.ErrBadRequest("Invalid direction ID")
		}

		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		if sessionUserId != id {
			return errors.ErrForbidden("User not allowed to access this resource")
		}

		// Define the input structure for updating an existing direction
		updateDirectionInput := struct {
			Country      string `json:"country" validate:"required"`       // Country is required
			State        string `json:"state" validate:"required"`         // State is required
			City         string `json:"city" validate:"required"`          // City is required
			AddressLine1 string `json:"address_line1" validate:"required"` // Address Line 1 is required
			AddressLine2 string `json:"address_line2,omitempty"`           // Address Line 2 (optional)
			PostalCode   string `json:"postal_code" validate:"required"`   // Postal code is required
		}{}

		if err := c.BodyParser(&updateDirectionInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}

		// Map input to the Direction model
		updatedDirection := user.Direction{
			UserID:       sessionUserId,
			ID:           directionId,
			Country:      updateDirectionInput.Country,
			State:        updateDirectionInput.State,
			City:         updateDirectionInput.City,
			AddressLine1: updateDirectionInput.AddressLine1,
			AddressLine2: updateDirectionInput.AddressLine2,
			PostalCode:   updateDirectionInput.PostalCode,
		}

		// Validate the updated direction
		if err := updatedDirection.Validate(); err != nil {
			return err
		}

		// Update the direction
		direction, err := service.UpdateDirection(c.Context(), &updatedDirection)
		if err != nil {
			return err
		}

		return c.JSON(direction)
	})
}
