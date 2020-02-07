package requests

import (
	"bookshelf_service/src/domains"
	"strings"
)

type AuthorRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (a *AuthorRequest) CleanAndValidate() error {
	a.LastName = strings.TrimSpace(a.LastName)
	a.FirstName = strings.TrimSpace(a.FirstName)
	if a.LastName == "" {
		return domains.NewValidationError("last_name can not be empty")
	} else if a.FirstName == "" {
		return domains.NewValidationError("first_name can not be empty")
	}
	return nil
}
