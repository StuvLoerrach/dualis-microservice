package internal

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// DualisClaims extends StandardClaims
type DualisClaims struct {
	StudentID      int                `json:"studentid,omitempty"`
	StandardClaims jwt.StandardClaims `json:"standardclaims,omitempty"`
}

// Valid checks if token is valid
func (c *DualisClaims) Valid() error {

	err := c.StandardClaims.Valid()

	if err != nil {
		return fmt.Errorf("Invalid token %v", err)
	}

	return nil
}
