# Database

SQLite database.

## Setup

The only prerequesite is the `sqlite3` binary (available in each distribution's package manager).

> [!CAUTION]
>
> The following commands will delete all database data.

### Clean

Initialize a clean version of the database:

```sh
make clean
```

### Mock

Initialize a clean version of the database with mock data:

```sh
make mock
```

## ERD

Visit [drawDB](https://www.drawdb.app) and import `doc/erd.ddb` diagram.
