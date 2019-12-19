/*
 * Promethium Daemon API
 *
 * API for Promethium Daemon
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Vm struct {
	Id string `json:"id,omitempty"`
	// VM status
	Status string `json:"status,omitempty"`
	Name string `json:"name,omitempty"`
	Cpus int64 `json:"cpus,omitempty"`
	Memory int64 `json:"memory,omitempty"`
	Volumes []VmVolume `json:"volumes,omitempty"`
	Drives []VmDrive `json:"drives,omitempty"`
	Interfaces []VmInterface `json:"interfaces,omitempty"`
}
