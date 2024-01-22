package common

import (
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Helpers struct {
	ErrorLogger log.Logger
	InfoLogger  log.Logger
	Validate    *validator.Validate
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
