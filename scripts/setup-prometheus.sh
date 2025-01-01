#!/bin/bash

# Create a namespace for monitoring
kubectl create namespace monitoring

# Install Prometheus
helm install prometheus prometheus-community/kube-prometheus-stack --namespace monitoring

# Install Grafana
helm install grafana grafana/grafana --namespace monitoring

# Apply additional configurations if needed
kubectl apply -f monitoring/prometheus/prometheus.yaml
kubectl apply -f monitoring/alertmanager/config.yaml