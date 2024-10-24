package productapi

import (
	"github.com/Abraxas-365/cabo/internal/auth"
	"github.com/Abraxas-365/cabo/internal/product"
	"github.com/Abraxas-365/cabo/internal/product/productsrv"
	"github.com/Abraxas-365/toolkit/pkg/errors"
	"github.com/Abraxas-365/toolkit/pkg/lucia"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for product endpoints
func SetupRoutes(app *fiber.App, service *productsrv.Service, authMiddleware *lucia.AuthMiddleware[*auth.AuthUser]) {
	// Create a new product
	app.Post("/products", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {

		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		newProductInput := struct {
			WarehouseID int            `json:"warehouse_id"`
			Status      product.Status `json:"status"`
			Invoices    []string       `json:"invoices"` // List of S3 URLs
		}{}

		if err := c.BodyParser(&newProductInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}

		newProduct := product.Product{
			UserID:      sessionUserId,
			WarehouseID: newProductInput.WarehouseID,
			Status:      newProductInput.Status,
			Invoices:    newProductInput.Invoices,
		}

		savedProduct, err := service.SaveProduct(c.Context(), &newProduct)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(savedProduct)
	})

	// Get a product by ID
	app.Get("/products/:id", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}

		product, err := service.GetProductById(c.Context(), id)
		if err != nil {
			if errors.IsNotFound(err) {
				return errors.ErrNotFound("Product not found")
			}
			return err
		}

		if product.UserID != sessionUserId {
			return errors.ErrForbidden("User not allowed to access this resource")
		}

		return c.JSON(product)
	})

	// Update an existing product
	app.Put("/products/:id", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ErrBadRequest("Invalid ID")
		}

		updateProductInput := struct {
			WarehouseID int            `json:"warehouse_id"`
			Status      product.Status `json:"status"`
			Invoices    []string       `json:"invoices"` // List of S3 URLs
		}{}

		if err := c.BodyParser(&updateProductInput); err != nil {
			return errors.ErrBadRequest("Invalid request body")
		}

		updatedProduct := product.Product{
			ID:          id,
			UserID:      sessionUserId,
			WarehouseID: updateProductInput.WarehouseID,
			Status:      updateProductInput.Status,
			Invoices:    updateProductInput.Invoices,
		}

		updatedProd, err := service.UpdateProduct(c.Context(), &updatedProduct)
		if err != nil {
			return err
		}

		return c.JSON(updatedProd)
	})

	// Generate presigned URL for uploading an invoice
	app.Get("/products/presigned-url", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		url, err := service.GeneratePresignedPutURL(c.Context(), sessionUserId)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"url": url})
	})

	// Get all products for a user
	app.Get("/users/:id/products", authMiddleware.RequireAuth(), func(c *fiber.Ctx) error {
		session := lucia.GetSession(c)
		sessionUserId, err := session.UserIDToInt()
		if err != nil {
			return err
		}

		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("page_size", 10)

		products, err := service.GetAllUserProducts(c.Context(), sessionUserId, page, pageSize)
		if err != nil {
			return err
		}

		return c.JSON(products)
	})
}
