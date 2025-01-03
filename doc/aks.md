# AKS setup

- Register `Azure for Students` account
- Go to `Settings`, `Resource providers` search for `Microsoft.Compute` and register it
- Create a new `Resource group`
- Enable `Microsoft.Compute`
- Add a new `Kubernetes service`
- Under `Kubernetes resources` - `Run command` run the follwing code block:
```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

helm install ingress-nginx ingress-nginx/ingress-nginx `
  --set controller.service.annotations."service\.beta\.kubernetes\.io/azure-load-balancer-health-probe-request-path"=/healthz `
  --set controller.service.externalTrafficPolicy=Local
```
- Under the new `Kubernetes service - Overview` tab, `Create` -> `Apply a YAML` and copy all YAML files from the repository
