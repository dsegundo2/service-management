package models

type CreateServiceRequest struct {
	// Unique Title of the service
	Title string `json:"title,omitempty"`
	// description of the service
	Description string `json:"description,omitempty"`
}
