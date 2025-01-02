# Database

PostgreSQL database.

## Setup

> [!CAUTION]
>
> The following commands will delete all database data.

### Start

Initialize a clean version of the database:

```sh
make start
```

### Mock

Initialize a clean version of the database with mock data:

```sh
make start-mock
```

### Stop

Stop the Docker container:

```sh
make stop
```

## ERD

Visit [drawDB](https://www.drawdb.app) and import `doc/erd.ddb` diagram.

