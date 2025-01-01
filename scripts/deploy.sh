#!/bin/bash

# Deploy backend
kubectl apply -f backend/

# Deploy monitoring stack
kubectl apply -f monitoring/prometheus/
kubectl apply -f monitoring/grafana/

# Deploy autoscaling
kubectl apply -f autoscaling/cluster-autoscaler.yaml
kubectl apply -f autoscaling/hpa/
