package gateway

// Services ...
type Services struct {
	buyers BuyerGateway
}

// WebServices .for the users
type WebServices struct {
	Services
}

// Start a new user service
func Start() *WebServices {
	return &WebServices{}
}
