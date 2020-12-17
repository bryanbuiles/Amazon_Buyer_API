package gateway

// Services ...
type Services struct {
	buyers BuyerGateway
}

// WebServices .for the users
type WebServices struct {
	Services
}

// NewServices ...
func NewServices() Services {
	return Services{
		buyers: &BuyerService{},
	}
}

// Start a new user service
func Start() *WebServices {
	return &WebServices{NewServices()}
}
