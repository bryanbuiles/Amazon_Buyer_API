package gateway

import (
	"github.com/bryanbuiles/tecnical_interview/internal/database"
)

// Services ...
type Services struct {
	data AllDataGateway
}

// WebServices .for the users
type WebServices struct {
	Services
}

// NewServices ...
func NewServices() Services {
	client := database.NewClient()
	return Services{
		data: &DataBaseService{DB: client},
	}
}

// Start a new user service
func Start() *WebServices {
	return &WebServices{NewServices()}
}
