package models

type UpdateServiceRequest struct {
	// id of the service to update
	ServiceId string `json:"service_id,omitempty"`
	// title for the service to update
	Title string `json:"title,omitempty"`
	// new description of the service to update
	Description string `json:"description,omitempty"`
}
