/*
 * Service Managemenet
 *
 * This is a sample proof of concept repo meant to be run on localhost at the moment
 *
 * API version: 1.0.0
 * Contact: diegosegundo2@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RemoveServiceVersionRequest struct {
	// version number of this service version. must follow number.number.number format
	Version string `json:"version,omitempty"`
	// parent id that the service version belongs to
	ServiceId string `json:"service_id,omitempty"`
}
