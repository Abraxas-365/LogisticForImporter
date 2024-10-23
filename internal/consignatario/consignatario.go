package consignatario

import (
	"fmt"
	"github.com/Abraxas-365/toolkit/pkg/errors"
)

type DocumentType string

// Consignatario represents a consignatario (can be associated with a user)
type Consignatario struct {
	ID             int          `json:"id"`
	UserID         int          `json:"user_id"`              // Reference to the associated user
	DocumentType   DocumentType `json:"document_type"`        // Document type (RUC 10, RUC 20, DNI, CI, Passport)
	DocumentNumber string       `json:"document_number"`      // The consignatario's document number
	FirstName      string       `json:"first_name,omitempty"` // First name (for individual consignatarios)
	LastName       *string      `json:"last_name,omitempty"`  // Last name (optional, for individual consignatarios)
	Phone          string       `json:"phone,omitempty"`      // Contact phone number
	Email          string       `json:"email,omitempty"`      // Contact email
}

// ValidDocumentTypes for Consignatario
const (
	RUC10    DocumentType = "RUC10"
	RUC20    DocumentType = "RUC20"
	DNI      DocumentType = "DNI"
	CI       DocumentType = "CI"
	Passport DocumentType = "Passport"
)

// IsValid checks if the Consignatario's DocumentType is valid
func (dt DocumentType) IsValid() error {
	switch dt {
	case RUC10, RUC20, DNI, CI, Passport:
		return nil
	}
	return errors.ErrBadRequest(fmt.Sprintf("invalid DocumentType: %s", dt))
}

func (c *Consignatario) IsOfUser(userId int) bool {
	return c.UserID == userId
}
