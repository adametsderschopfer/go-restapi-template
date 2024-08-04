# Go Rest Api Template

# App
To launch the application do the following

1. Set up local/prod config files
2. Run the build command ```$ make build-app```
3. Run the application launch command ```$ make run-app-local``` or ```$ make run-app-prod```

# Migrations
To deploy migrations, you can use the cmd/migrator tool

It is based on the tool https://github.com/golang-migrate/migrate for this occasion you can use the commands from this library

For work, you need:
1. Set up the local/prod configuration file
    1. Create migrations if they don't exist yet
2. Build a tool for rolling up migrations with the command ```$ make build-migrator-up```
3. Run migrations using the command ```$ make run-migrate-up-local``` or ```$ make run-migrate-up-prod```

# Locale Development

### Docker Compose Local environments
Docker Compose Local file contains the following environment variables:

In this solution, local compose is used exclusively to create a database

* `POSTGRES_USER`
* `POSTGRES_PASSWORD`
* `PGADMIN_DEFAULT_EMAIL`
* `PGADMIN_DEFAULT_PASSWORD`

### Access to postgres:
* `localhost:5432`
* **Username:** postgres (as a default)
* **Password:** test_password (as a default)

### Access to PgAdmin:
* **URL:** `http://localhost:5050`
* **Username:** admin@service.example (as a default)
* **Password:** admin (as a default)

### Add a new server in PgAdmin:
* **Host name/address** `postgres`
* **Port** `5432`
* **Username** as `POSTGRES_USER`
* **Password** as `POSTGRES_PASSWORD`