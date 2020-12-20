package gateway

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
	return Services{
		data: &DataBaseService{},
	}
}

// Start a new user service
func Start() *WebServices {
	return &WebServices{NewServices()}
}
