package consignee

import (
	"fmt"
	"strings"

	"github.com/Abraxas-365/toolkit/pkg/errors"
)

type DocumentType string

// Consignee represents a consignatario (can be associated with a user)
type Consignee struct {
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

func NewDocumentType(docType string) (DocumentType, error) {
	switch DocumentType(strings.ToUpper(docType)) {
	case RUC10:
		return RUC10, nil
	case RUC20:
		return RUC20, nil
	case DNI:
		return DNI, nil
	case CI:
		return CI, nil
	case Passport:
		return Passport, nil
	default:
		return "", errors.ErrBadRequest("invalid DocumentType")
	}
}

// IsValid checks if the Consignatario's DocumentType is valid
func (dt DocumentType) IsValid() error {
	switch dt {
	case RUC10, RUC20, DNI, CI, Passport:
		return nil
	}
	return errors.ErrBadRequest(fmt.Sprintf("invalid DocumentType: %s", dt))
}

func (c *Consignee) IsOfUser(userId int) bool {
	return c.UserID == userId
}
