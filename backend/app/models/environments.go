package models

type Environment struct {
	Name        string   `json:"environmentName"`
	BaseImage   string   `json:"baseImage"`
	Packages    []string `json:"packages"`
	CPURequest  string   `json:"cpuRequest"`
	MemoryRequest string `json:"memoryRequest"`
	GPURequest  string   `json:"gpuRequest,omitempty"`
}
