# :memo: BugBase

Project management software

## :rocket: Main statuses

[![(framework) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/framework.yaml/badge.svg?branch=main&event=push)](https://github.com/RSOPMS/app/actions/workflows/framework.yaml)

[![(app-static) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-static.yaml/badge.svg?branch=main&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-static.yaml)

[![(app-issue) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-issue.yaml/badge.svg?branch=main&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-issue.yaml)

[![(app-bulk) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-bulk.yaml/badge.svg?branch=main&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-bulk.yaml)

[![(app-login) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-login.yaml/badge.svg?branch=main&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-login.yaml)

## :construction: Dev statuses

[![(framework) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/framework.yaml/badge.svg?branch=dev&event=push)](https://github.com/RSOPMS/app/actions/workflows/framework.yaml)

[![(app-static) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-static.yaml/badge.svg?branch=dev&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-static.yaml)

[![(app-issue) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-issue.yaml/badge.svg?branch=dev&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-issue.yaml)

[![(app-bulk) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-bulk.yaml/badge.svg?branch=dev&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-bulk.yaml)

[![(app-login) Vet and test GO code](https://github.com/RSOPMS/app/actions/workflows/app-login.yaml/badge.svg?branch=dev&event=push)](https://github.com/RSOPMS/app/actions/workflows/app-login.yaml)

## :wrench: Development

Run the following command to initialize the project for local development:

```sh
make init
```

> [!TIP]
>
> Optionally install [editorconfig](https://editorconfig.org/) editor plugin for a consistent coding style

### Repository structure

Root directory contains mostly common configuration files and kubernetes configurations.
Each microservice is its own module.
Refer to each services' `README.md` file for development instructions.

## :computer: Local Kubernetes setup

### Prerequsites

Install (in the specified order) [Docker](https://docs.docker.com/engine/install/), [minikube](https://minikube.sigs.k8s.io/docs/) and [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) for local development.

### Start the local cluster

Start the local minikube cluster:

```sh
minikube start
```

> [!TIP]
>
> Open kubernetes dashboard proxy in the browser:
>
> ```sh
> minikube dashboard
> ```

Enable ingress addon:

```sh
minikube addons enable ingress
```

### Building Docker images and starting the cluster

Set the Docker context to the local minikube cluster:

```sh
eval $(minikube docker-env)
```

> [!WARNING]
>
> This command sets the Docker daemon namespace to the minikube cluster only for the current terminal session.

Validate connection to Docker daemon:

```sh
# This command should output "minikube"
echo $MINIKUBE_ACTIVE_DOCKERD
```

Push all Docker images to the in-cluster Docker daemon and apply k8s deployments:

```sh
make k8s/dev
```

Remove k8s deployments:

```sh
make k8s/dev/delete
```

> [!NOTE]
>
> Run `make` to display Makefile help:
>
> ```
> Tip:
>   review Makefile for more detailed targets
> Usage:
>   help                 Print this help message
>   init                 Initialize the repository for local development
>   clean                Remove all generated files
>   k8s/dev              Build and start the k8s cluster                   (alias: kd)
>   k8s/dev/build        Build Docker container images                     (alias: kdb)
>   k8s/dev/start        Apply k8s configurations                          (alias: kds)
>   k8s/dev/delete       Delete the k8s cluster                            (alias: kdd)
>   migrate/up           Run up migrations                                 (alias: mu)
>   migrate/down         Run down migrations                               (alias: md)
>   migrate/fresh        Rebuild the database                              (alias: mf)
>   migrate/fresh/seed   Rebuild the database and seed it                  (alias: mfs)
> ```

### Stop the local cluster

Stop the local minikube cluster:

```sh
minikube stop
```

## :bug: Debugging

### Prerequsites

For debugging GO code install [delve](https://github.com/go-delve/delve).

### Debug

VSCode is preconfigured with necessary run configurations.
They can be accessed and ran under the `Run and Debug` tab (`Ctrl+Shift+D`).
