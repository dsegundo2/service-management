package models

type ListServicesResponse struct {
	// Total number of objects found not only on this page
	TotalCount int32 `json:"total_count,omitempty"`
	// List of services
	Services []Service `json:"services,omitempty"`
}
