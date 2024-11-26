# Project management software

## Main statuses

[![(framework) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/framework.yaml/badge.svg?branch=main&event=push)](https://github.com/RSOPMS/app/actions/workflows/framework.yaml)

[![(app-static) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-static.yaml/badge.svg?branch=main&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-static.yaml)

[![(app-issue) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-issue.yaml/badge.svg?branch=main&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-issue.yaml)

## Dev statuses

[![(framework) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/framework.yaml/badge.svg?branch=dev&event=push)](https://github.com/RSOPMS/app/actions/workflows/framework.yaml)

[![(app-static) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-static.yaml/badge.svg?branch=dev&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-static.yaml)

[![(app-issue) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-issue.yaml/badge.svg?branch=dev&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-issue.yaml)

## Local Kubernetes setup

### Prerequsites

Install (in the specified order) [Docker](https://docs.docker.com/engine/install/), [minikube](https://minikube.sigs.k8s.io/docs/) and [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) for local development.

### Start the local cluster

Start the local minikube cluster:

```sh
minikube start
```

Validate connection to Docker daemon:

```sh
# This command should output "minikube"
echo $MINIKUBE_ACTIVE_DOCKERD
```

> [!TIP]
>
> Open kubernetes dashboard proxy in the browser:
>
> ```sh
> minikube dashboard
> ```

### Building Docker images and starting the cluster

Set the Docker context to the local minikube cluster:

```sh
eval $(minikube docker-env)
```

> [!WARNING]
>
> This command sets the Docker daemon namespace to the minikube cluster only for the current terminal session.

Push all Docker images to the in-cluster Docker daemon and apply k8s deployments:

```sh
make k8s-dev
```

Remove k8s deployments:

```sh
make k8s-dev-stop
```

> [!NOTE]
>
> Review `Makefile` for detailed build commands.

### Stop the local cluster

Stop the local minikube cluster:

```sh
minikube stop
```
