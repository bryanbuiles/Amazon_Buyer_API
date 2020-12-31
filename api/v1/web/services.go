package v1

import (
	v1 "github.com/bryanbuiles/tecnical_interview/api/v1/gateway"
	"github.com/bryanbuiles/tecnical_interview/internal/database"
)

// Services ...
type Services struct {
	data v1.AllDataGateway
}

// WebServices .for the users
type WebServices struct {
	Services
}

// NewServices ...
func NewServices() Services {
	client := database.NewClient()
	return Services{
		data: &v1.DataBaseService{DB: client},
	}
}

// Start a new user service
func Start() *WebServices {
	return &WebServices{NewServices()}
}
