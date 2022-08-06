#!/bin/bash

# Using kind to create a cluster at a specific api version https://kind.sigs.k8s.io/docs/user/quick-start/
kind create cluster --image kindest/node:v1.21.2 --name kind-1.21.2 --config ./test/cluster/kind-config.yaml
# helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
# helm repo add podinfo https://stefanprodan.github.io/podinfo
# helm repo add fluxcd https://charts.fluxcd.io
# helm repo update

# # Existing Ingress
# helm upgrade --install ingress3 ingress-nginx/ingress-nginx \
#     --set controller.name=ingress3 \
#     --set controller.ingressClass=ingress3 \
#     --version 3.26.0 \
#     --namespace ingress3 \
#     --create-namespace \
#     -f ingress-3-26-0-values.yaml

# # Something to test with
# helm upgrade --install my-release podinfo/podinfo \
#     --namespace dev \
#     --create-namespace \
#     -f podinfo-values.yaml

# # New Ingress
# helm upgrade --install ingress4 ingress-nginx/ingress-nginx \
#     --set controller.ingressClass=ingress4 \
#     --set controller.watchIngressWithoutClass=true \
#     --set controller.ingressClassResource.controllerValue=k8s.io/ingress4 \
#     --set controller.ingressClassResource.name=ingress4 \
#     -f ingress-4-1-4-values.yaml \
#     --namespace ingress4 \
#     --create-namespace \
#     --version 4.0.18

# # Installing Flux v1 from https://fluxcd.io/legacy/flux/tutorials/get-started/
# # Create namespace and apply CRD's
# kubectl create ns flux
# kubectl apply -f crds.yaml

# # Create secret with credential in the cluster https://github.com/fluxcd/flux/blob/master/chart/flux/README.md#flux-with-git-over-https

# # Run Helm chart with secret 
# helm upgrade -i flux fluxcd/flux \
# --namespace flux \
# --set git.branch="main" \
# --set git.url="disp@vs-ssh.visualstudio.com:v3/disp/RD%20Enablement/concurrent-nginx-ingresss-poc" \
# --set git.path="clusters/kind-1.21.2" \
# --set git.secretName="flux-ssh" \
# --set registry.acr.enabled="false" \
# --set syncGarbageCollection.enabled="true" \
# --set sync.timeout=3m \
# --set manifestGeneration="true"

# # Install flux helm controller
# helm upgrade --install helm-operator fluxcd/helm-operator \
#    --namespace flux \
#    --set git.ssh.secretName="flux-ssh" \
#    --set helm.versions=v3