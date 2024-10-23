package consignatario

import (
	"github.com/Abraxas-365/cabo/internal/auth"
	"github.com/Abraxas-365/toolkit/pkg/errors"
	"github.com/Abraxas-365/toolkit/pkg/lucia"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for consignatario endpoints
func SetupRoutes(app *fiber.App, service *Service, authMiddleware *lucia.AuthMiddleware[*auth.AuthUser]) {
	// Create a new consignatario
	app.Post("/consignatarios", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		session := lucia.GetSession(c)
		userId, err := session.UserIDToInt()
		if err != nil {
			return err
		}
		newConsignatarioInput := struct {
			DocumentType   DocumentType `json:"document_type"`        // Document type (RUC 10, RUC 20, DNI, CI, Passport)
			DocumentNumber string       `json:"document_number"`      // The consignatario's document number
			FirstName      string       `json:"first_name,omitempty"` // First name (for individual consignatarios)
			LastName       *string      `json:"last_name,omitempty"`  // Last name (optional, for individual consignatarios)
			Phone          string       `json:"phone,omitempty"`      // Contact phone number
			Email          string       `json:"email,omitempty"`      // Contact email
		}{}
		if err := c.BodyParser(&newConsignatarioInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}

		newConsignatario := Consignatario{
			DocumentType:   newConsignatarioInput.DocumentType,
			DocumentNumber: newConsignatarioInput.DocumentNumber,
			FirstName:      newConsignatarioInput.FirstName,
			LastName:       newConsignatarioInput.LastName,
			Phone:          newConsignatarioInput.Phone,
			Email:          newConsignatarioInput.Email,
			UserID:         userId,
		}

		nc, err := service.CreateConsignatario(c.Context(), &newConsignatario)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(nc)
	})

	// Get a consignatario by ID
	app.Get("/consignatarios/:id", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		session := lucia.GetSession(c)
		userId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}

		consignatario, err := service.GetConsignatario(c.Context(), id)
		if err != nil {
			return err
		}

		if !consignatario.IsOfUser(userId) {
			return errors.ErrUnauthorized("You are not authorized to view this consignatario")
		}

		return c.JSON(consignatario)
	})

	// Update an existing consignatario
	app.Put("/consignatarios/:id", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		session := lucia.GetSession(c)
		userId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}

		updateConsignatarioInput := struct {
			DocumentType   DocumentType `json:"document_type"`        // Document type (RUC 10, RUC 20, DNI, CI, Passport)
			DocumentNumber string       `json:"document_number"`      // The consignatario's document number
			FirstName      string       `json:"first_name,omitempty"` // First name (for individual consignatarios)
			LastName       *string      `json:"last_name,omitempty"`  // Last name (optional, for individual consignatarios)
			Phone          string       `json:"phone,omitempty"`      // Contact phone number
			Email          string       `json:"email,omitempty"`      // Contact email
		}{}

		if err := c.BodyParser(&updateConsignatarioInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}

		updatedConsignatario := Consignatario{
			ID:             id,
			DocumentType:   updateConsignatarioInput.DocumentType,
			DocumentNumber: updateConsignatarioInput.DocumentNumber,
			FirstName:      updateConsignatarioInput.FirstName,
			LastName:       updateConsignatarioInput.LastName,
			Phone:          updateConsignatarioInput.Phone,
			Email:          updateConsignatarioInput.Email,
			UserID:         userId,
		}

		uc, err := service.UpdateConsignatario(c.Context(), &updatedConsignatario)
		if err != nil {
			return err
		}

		return c.JSON(uc)
	})
	// Get all consignatarios by user ID with pagination
	app.Get("/users/:userId/consignatarios", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		userId, err := c.ParamsInt("userId")
		if err != nil {
			return errors.ErrBadRequest("Invalid User ID")
		}

		pageNumber := c.QueryInt("page", 1) // Default to page 1
		pageSize := c.QueryInt("size", 10)  // Default to 10 items per page

		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		if userId != sessionUserId {
			return errors.ErrUnauthorized("You are not authorized to view this user's consignatarios")
		}

		paginatedConsignatarios, err := service.GetAllByUserId(c.Context(), userId, pageNumber, pageSize)
		if err != nil {
			return err
		}

		return c.JSON(paginatedConsignatarios)
	})
}
