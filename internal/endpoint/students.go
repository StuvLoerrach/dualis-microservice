package endpoint

import (
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// HandleStudents finds students by various criteria
func HandleStudents(params operations.StudentsParams, principal interface{}) middleware.Responder {

}
