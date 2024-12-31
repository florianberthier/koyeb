package models

type CreateServiceRequest struct {
	Script bool   `json:"script"`
	URL    string `json:"url" validate:"required"`
}
