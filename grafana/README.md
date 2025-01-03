# Grafana

Grafana

## Local setup

### Start

Initialize a clean version of the database:

```sh
make start
```

### Access

Open the Grafana service in the browser:

```sh
minikube service grafana
```

> [!NOTE]
>
> Default credentials are `admin` and `admin`.

### Setup

Under `Connections` menu select `Data sources`.
Insert the following values:

- Name: `BugBase database`
- Host URL: `<host_url>:<host_port>`
- Database name: `bugbase`
- Username: `bugbase`
- Password: `password`
- TLS/SSL Mode: `disable`

Press the `Save & test` button to apply the configuration.
Under `Dashboards` menu import a new dashboard from the JSON configuration file in the `dashboards` folder.

### Stop

Stop the Docker container:

```sh
make stop
```
