# Dynamic Environment Helm Chart

This repository contains a Helm chart to deploy dynamic environments for developers on any Kubernetes cluster. The setup includes a Go application, an SSH server, monitoring with Prometheus and Grafana, and autoscaling.

## Prerequisites

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Helm](https://helm.sh/docs/intro/install/)
- [Docker](https://docs.docker.com/get-docker/)


## Automated Setup

The setup process is automated using GitHub Actions. Simply push your changes to the `main` branch, and the workflow will handle the rest.

## Manual Setup

If you prefer to set up the environment manually, follow these steps:

### 1. Create a Kubernetes Cluster

Ensure you have a Kubernetes cluster running. You can use any Kubernetes provider (e.g., Minikube, kind, EKS, AKS, etc.).

### 2. Configure kubectl to Use Your Kubernetes Cluster

Ensure `kubectl` is configured to interact with your Kubernetes cluster.

### 3. Install Helm

```sh
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash
```
### 4. Add Helm Repositories
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
```

### 5. Deploy Prometheus and Grafana
Run the `setup-prometheus.sh` script to deploy Prometheus and Grafana:
```
./scripts/setup-prometheus.sh
```
### 6. Build and Push Docker Image
Build and push the docker image for the Go application 
```
docker build -t your-docker-username/my-go-app:latest .
docker push your-docker-username/my-go-app:latest
```
### 7. Deploy the Helm Chart
Create a `values.yaml` file for the Helm chart configuration:
```
environmentName: default-environment
baseImage: golang:1.17
cpuRequest: "500m"
memoryRequest: "1Gi"
gpuRequest: ""
packages:
  - flask
  - requests
service:
  type: ClusterIP
  port: 8080
ingress:
  enabled: true
  host: environment.example.com

ssh:
  enabled: true
  user: "user"
  password: "password"
  host: "ssh.example.com"

```

Deploy the Helm chart:
```
helm install my-dynamic-environment /path/to/helm-chart -f values.yaml
```
### Accessing the Services
#### Access the Go Application
Use the following command to get the external IP of the Go application:
```
kubectl get ingress my-dynamic-environment-ingress
```
#### Access the SSH Server
Use the following command to get the external IP of the SSH server:
```
kubectl get ingress my-dynamic-environment-ssh-ingress
```
#### Access Prometheus and Grafana
Use the following commands to get the external IPs of Prometheus and Grafana:
```
kubectl get svc -n monitoring prometheus-kube-prometheus-prometheus
kubectl get svc -n monitoring grafana
```
## Contributing
Feel free to open issues or submit pull requests if you have any improvements or bug fixes.

## License
This project is licensed under the MIT License.