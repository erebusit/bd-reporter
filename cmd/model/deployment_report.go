package model

type DeploymentStatus string

const (
	Success = "success"
	Error   = "error"
)

type DeploymentReport struct {
	Status      DeploymentStatus `json:"status"`
	Project     string           `json:"project"`
	Version     string           `json:"version,omitempty"`
	Environment string           `json:"environment,omitempty"`
}
