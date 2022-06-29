package models

type AddServiceVersionRequest struct {
	// version number of this service version. must follow number.number.number format
	Version string `json:"version,omitempty"`
	// proof of concept random field on a version
	ServiceExampleField string `json:"service_example_field,omitempty"`
	// parent id that the service version belongs to
	ServiceId string `json:"service_id,omitempty"`
}
