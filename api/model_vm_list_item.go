/*
 * Promethium Daemon API
 *
 * API for Promethium Daemon
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type VmListItem struct {

	Id string `json:"id,omitempty"`

	// VM status
	Status string `json:"status,omitempty"`

	Name string `json:"name,omitempty"`
}
