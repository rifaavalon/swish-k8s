package utils

import (
	"go-app/models"
)

func ValidateEnvironment(env models.Environment) string {
	var errors string

	if env.Name == "" {
		errors += "environmentName is required. "
	}
	if env.BaseImage == "" {
		errors += "baseImage is required. "
	}
	if len(env.Packages) == 0 {
		errors += "At least one package is required. "
	}
	if env.CPURequest == "" {
		errors += "cpuRequest is required. "
	}
	if env.MemoryRequest == "" {
		errors += "memoryRequest is required. "
	}

	return errors
}
