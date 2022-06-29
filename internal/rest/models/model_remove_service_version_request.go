package models

type RemoveServiceVersionRequest struct {
	// version number of this service version. must follow number.number.number format
	Version string `json:"version,omitempty"`
	// parent id that the service version belongs to
	ServiceId string `json:"service_id,omitempty"`
}
