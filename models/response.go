package models

type CreateServiceResponse struct {
	URL string `json:"url"`
}

type ServiceResponse struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
