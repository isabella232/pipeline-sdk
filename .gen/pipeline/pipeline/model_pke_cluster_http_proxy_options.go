/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// PkeClusterHttpProxyOptions struct for PkeClusterHttpProxyOptions
type PkeClusterHttpProxyOptions struct {
	// host of the proxy
	Host string `json:"host"`
	// port the proxy is available on
	Port int32 `json:"port,omitempty"`
	// ID of the secret containing the username and password for the proxy
	SecretId string `json:"secretId,omitempty"`
}
