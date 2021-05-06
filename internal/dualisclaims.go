package internal

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// DualisClaims extends StandardClaims
type DualisClaims struct {
	StudentID      int64              `json:"studentid,omitempty"`
	StudentCourse  int64              `json:"studentcourse,omitempty"`
	Organization   int64              `json:"organization,omitempty"`
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
