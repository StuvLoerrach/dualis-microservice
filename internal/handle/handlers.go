package handle

import (
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi/operations"
	"dhbw-loerrach.de/dualis/microservice/internal/endpoint"
	"github.com/go-openapi/loads"
)

// InitializeHandlers initializes Lambda Handlers
func InitializeHandlers(swaggerSpec *loads.Document) *operations.DualisMicroserviceAPI {
	api := operations.NewDualisMicroserviceAPI(swaggerSpec)
	api.DualisKeyAuth = endpoint.VerifyToken
	api.LoginHandler = operations.LoginHandlerFunc(endpoint.HandleLogin)
	api.StudentsHandler = operations.StudentsHandlerFunc(endpoint.HandleStudents)
	api.StudentPerformanceHandler = operations.StudentPerformanceHandlerFunc(endpoint.HandleStudentPerformance)
	api.StudentModuleStatisticsHandler = operations.StudentModuleStatisticsHandlerFunc(endpoint.HandleStudentModuleStatistics)

	return api
}
