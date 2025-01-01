package services

import (
	"fmt"
	"os/exec"
	"strings"
	"go-app/models"
)

func DeployEnvironment(env models.Environment) error {
	packages := strings.Join(env.Packages, ",")
	helmCommand := []string{
		"helm", "install", env.Name, "./helm-chart",
		"--set", fmt.Sprintf("environmentName=%s", env.Name),
		"--set", fmt.Sprintf("baseImage=%s", env.BaseImage),
		"--set", fmt.Sprintf("packages=%s", packages),
		"--set", fmt.Sprintf("cpuRequest=%s", env.CPURequest),
		"--set", fmt.Sprintf("memoryRequest=%s", env.MemoryRequest),
	}
	if env.GPURequest != "" {
		helmCommand = append(helmCommand, "--set", fmt.Sprintf("gpuRequest=%s", env.GPURequest))
	}

	cmd := exec.Command(helmCommand[0], helmCommand[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Helm command failed: %s", string(output))
	}

	return nil
}
