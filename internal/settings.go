package internal

// Settings for dualis-microservice
type Settings struct {
	TokenSecret string `json:"token_secret"`

	DbHost     string `json:"db_host"`
	DbUsername string `json:"db_username"`
	DbPassword string `json:"db_password"`
	DbName     string `json:"db_name"`
}
