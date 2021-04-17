package handle

import (
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"dhbw-loerrach.de/dualis/microservice/internal/endpoint"
	"github.com/go-openapi/loads"
)

// InitializeHandlers initializes Lambda Handlers
func InitializeHandlers(swaggerSpec *loads.Document) *operations.DualisMicroserviceAPI {
	api := operations.NewDualisMicroserviceAPI(swaggerSpec)
	api.StudentsHandler = operations.StudentsHandlerFunc(endpoint.HandleStudents)

	return api
}
