package main

import (
	"dhbw-loerrach.de/dualis/microservice/internal/api/restapi"
	"dhbw-loerrach.de/dualis/microservice/internal/endpoint"
	"dhbw-loerrach.de/dualis/microservice/internal/handle"
	"github.com/go-openapi/loads"
)

//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger generate server --exclude-main --target ../internal/api --spec ../api/definition.yaml

func main() {

	endpoint.LoadService()
	endpoint.CreateDBClient()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	api := handle.InitializeHandlers(swaggerSpec)
	api.UseSwaggerUI()

	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"unix", "http"}
	server.Port = 44444
	server.ConfigureAPI()
	server.Serve()

}
