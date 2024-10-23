package user

import (
	"fmt"
	"strings"

	"github.com/Abraxas-365/toolkit/pkg/errors"
)

type DocumentType string

const (
	DNI      DocumentType = "DNI"
	CE       DocumentType = "CE"
	Passport DocumentType = "Passport"
)

func NewDocumentType(docType string) (DocumentType, error) {
	switch DocumentType(strings.ToUpper(docType)) {
	case DNI:
		return DNI, nil
	case CE:
		return CE, nil
	case Passport:
		return Passport, nil
	default:
		return "", errors.ErrBadRequest("Invalid DocumentType")
	}
}

type User struct {
	ID             int          `json:"id"`
	IsComplete     bool         `json:"is_complete"`
	DocumentType   DocumentType `json:"document_type" validate:"required,oneof=DNI CE Passport"`
	DocumentNumber string       `json:"document_number" validate:"required"`
	FirstName      string       `json:"first_name" validate:"required"`
	LastName       string       `json:"last_name,omitempty" validate:"omitempty"`
	Email          string       `json:"email,omitempty" validate:"omitempty,email"`
	Phone          string       `json:"phone,omitempty" validate:"omitempty"`
}

func (dt DocumentType) IsValid() error {
	switch dt {
	case DNI, CE, Passport:
		return nil
	}
	return errors.ErrBadRequest(fmt.Sprintf("Invalid DocumentType: %s", dt))
}

func (u *User) Validate() error {
	if u.FirstName == "" {
		return errors.ErrBadRequest("FirstName is required")
	}

	if u.LastName == "" {
		return errors.ErrBadRequest("LastName is required")
	}

	if err := u.DocumentType.IsValid(); err != nil {
		return err // DocumentType validation error
	}

	if u.DocumentNumber == "" {
		return errors.ErrBadRequest("DocumentNumber is required")
	}

	return nil // Validation successful
}
