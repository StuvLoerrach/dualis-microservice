package endpoint

import (
	"fmt"
	"log"
	"time"

	"dhbw-loerrach.de/dualis/microservice/internal"
	"dhbw-loerrach.de/dualis/microservice/internal/api/models"
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
)

// HandleLogin takes email and password, if they can be verified a new dualis token is created
func HandleLogin(params operations.LoginParams) middleware.Responder {

	var (
		id       int64
		email    string
		password string
		courseFk int64

		hadNext bool
	)

	rows, err := db.Query("SELECT * FROM dualis.student WHERE email = ?", *params.Login.Email)

	if err != nil {
		errMsg := err.Error()
		return operations.NewLoginInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	defer rows.Close()

	hadNext = false

	for rows.Next() {
		err = rows.Scan(&id, &email, &password, &courseFk)

		if err != nil {
			errMsg := err.Error()
			return operations.NewLoginInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
		}

		hadNext = true
	}

	if !hadNext {
		errMsg := "Couldn't safely identify this student!"
		return operations.NewLoginInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	if password != *params.Login.Password {
		errMsg := fmt.Sprintf("Couldn't get authenticated!")
		return operations.NewLoginInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	dualisToken, err := generateDualisToken(id, email, password, courseFk)

	if err != nil {
		errMsg := fmt.Sprintf("Couldn't generate Dualis token: %v", err)
		return operations.NewLoginInternalServerError().WithPayload(&models.SimpleError{Error: &errMsg})
	}

	dualisTokenResp := &models.LoginResponse{Jwt: &dualisToken}

	return operations.NewLoginOK().WithPayload(dualisTokenResp)

}

func generateDualisToken(studentId int64, email string, password string, courseFk int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &internal.DualisClaims{
		StudentID:     studentId,
		StudentCourse: courseFk,
		StandardClaims: jwt.StandardClaims{
			Subject:   email,
			Issuer:    "dualis-microservice",
			ExpiresAt: time.Now().Unix() + 600,
		},
	})

	signedToken, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", fmt.Errorf("Signing failed: %v", err)
	}

	return signedToken, nil
}

// VerifyToken verifies the Dualis token
func VerifyToken(rawToken string) (interface{}, error) {

	token, err := jwt.ParseWithClaims(rawToken, &internal.DualisClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return tokenSecret, nil
	})

	if err != nil {
		log.Printf("Token not parseable: %v", err)
		return nil, fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(*internal.DualisClaims)

	if !ok || !token.Valid {
		log.Printf("Invalid token: %v", err)
		return nil, fmt.Errorf("Invalid token")
	}

	err = claims.Valid()

	if err != nil {
		log.Printf("Invalid token claims: %v", err)
		return nil, fmt.Errorf("Invalid token")
	}

	if ok && token.Valid && err == nil {
		return token.Claims, nil
	}

	return nil, fmt.Errorf("Couldn't validate Dualis token")
}
