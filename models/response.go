package models

type CreateServiceResponse struct {
	URL string `json:"url"`
}

type ServiceResponse struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Status string `json:"status"`
}
