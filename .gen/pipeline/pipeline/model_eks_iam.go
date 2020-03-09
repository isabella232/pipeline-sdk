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
// EksIam struct for EksIam
type EksIam struct {
	// The identifier of existing IAM role to be used for creating the EKS cluster. If not provided a new IAM role is created for the cluster (requires IAM Write Access).
	ClusterRoleId string `json:"clusterRoleId,omitempty"`
	// The identifier of existing IAM role to be used for creating the EKS nodes. If not provided a new IAM role is created for the nodes (requires IAM Write Access).
	NodeInstanceRoleId string `json:"nodeInstanceRoleId,omitempty"`
	// Marks if the userid associated with the clusters AWS secret has to be used in kubeconfig (bypasses IAM user creation).
	DefaultUser bool `json:"defaultUser,omitempty"`
}
